package service

import (
	. "github.com/space-tech-dev/sms-service-notification/app/models"
	"github.com/space-tech-dev/sms-service-notification/config"
	"github.com/space-tech-dev/sms-service-notification/database"
)

type UserService struct {
	User    *User
}

var authUser *User

/**
 ** 初始化构造函数
 */

// 模块初始化函数 import 包时被调用
func init() {
	//fmt.Println("User service module init function")
}

func NewUserService() (r *UserService) {
	r = &UserService{
		User:    NewUser(),
	}
	return r
}

func SetAuthUser(user *User) {
	authUser = user
}

func GetAuthUser() (user *User) {
	return authUser
}

func (srv *UserService) GetUser(userName string) (user *User, r int) {

	user = &User{}

	db := database.DBConnection.Scopes(
		srv.User.WhereUserName(userName),
		//srv.User.WhereIsValid,
	)

	result := db.Preload("Account").First(user)

	if result.RowsAffected > 0 {
		//fmt.Printf("user: %v", user.Account)
		return user, config.API_RESULT_CODE_INIT

	} else {
		return nil, config.API_ERR_CODE_USER_UNREGISTER
	}

}
