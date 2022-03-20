package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/huyanhvn/quick-rest-go/controller"
	"github.com/huyanhvn/quick-rest-go/docs"
	"github.com/huyanhvn/quick-rest-go/model"
	swaggerfiles "github.com/swaggo/files"
	ginswagger "github.com/swaggo/gin-swagger"
	ginprometheus "github.com/zsais/go-gin-prometheus"
)

// @title           quick-rest
// @version         1.0
// @description     Boilerplate for a Go REST web service.

// @contact.name   Huy Nguyen

// @host      localhost:8080
// @BasePath  /api/v1

var port = flag.Int("port", 8080, "Port to listen for HTTP")

func main() {
	flag.Parse()
	// initialize connection to local database
	model.ConnectDatabase("test.db")

	// setup routes
	r := gin.Default()
	v1 := r.Group("/api/v1")
	docs.SwaggerInfo.BasePath = "/api/v1"
	c := controller.NewController()
	v1.GET("/items", c.GetAllItems)
	v1.GET("/items/:id", c.GetItemByID)
	v1.PATCH("/items/:id", c.UpdateItem)
	v1.DELETE("/items/:id", c.DeleteItem)
	v1.POST("/items", c.CreateItem)

	// setup swagger
	r.GET("/swagger/*any", ginswagger.WrapHandler(swaggerfiles.Handler))

	// setup Prometheus exporter
	p := ginprometheus.NewPrometheus("gin")
	p.Use(r)

	// start server
	log.Printf("Listening on :%d", *port)
	r.Run(fmt.Sprintf(":%d", *port))
}
