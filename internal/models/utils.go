package models


func ToUserModel(userDB *UserDB) (*User) {
	user := &User{
		ID: userDB.ID,
		Name: userDB.Name,
		Surname: userDB.Surname,
		Email: userDB.Email,
		Password: userDB.Password,
	}

	if userDB.About.Valid {
		user.About = userDB.About.String
	}

	if userDB.ImgUrl.Valid {
		user.ImgUrl = userDB.ImgUrl.String
	}

	return user
}

func ToPlaceModel(placeDB *PlaceDB) (*Place) {
	place := &Place{
		ID: placeDB.ID,
		Name: placeDB.Name,
		Description: placeDB.Description,
		Category: placeDB.Category,
	}

	if placeDB.About.Valid {
		place.About = placeDB.About.String
	}

	if placeDB.ImgUrl.Valid {
		place.ImgUrl = placeDB.ImgUrl.String
	}

	return place
}