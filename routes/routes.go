package routes

import (
	"github.com/irenegujw/ecommerce-go/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {

	incomingRoutes.POST("/users/signup",controllers.SignUp())
	incomingRoutes.POST("/users/login",controllers.Login())
	incomingRoutes.POST("/admin/addproduct",controllers.ProductViewerAdmin())
	incomingRoutes.GET("/users/productview",controllers.searchProduct())
	incomingRoutes.GET("/users/search",controllers.searchProductByQuery())

}