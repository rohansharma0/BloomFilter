package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rohansharma0/bloomfiler/internal/config"
	"github.com/rohansharma0/bloomfiler/internal/service"
)

type UsernameRequest struct {
	Username string `json:"username"`
}

func SetupRouter(cfg *config.Config) *gin.Engine {
	router := gin.Default()

	routes := router.Group("/api")
	{
		routes.GET("/attempt", func(c *gin.Context) {
			username := c.Query("username")

			if username == "" {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "username is not present.",
				})
				return
			}

			c.JSON(http.StatusOK, !service.IsUsernameExists(username))
		})

		routes.POST("/register", func(c *gin.Context) {
			var req UsernameRequest

			if err := c.BindJSON(&req); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "Invalid JSON",
				})
				return
			}

			if req.Username == "" {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "bad request.",
				})
				return
			}
			if service.IsUsernameExists(req.Username) {
				c.JSON(http.StatusConflict, gin.H{
					"error": "username may already be taken.",
				})
				return
			}
			service.AddUsername(req.Username)
			c.JSON(http.StatusCreated, gin.H{
				"message": req.Username + " is registered.",
			})
		})
	}
	return router
}
