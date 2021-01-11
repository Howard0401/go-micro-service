package main

import (
	// "go-micro-service/domain/service"

	"go-micro-service/domain/repository"
	userService "go-micro-service/domain/service"
	"go-micro-service/handler"
	pb "go-micro-service/proto/user"
	"log"

	"github.com/micro/go-micro/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// pb "user/proto"

func main() {
	// Create service //microv3 TODO
	// srv := service.New(service.Name("go-micro-service"), service.Version("latest"))
	// srv := service.New()

	srv := micro.NewService(
		micro.Name("micro-server"),
	)

	// Register handler
	srv.Init()

	config := "root:@Aamother0120@tcp(127.0.0.1:3306)/micro?parseTime=true&charset=utf8&parseTime=true&loc=Local"
	orm, err := gorm.Open(mysql.Open(config), &gorm.Config{NamingStrategy: schema.NamingStrategy{SingularTable: true}})

	if err != nil {
		log.Printf("open mysql err: %v", err)
	}

	db, err := orm.DB()

	if err != nil {
		log.Printf("create sql DB err: %v", err)
	}

	defer db.Close()

	// Create Table
	// rp := repository.NewUserRepository(orm)
	// rp.InitTable()

	// Register Service
	userDataService := userService.NewUserDataService(repository.NewUserRepository(orm))
	// Register handler
	err = pb.RegisterUserHandler(srv.Server(), &handler.User{UserDataService: userDataService})
	// err = srv.Handle(&handler.User{UserDataService: userDataService})
	if err != nil {
		log.Printf("register pb failed: %v", err)
	}
	// Run service
	if err := srv.Run(); err != nil {
		log.Printf("create sql DB err: %v", err)
	}
}

// func main() {
// 	// Create service
// 	srv := service.New(
// 		service.Name("user"),
// 		service.Version("latest"),
// 	)

// 	// Register handler
// 	pb.RegisterUserHandler(srv.Server(), new(handler.User))

// 	// Run service
// 	if err := srv.Run(); err != nil {
// 		logger.Fatal(err)
// 	}
// }
