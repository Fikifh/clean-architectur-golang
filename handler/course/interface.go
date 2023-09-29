package handler_course

import "github.com/gin-gonic/gin"

type RestHandler interface {
	GetCourses(c *gin.Context)
	GetCourse(c *gin.Context)
	CreateCourse(c *gin.Context)
	UpdateCourse(c *gin.Context)
	DeleteCourse(c *gin.Context)
}
