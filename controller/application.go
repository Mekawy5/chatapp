package controller

import (
	"net/http"
	"strconv"

	"github.com/Mekawy5/chatapp/domain/application"
	"github.com/gin-gonic/gin"
)

type ApplicationController struct {
	Srv     *gin.Engine
	Service *application.ApplicationService
}

func NewApplicationController(srv *gin.Engine, s *application.ApplicationService) *ApplicationController {
	ac := &ApplicationController{
		Srv:     srv,
		Service: s,
	}

	ac.Srv.GET("/applications", ac.All)
	ac.Srv.GET("/applications/:token", ac.Find)
	ac.Srv.DELETE("/applications/:token", ac.Delete)

	return ac
}

func (ac *ApplicationController) All(c *gin.Context) {
	apps := ac.Service.GetAll()
	c.JSON(http.StatusOK, gin.H{"apps": application.GetApplications(apps)})
}

func (ac *ApplicationController) Find(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	app := ac.Service.Get(uint(id))
	c.JSON(http.StatusOK, gin.H{"app": application.GetApplication(app)})
}

func (ac *ApplicationController) Delete(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Application Deleted"})
}
