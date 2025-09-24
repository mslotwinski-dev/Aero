package src

import (
	"aero/src/app/cmd"
	"aero/src/app/gateway"
	"aero/src/config"
	"aero/src/database/repository"
	"aero/src/internal/auth"
	"aero/src/internal/chat"
	"aero/src/internal/media"
	"aero/src/internal/user"
	"aero/src/utility/helpers"
	"aero/src/utility/monitoring"
	"aero/src/web"
	"aero/src/window"
	"fmt"
)

func Init() {
	fmt.Println("App initialized")

	cmd.Run()
	gateway.Run()
	config.Run()
	repository.Run()
	chat.Run()
	media.Run()
	user.Run()
	helpers.Run()
	monitoring.Run()
	web.Run()
	window.Run()
	auth.Run()
	repository.Run()
}
