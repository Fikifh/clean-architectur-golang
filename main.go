package main

import (
	"fmt"
	"log"

	"incentrick-restful-api/app/usecase/authentication"
	"incentrick-restful-api/config"
	"incentrick-restful-api/config/database"
	auth_handler "incentrick-restful-api/handler/auth"
	"incentrick-restful-api/repository/auth"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.New()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?loc=Local&parseTime=true", cfg.DBUserName, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBDatabaseName)
	db, err := database.Connect(dsn)
	if err != nil {
		log.Panicln("Failed to Initialized mysql DB:", err)
	}

	// video_repo := video_repository_v1.New(db)
	// crud_video_uc := crud_video.NewUseCase(video_repo)
	// hVideo := handler.NewHandler(crud_video_uc)

	// course_repo := course_repository_v1.New(db)
	// crud_course_uc := crud_course.NewUseCase(course_repo)
	// hCourse := handler_course.NewHandler(crud_course_uc)

	// router.GET("/video", hVideo.GetVideos)
	// router.GET("/video/:id", hVideo.GetVideo)
	// router.POST("/video", hVideo.CreateVideo)
	// router.POST("/video/:id", hVideo.UpdateVideo)
	// router.DELETE("/video/:id", hVideo.DeleteVideo)

	// router.GET("/course", hCourse.GetCourses)

	authRepo := auth.New(db)
	authUseCase := authentication.NewUseCase(authRepo)
	handlerAuth := auth_handler.NewHandler(authUseCase)

	router := gin.Default()
	v1 := router.Group("v1")       //V1
	authRouter := v1.Group("auth") //AUTH GROUP
	authRouter.GET("/login", handlerAuth.Login)
	authRouter.GET("/google/login", handlerAuth.OauthGoogleLogin)
	authRouter.POST("/google/callback", handlerAuth.OauthGoogleCallback)

	router.Run(fmt.Sprintf(":%d", cfg.Port))
}
