package database

import (
	"errors"
	
)

var (
	ErrCantFindProduct = errors.New("Product not found")
	ErrCantDecodeProducts = errors.New("Can't find product")
	ErrUserIdIsNotValid = errors.New("User ID is not valid")
	ErrCantUpdateUser = errors.New("Can't add product to cart")
	ErrCantRemoveItemCart = errors.New("Can't remove item from cart")
	ErrCantGetItem = errors.New("Can't get item from cart")
	ErrCantBuyCartItem = errors.New("Can't update the purchase")

)

func AddProductToCart(){

	
}

func RemoveProductFromCart(){
	
}

func BuyProductInCart(){
	
}	

func InstantBuyer(){
	
}