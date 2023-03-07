package router

import "github.com/gin-gonic/gin"

func Init() (r *gin.Engine) {
	r = gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// routes
	r.POST("/auth")

	apiv1 := r.Group("/api/v1")
	// apiv1.Use(jwt) // middleware
	// protected routes
	{
		somethingApi := apiv1.Group("/something")
		{
			// get all
			somethingApi.GET("/")
			// create one
			somethingApi.POST("/")
			// get one
			somethingApi.GET("/:id")
			// delete one
			somethingApi.DELETE("/:id")
		}
	}

	return
}
