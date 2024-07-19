package user_controllers

import (
	"github.com/adi-kmt/brew-scylla/internal/user"
	"github.com/gofiber/fiber/v2"
)

func UserHandler(router fiber.Router, o *user.OrderService, p *user.ProductService) {
	router.Get("/stores", getAllStores(p))

	router.Get("/products/packs/:storeId", getProductPacks(p))
	router.Get("/products/collections/:storeId", getProductsCollection(p))
	router.Get("/products/featured/:storeId", getFeaturedProductCollections(p))

	router.Get("/products/:storeId", getProductsByStore(p))
	router.Get("/products/search", searchProducts(p))
	router.Get("/products/:productId", getProductById(p))

	router.Post("/cart", addProductToCart(o))
	router.Post("/order", checkoutOrder(o))
	router.Get("/order", getOrdersByUserId(o))
}
