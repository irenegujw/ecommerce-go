package controllers

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os/user"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/irenegujw/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

	type Application struct {
		prodCollection *mongo.Collection
		userCollection *mongo.Collection
	}

	func NewApplication(prodCollection *mongo.Collection, userCollection *mongo.Collection) *Application {
		return &Application{
			prodCollection: prodCollection, 
			userCollection: userCollection,
		}
	}

func (app *Application) AddToCart() gin.HandlerFunc{
	return func(c *gin.Context){
		productQueryID := c.Query("product_id")
		if productQueryID == "" {
			log.Println("product id is empty")
			_ = c.AbortWithError(http.StatusBadRequest,errors.New("product id is empty"))
			return
		}

		userQueryID := c.Query("user_id")
		if userQueryID == "" {
			log.Println("user id is empty")
			_ = c.AbortWithError(http.StatusBadRequest,errors.New("user id is empty"))
			return
		}

		ProductID, err := primitive.ObjectIDFromHex(productQueryID)
		if err != nil {
			log.Println(err)
			_ = c.AbortWithError(http.StatusInternalServerError,err)
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		database.AddProductToCart(ctx, app.prodCollection, app.userCollection,ProductID, userQueryID)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError,err)
		}
		c.IndentedJSON(200,"Successfully added to cart")


	}
	
}

func RemoveItem() gin.HandlerFunc{
	
}

func GetItemFromCart() gin.HandlerFunc{
	
}

func BuyFromCart() gin.HandlerFunc{
	
}

func InstantBuy() gin.HandlerFunc{
	
}