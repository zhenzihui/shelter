package user

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"projectsuika.com/shelter/src/model"
)

// User 用户表
type User struct {
	Id         string `json:"id" gorm:"primary_key,default:uuid_generate_v3()"`
	UserName   string `json:"user_name" gorm:"unique"`
	Password   string `json:"password" gorm:"default:e10adc3949ba59abbe56e057f20f883e"`
	BucketName string `json:"bucket_name"`
}

func Create(user User) {
	res := model.DB.Debug().Create(&user)
	log.Panicln(res.Error.Error())
}

func FindById(id string) User {
	result := User{}
	model.DB.Model(&User{Id: id}).First(&result)
	return result
}
func FindOne() User {
	result := User{}
	model.DB.Model(&User{}).First(&result)
	return result
}

func CreateDefault() User {
	user := FindById("0000")
	defaultUser := User{
		Id:         "0000",
		UserName:   "admin",
		Password:   "e10adc3949ba59abbe56e057f20f883e",
		BucketName: "default",
	}
	if len(user.UserName) == 0 {
		Create(defaultUser)
	}
	return defaultUser
}

func Login(userName string, password string) *User {
	h := md5.New()
	io.WriteString(h, password)
	var query = User{UserName: userName, Password: fmt.Sprintf("%x", h.Sum(nil))}
	var result User
	err := model.DB.Debug().Where(query).First(&result).Error
	if err != nil {
		return nil
	}
	return &result
}
