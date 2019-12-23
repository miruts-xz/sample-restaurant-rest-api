package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/julienschmidt/httprouter"
	"github.com/miruts/sample-restaurant-rest-api/delivery/http/handler"
	"github.com/miruts/sample-restaurant-rest-api/user/repository"
	"github.com/miruts/sample-restaurant-rest-api/user/service"
	"net/http"
)

func main() {

	dbconn, err := gorm.Open("postgres", "postgres://postgres:postgres@localhost/restrestaurantdb?sslmode=disable")

	if err != nil {
		panic(err)
	}
	defer dbconn.Close()

	//role related infrastructure
	roleRepo := repository.NewRoleGormRepo(dbconn)
	roleSrv := service.NewRoleService(roleRepo)
	adminRoleHandler := handler.NewAdminRoleHandler(roleSrv)

	router := httprouter.New()
	router.GET("/v1/admin/roles", adminRoleHandler.GetRoles)
	http.ListenAndServe(":8181", router)
	/*commentRepo := repository.NewCommentGormRepo(dbconn)
	commentSrv := service.NewCommentService(commentRepo)

	adminCommentHandler := handler.NewAdminCommentHandler(commentSrv)

	router := httprouter.New()

	router.GET("/v1/admin/comments/:id", adminCommentHandler.GetSingleComment)
	router.GET("/v1/admin/comments", adminCommentHandler.GetComments)
	router.PUT("/v1/admin/comments/:id", adminCommentHandler.PutComment)
	router.POST("/v1/admin/comments", adminCommentHandler.PostComment)
	router.DELETE("/v1/admin/comments/:id", adminCommentHandler.DeleteComment)
	http.ListenAndServe(":8181", router)
	*/
}
