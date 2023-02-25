package utils

import (
	"Diploma/internal/customErrors"
	"Diploma/internal/models"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"reflect"
	"strings"

	log "Diploma/pkg/logger"
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
		return &models.AccessDetails{}, customErrors.ErrNoTokenInContext
	}

	au, ok := ctxau.(models.AccessDetails)
	if !ok {
		return &models.AccessDetails{}, customErrors.ErrNoTokenInContext
	}

	if reflect.ValueOf(au).IsZero() {
		return &models.AccessDetails{}, customErrors.ErrNoTokenInContext
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

	Validate, err := govalidator.ValidateStruct(object)
	if err != nil || !Validate {
		log.Error(err.Error())
		return customErrors.ErrValidation
	}
	return err
}

func SaveImageFromRequest(c *gin.Context, httpRequestKey string) (string, error) {
	avatarFile, handler, err := c.Request.FormFile(httpRequestKey)
	if err != nil {
		log.Error("first err = ", err.Error())
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
	case "webp":
	default:
		return "", customErrors.ErrWrongExtension
	}

	options, err := encoder.NewLossyEncoderOptions(encoder.PresetDefault, 40)
	if err != nil {
		log.Error("second err = ", err.Error())
		return "", err
	}

	resultFileName := newFilename + "." + filenameExtension
	output, err := os.Create(viper.GetString("img_path") + "/" + resultFileName)
	if err != nil {
		log.Error("third err = ", err.Error())
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
		return newFilename, nil
	}

	if err := webp.Encode(output, img, options); err != nil {
		log.Error("fourth err = ", err.Error())
		return "", err
	}
	return newFilename, nil
}

func TryBuildImgUrl(imgUUID string) string {
	if imgUUID == "" {
		return ""
	}

	serverName := viper.GetString("server.name")
	return serverName + "/images/" + imgUUID + ".webp"
}

func GetPageQueryParamFromRequest(c *gin.Context) string {
	pageParam := c.DefaultQuery("page", "1")
	return pageParam
}

func SendMessage(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, models.Message{
		Message: message,
	})
}
