package services

import (
	"fmt"

	"github.com/ruilisi/Rocket.Chat.Go.SDK/conf"
	"github.com/ruilisi/Rocket.Chat.Go.SDK/models"
	"github.com/ruilisi/Rocket.Chat.Go.SDK/rest"
)

func RegisterAdmin(reg *models.RegisterUserRequest, client *rest.Client) (*rest.AddUserToRoleResponse, error) {

	client.RegisterUser(reg)

	fmt.Println(conf.GetEnv("ADMIN_USERNAME"), conf.GetEnv("ADMIN_PASSWORD"))
	login := models.UserCredentials{
		Username: conf.GetEnv("ADMIN_USERNAME"),
		Password: conf.GetEnv("ADMIN_PASSWORD"),
	}
	client.Login(&login)

	addUserToRoleParams := &models.AddUserToRole{RoleName: "admin", Username: reg.Username}
	response, err := client.AddUserToRole(addUserToRoleParams)
	return response, err

}
