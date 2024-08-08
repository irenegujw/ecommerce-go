package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/irenegujw/controllers"
)

func UserRoutes(incomingRoutes *gin.Engine) {

	incomingRoutes.POST("/users/signup", controllers.SignUp())
	incomingRoutes.POST("/users/login", controllers.Login())
	incomingRoutes.POST("/admin/addproduct", controllers.ProductViewerAdmin())
	incomingRoutes.GET("/users/productview", controllers.searchProduct())
	incomingRoutes.GET("/users/search", controllers.searchProductByQuery())

}
