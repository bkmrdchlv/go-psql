package api

import (
	"github.com/gin-gonic/gin"
	"github.com/go-psql/api/handlers"
	"github.com/go-psql/config"
)

func SetUpRouter(h handlers.Handler, cfg config.Config) (r *gin.Engine) {
	r = gin.New()

	// cors
	r.Use(CORSMiddleware())

	r.GET("/something_list", h.GetSomethingList)
	r.GET("/something:id", h.GetSomethingByID)
	r.POST("/something", h.CreateSomething)
	r.PUT("/something", h.UpdateSomething)
	r.PATCH("/something", h.PatchSomething)
	r.DELETE("/something", h.DeleteSomething)

	return
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
