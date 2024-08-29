package main

import (
	"api_geteway/api/router"
	_ "api_geteway/cmd/docs"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
)

// @title           Booking Hotel API
// @version         2.0
// @description     This is an API for booking Hotels.
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @host localhost:8085
// @BasePath /
func main() {
	http.Handle("/swagger/", httpSwagger.WrapHandler)
	router.NewRouter()
}
