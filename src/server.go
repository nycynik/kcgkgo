package main

import (
	"io/ioutil"
	"os"
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"controllers"
)

func main() {

	// Print banner
	dat, _ := ioutil.ReadFile("./src/config/banner.txt")
	fmt.Print(string(dat))

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	// Router
	controllers.APISetup(e)
	controllers.LiveOpsSetup(e)

	// Port
	port := os.Getenv("PORT") //Get port from .env file, we did not specify any port so this should return an empty string when tested locally
	if port == "" {
		port = ":8000" //localhost
	}

	// CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"X-Requested-With", "Content-Type", "Authorization"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	e.Logger.Fatal(e.Start(port))

}