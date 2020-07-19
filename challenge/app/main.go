package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"

	"echotom.dev/hellofresh/router"
	"github.com/joho/godotenv"
)

func main() {
	//must remove the line below before deploying to minikube
	godotenv.Load(".env")
	p := os.Getenv("SERVE_PORT")
	validPort, _ := regexp.MatchString("^\\d{1,5}$", p)
	if validPort {
		log.Fatal(http.ListenAndServe(":"+p, router.Handlers()))
	} else {
		fmt.Println("Invalid port number:", p)
		os.Exit(1)
	}
}
