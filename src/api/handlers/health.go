package handlers

import (
	"github.com/Masih-Ghasri/GolangBackend/api/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

type header struct {
	UserId  string
	Browser string
}

type personData struct {
	FirstName    string `json:"first_name" binding:"required,alpha,min=4,max=10"`
	LastName     string `json:"last_name" binding:"required,alpha,min=6,max=20"`
	MobileNumber string `json:"mobile_number" binding:"required,mobile,min=11,max=11"`
}

type HealthHandler struct {
}

// @Summary Health Check
// @Description Health Check
// @Tags Health
// @Accept json
// @Produce json
// @Success 200 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse{ValidationErrors=[]validation.ValidationError} "Failed"
// @Router /v1/health/ [get]
func GetUsers(c *gin.Context) {
	// handler code
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, helper.GenerateBaseResponse("working", true, 0))
}
