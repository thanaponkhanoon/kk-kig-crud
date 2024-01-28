package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thanaponkhanoon/sa-65-example/controller"
	"github.com/thanaponkhanoon/sa-65-example/entity"
	"github.com/thanaponkhanoon/sa-65-example/middlewares"
)

const PORT = "8080"

func main() {
	entity.SetupDatabase()
	r := gin.Default()
	r.Use(CORSMiddleware())
	router := r.Group("/")
	{
		router.Use(middlewares.Authorizes())
		{
			// User Routes
			router.GET("/employees", controller.ListEmployees)
			router.GET("/employee/:id", controller.GetEmployee)
			router.PATCH("/employees", controller.UpdateEmployee)
			router.DELETE("/employees/:id", controller.DeleteEmployee)
			router.PATCH("/employee", controller.CreateEmployee)
			// Pcode Routes
			router.GET("/pcodes", controller.ListPcode)
			router.GET("/pcode/:id", controller.GetPcode)
			router.POST("/pcodes", controller.CreatePcode)
			router.PATCH("/pcodes", controller.UpdatePcode)
			router.DELETE("pcodes/:id", controller.DeletePcode)
			// detail Routes
			router.GET("/details", controller.ListDetail)
			router.GET("/detail/:id", controller.GetDetail)
			router.POST("/details", controller.CreateDetail)
			router.PATCH("/details", controller.UpdateDetail)
			router.DELETE("details/:id", controller.DeleteDetail)
			// Product Routes
			router.GET("/products", controller.ListProducts)
			router.GET("/product/:id", controller.GetProduct)
			router.POST("/product", controller.CreateProduct)
			router.PATCH("/products", controller.UpdateProduct)
			router.DELETE("/products/:id", controller.DeleteProduct)
			// Ccode Routes
			router.GET("/ccode", controller.ListCcode)
			router.GET("/ccode/:id", controller.GetCcode)
			router.POST("/ccode", controller.CreateCcode)
			router.PATCH("/ccode", controller.UpdateCcode)
			router.DELETE("ccode/:id", controller.DeleteCcode)
			// cusdetail Routes
			router.GET("/cusdetail", controller.ListCusdetail)
			router.GET("/cusdetail/:id", controller.GetCusdetail)
			router.POST("/cusdetail", controller.CreateCusdetail)
			router.PATCH("/cusdetails", controller.UpdateCusdetail)
			router.DELETE("cusdetails/:id", controller.DeleteCusdetail)
			// Customer Routes
			router.GET("/customer", controller.ListCustomer)
			router.GET("/customer/:id", controller.GetCustomer)
			router.POST("/customer", controller.CreateCustomer)
			router.PATCH("/customer", controller.UpdateCustomer)
			router.DELETE("/customer/:id", controller.DeleteCustomer)
		}
	}
	// Signup User Route
	// r.POST("/signup", controller.CreateUser)
	// login User Route
	r.POST("/login", controller.Login)
	// Run the server go run main.go
	r.Run("localhost: " + PORT)
}
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
