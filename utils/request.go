package utils

import (
	"Diploma/internal/customErrors"
	"Diploma/internal/models"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/go-sanitize/sanitize"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/viper"

	"github.com/kolesa-team/go-webp/encoder"
	"github.com/kolesa-team/go-webp/webp"
)

func GetAUFromContext(c *gin.Context) (*models.AccessDetails, error) {
	ctxau, ok := c.Get("access_details")
	if !ok {
		return nil, customErrors.ErrNoTokenInContext
	}

	au, ok := ctxau.(models.AccessDetails)
	if !ok {
		return nil, customErrors.ErrNoTokenInContext
	}
	
	return &au, nil
}

func ValidateAndSanitize(object interface{}) error {
	s, err := sanitize.New()
	if err != nil {
		return customErrors.ErrSanitizer
	}

	err = s.Sanitize(object)
	if err != nil {
		return customErrors.ErrSanitizing
	}

	valid, err := govalidator.ValidateStruct(object)
	if err != nil || !valid {
		return customErrors.ErrValidation
	}
	return err
}

func SaveImageFromRequest(c *gin.Context, httpRequestKey string) (string, error) {
	serverName := viper.GetString("server.name")

	avatarFile, handler, err := c.Request.FormFile(httpRequestKey)
	if err != nil {
		return "", err
	}
	defer avatarFile.Close()

	var img image.Image

	filenameParts := strings.Split(handler.Filename, ".")
	filenameExtension := strings.ToLower(filenameParts[len(filenameParts)-1])
	newFilename := uuid.NewV4().String()
	
	switch filenameExtension {
		case "jpg", "jpeg", "png":
			filenameExtension = "webp"
		case "ico","woff","swg","webp","webm","gif":
		default:
			return "", customErrors.ErrWrongExtension
	}

	options, err := encoder.NewLossyEncoderOptions(encoder.PresetDefault, 40)
	if err != nil {
		return "", err
	}

	resultFileName := newFilename + "." + filenameExtension
	output, err := os.Create(viper.GetString("img_path") + "/" + resultFileName)
	if err != nil {
		return "", err
	}
	defer output.Close()

	switch filenameExtension {
	case "jpg", "jpeg":
		img, err = jpeg.Decode(avatarFile)
		if err != nil {
			return "", err
		}
	case "png":
		img, err = png.Decode(avatarFile)
		if err != nil {
			return "", err
		}
	default:
		_, err := io.Copy(output, avatarFile)
		if err != nil {
			return "", err
		}
		return serverName + "/images/" + resultFileName, nil
	}

	if err := webp.Encode(output, img, options); err != nil {
		return "", err
	}
	return serverName + "/images/" + resultFileName, nil
}

func GetPageQueryParamFromRequest(c *gin.Context) (int, error) {
	pageParam := c.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageParam)
	if err != nil {
		return 0, err
	}
	return page, nil
}

func SendErrorMessage(c *gin.Context, statusCode int, errorMessage string) () {
	c.JSON(statusCode, models.ErrorMessage{
		Message: errorMessage,
	})
}