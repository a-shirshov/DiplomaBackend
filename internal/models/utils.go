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