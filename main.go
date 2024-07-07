package main

import (
	"github.com/yellowsky2000/users/handler"
	proto "github.com/yellowsky2000/users/proto"

	otp "github.com/yellowsky2000/otp/proto"
	adminpb "github.com/yellowsky2000/pkg/service/proto"

	"micro.dev/v4/service"
	"micro.dev/v4/service/logger"
	"micro.dev/v4/service/store"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("user"),
	)
	srv.Init()

	hd := handler.NewUser(
		store.DefaultStore,
		otp.NewOtpService("otp", srv.Client()),
	)

	proto.RegisterUserHandler(srv.Server(), hd)
	adminpb.RegisterAdminHandler(srv.Server(), hd)

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
