package handler_course

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"incentrick-restful-api/app/usecase/crud_course"
	"incentrick-restful-api/entity"
)

type restHandler struct {
	crud_course_uc crud_course.UseCase
}

func NewHandler(crud_course_uc crud_course.UseCase) RestHandler {
	return &restHandler{crud_course_uc: crud_course_uc}
}

func (h *restHandler) GetCourses(c *gin.Context) {
	data, err := h.crud_course_uc.GetAll()
	if err == nil {
		c.JSON(http.StatusOK, SuccessResponse{Data: data})
	} else {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
	}
}

func (h *restHandler) GetCourse(c *gin.Context) {
	paramId := c.Param("id")
	id, _ := strconv.Atoi(paramId)

	data, err := h.crud_course_uc.Get(id)
	if err == nil {
		c.JSON(http.StatusOK, SuccessResponse{Data: data})
	} else {
		if errors.Is(err, crud_course.ErrCourseNotFound) {
			c.JSON(http.StatusNotFound, ErrorResponse{Message: err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		}
	}
}

func (h *restHandler) CreateCourse(c *gin.Context) {
	param := &entity.Courses{}
	err := c.ShouldBindBodyWith(&param, binding.JSON)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return
	}

	result, err := h.crud_course_uc.Create(*param)
	if err == nil {
		c.JSON(http.StatusOK, SuccessResponse{Data: result})
	} else {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
	}
}

func (h *restHandler) UpdateCourse(c *gin.Context) {
	param := &entity.Courses{}
	err := c.ShouldBindBodyWith(&param, binding.JSON)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return
	}
	paramId := c.Param("id")
	param.Id, _ = strconv.Atoi(paramId)

	result, err := h.crud_course_uc.Update(*param)
	if err == nil {
		c.JSON(http.StatusOK, SuccessResponse{Data: result})
	} else {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
	}
}

func (h *restHandler) DeleteCourse(c *gin.Context) {
	paramId := c.Param("id")
	id, _ := strconv.Atoi(paramId)

	err := h.crud_course_uc.Delete(id)
	if err == nil {
		c.JSON(http.StatusOK, SuccessResponse{Data: fmt.Sprintf("id:%d. successfully deleted", id)})
	} else {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
	}
}
