package databases

import (
	"test_erajaya/config"
	"test_erajaya/middlewares"
	"test_erajaya/models"
)

// function database untuk menambahkan user baru (registrasi)
func CreateUser(user *models.Users) (interface{}, error) {
	if err := config.DB.Create(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func LoginUser(user *models.Users) (interface{}, error) {
	var get_login models.GetLoginUser
	err := config.DB.Where("email = ?", user.Email).First(&user).Error
	if err != nil {
		return nil, err
	}

	user.Token, err = middlewares.CreateToken(int(user.ID)) // generate token
	if err != nil {
		return nil, err
	}
	// restruktur respons
	get_login.ID = user.ID
	get_login.Email = user.Email
	get_login.Password = user.Password
	get_login.Token = user.Token
	if err = config.DB.Save(user).Error; err != nil {
		return nil, err
	}
	return get_login, nil
}
