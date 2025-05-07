package handlers

import (
	"github.com/Masih-Ghasri/GolangBackend/api/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Header struct {
	UserId  string `json:"UserId" example:"user123"`
	Browser string `json:"Browser" example:"Chrome"`
}

type PersonData struct {
	FirstName string `json:"first_name" binding:"required,alpha,min=4,max=10" example:"John"`

	LastName string `json:"last_name" binding:"required,alpha,min=6,max=20" example:"Doe"`

	MobileNumber string `json:"mobile_number" binding:"required,mobile,min=11,max=11" example:"09123456789"`
}

type TestHandler struct{}

func NewTestHandler() *TestHandler {
	return &TestHandler{}
}

// Test godoc
// @Summary Service health check
// @Description Basic endpoint to test service availability
// @Tags Test
// @Accept json
// @Produce json
// @Success 200 {object} helper.BaseHttpResponse{result=string} "Service is working"
// @Router /v1/test/ [get]
func (h *TestHandler) Test(c *gin.Context) {
	c.JSON(http.StatusOK, helper.GenerateBaseResponse("Test", true, 0))
}

// Users godoc
// @Summary Get all users
// @Description Retrieves a list of all system users
// @Tags Test
// @Accept json
// @Produce json
// @Success 200 {object} helper.BaseHttpResponse{result=string} "List of users"
// @Router /v1/test/users/ [get]
func (h *TestHandler) Users(c *gin.Context) {
	c.JSON(http.StatusOK, helper.GenerateBaseResponse("Users", true, 0))
}

// UserById godoc
// @Summary Get user by ID
// @Description Retrieves user details by user ID
// @Tags Test
// @Accept json
// @Produce json
// @Param id path string true "User ID" example("123")
// @Success 200 {object} helper.BaseHttpResponse{result=object} "User details"
// @Failure 404 {object} helper.BaseHttpResponse "User not found"
// @Router /v1/test/users/{id} [get]
func (h *TestHandler) UserById(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"result": "UserById",
		"id":     id,
	}, true, 0))
}

// UserByUsername godoc
// @Summary Get user by username
// @Description Retrieves user details by username
// @Tags Test
// @Accept json
// @Produce json
// @Param username path string true "Username" example("johndoe")
// @Success 200 {object} helper.BaseHttpResponse{result=object} "User details"
// @Failure 404 {object} helper.BaseHttpResponse "User not found"
// @Router /v1/test/users/{username} [get]
func (h *TestHandler) UserByUsername(c *gin.Context) {
	username := c.Param("username")
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"result":   "UserByUsername",
		"username": username,
	}, true, 0))
}

// Accounts godoc
// @Summary Get user accounts
// @Description Retrieves all user accounts
// @Tags Test
// @Accept json
// @Produce json
// @Success 200 {object} helper.BaseHttpResponse{result=object} "List of accounts"
// @Router /v1/test/accounts/ [get]
func (h *TestHandler) Accounts(c *gin.Context) {
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"result": "Accounts",
	}, true, 0))
}

// AddUser godoc
// @Summary Add new user
// @Description Creates a new user in the system
// @Tags Test
// @Accept json
// @Produce json
// @Param id path string true "User ID" example("123")
// @Success 200 {object} helper.BaseHttpResponse{result=object} "User created"
// @Failure 400 {object} helper.BaseHttpResponse "Invalid request"
// @Router /v1/test/users/{id} [post]
func (h *TestHandler) AddUser(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"result": "AddUser",
		"id":     id,
	}, true, 0))
}

// HeaderBinder1 godoc
// @Summary Get headers (Method 1)
// @Description Demonstrates header binding - direct access
// @Tags Test
// @Accept json
// @Produce json
// @Param UserId header string true "User ID" example("user123")
// @Success 200 {object} helper.BaseHttpResponse{result=object} "Header data"
// @Router /v1/test/header1/ [get]
func (h *TestHandler) HeaderBinder1(c *gin.Context) {
	userId := c.GetHeader("UserId")
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"result": "HeaderBinder1",
		"userId": userId,
	}, true, 0))
}

// HeaderBinder2 godoc
// @Summary Get headers (Method 2)
// @Description Demonstrates header binding - struct binding
// @Tags Test
// @Accept json
// @Produce json
// @Param UserId header string true "User ID" example("user123")
// @Param Browser header string true "Browser Info" example("Chrome")
// @Success 200 {object} helper.BaseHttpResponse{result=object} "Header data"
// @Router /v1/test/header2/ [get]
func (h *TestHandler) HeaderBinder2(c *gin.Context) {
	header := Header{}
	_ = c.BindHeader(&header)
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"result": "HeaderBinder2",
		"header": header,
	}, true, 0))
}

// QueryBinder1 godoc
// @Summary Query binding (single)
// @Description Demonstrates single-value query parameter binding
// @Tags Test
// @Accept json
// @Produce json
// @Param id query string true "Item ID" example("123")
// @Param name query string true "Item Name" example("test")
// @Success 200 {object} helper.BaseHttpResponse{result=object} "Query parameters"
// @Router /v1/test/query1 [get]
func (h *TestHandler) QueryBinder1(c *gin.Context) {
	id := c.Query("id")
	name := c.Query("name")
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"result": "QueryBinder1",
		"id":     id,
		"name":   name,
	}, true, 0))
}

// QueryBinder2 godoc
// @Summary Query binding (array)
// @Description Demonstrates array query parameter binding
// @Tags Test
// @Accept json
// @Produce json
// @Param id query []string true "Item IDs" collectionFormat(multi) example("123,456")
// @Param name query string true "Item Name" example("test")
// @Success 200 {object} helper.BaseHttpResponse{result=object} "Query parameters"
// @Router /v1/test/query2 [get]
func (h *TestHandler) QueryBinder2(c *gin.Context) {
	ids := c.QueryArray("id")
	name := c.Query("name")
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"result": "QueryBinder2",
		"ids":    ids,
		"name":   name,
	}, true, 0))
}

// UriBinder godoc
// @Summary URI parameter binding
// @Description Demonstrates URI path parameter binding
// @Tags Test
// @Accept json
// @Produce json
// @Param id path int true "User ID" example(123)
// @Param name path string true "User Name" example("john")
// @Success 200 {object} helper.BaseHttpResponse{result=object} "URI parameters"
// @Failure 400 {object} helper.BaseHttpResponse "Invalid parameters"
// @Router /v1/test/binder/uri/{id}/{name} [post]
// @Security AuthBearer
func (h *TestHandler) UriBinder(c *gin.Context) {
	id := c.Param("id")
	name := c.Param("name")
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"result": "UriBinder",
		"id":     id,
		"name":   name,
	}, true, 0))
}

// BodyBinder godoc
// @Summary JSON body binding
// @Description Demonstrates JSON request body binding
// @Tags Test
// @Accept json
// @Produce json
// @Param person body PersonData true "Person Data"
// @Success 200 {object} helper.BaseHttpResponse{result=object} "Person data"
// @Failure 400 {object} helper.BaseHttpResponse "Validation error"
// @Router /v1/test/binder/body [post]
// @Security AuthBearer
func (h *TestHandler) BodyBinder(c *gin.Context) {
	p := PersonData{}
	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil,
				false, helper.ValidationError, err))
		return
	}
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"result": "BodyBinder",
		"person": p,
	}, true, 0))
}

// FormBinder godoc
// @Summary Form data binding
// @Description Demonstrates form data binding
// @Tags Test
// @Accept multipart/form-data
// @Produce json
// @Param first_name formData string true "First Name" example("John")
// @Param last_name formData string true "Last Name" example("Doe")
// @Param mobile_number formData string true "Mobile Number" example("09123456789")
// @Success 200 {object} helper.BaseHttpResponse{result=object} "Form data"
// @Router /v1/test/binder/form [post]
func (h *TestHandler) FormBinder(c *gin.Context) {
	p := PersonData{}
	_ = c.ShouldBind(&p)
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"result": "FormBinder",
		"person": p,
	}, true, 0))
}

// FileBinder godoc
// @Summary File upload
// @Description Demonstrates file upload handling
// @Tags Test
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "File to upload"
// @Success 200 {object} helper.BaseHttpResponse{result=object} "File uploaded"
// @Failure 500 {object} helper.BaseHttpResponse "Upload failed"
// @Router /v1/test/binder/file [post]
func (h *TestHandler) FileBinder(c *gin.Context) {
	file, _ := c.FormFile("file")
	err := c.SaveUploadedFile(file, "file")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			helper.GenerateBaseResponseWithError(nil, false, helper.ValidationError, err))
		return
	}
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"result": "FileBinder",
		"file":   file.Filename,
	}, true, 0))
}
