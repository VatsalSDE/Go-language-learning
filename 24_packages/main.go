package mai

import (
	"fmt"

	"github.com/codersgyan/podcast/auth"
	"github.com/codersgyan/podcast/user"
	"github.com/fatih/color"
)

func mai() {
	auth.LoginWithCredentials("codersgyan", "secret")
	session := auth.GetSession()

	fmt.Println("session", session)

	user := user.User{
		Email: "user@email.com",
		// Name:  "John Doe",
	}

	// fmt.Println(user.Email, user.Name)
	color.Green(user.Email)

}
