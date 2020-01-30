package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/julienschmidt/httprouter"
	"github.com/miruts/food-api/delivery/http/handler"
	"github.com/miruts/food-api/entity"
	"github.com/miruts/food-api/user/repository"
	"github.com/miruts/food-api/user/service"
	"net/http"
)

func main() {

	dbconn, err := gorm.Open("postgres", "postgres://postgres:postgres@localhost/restrestaurantdb?sslmode=disable")

	if err != nil {
		panic(err)
	}
	defer dbconn.Close()
	_ = dbconn.CreateTable(&entity.Role{}).GetErrors()
	//role related infrastructure
	roleRepo := repository.NewRoleGormRepo(dbconn)
	roleSrv := service.NewRoleService(roleRepo)
	adminRoleHandler := handler.NewAdminRoleHandler(roleSrv)

	router := httprouter.New()

	router.GET("/v1/admin/roles", adminRoleHandler.GetRoles)
	router.GET("/v1/admin/roles/:id", adminRoleHandler.GetRole)
	router.POST("/v1/admin/roles", adminRoleHandler.PostRole)
	router.PUT("/v1/admin/roles/:id", adminRoleHandler.PutRole)
	router.DELETE("/v1/admin/roles/:id", adminRoleHandler.DeleteRole)
	err = http.ListenAndServe(":5000", router)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
}
