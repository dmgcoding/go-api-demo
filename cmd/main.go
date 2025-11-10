package main

import (
	"api_demo/internal/routes"
	"log"
	"net/http"
)

func main(){
	// COMMENTED BELOW LINES SO THAT IT IS EASIER TO SETUP/RUN
	// godotenv.Load()

	// portString := os.Getenv("PORT")
	// if portString == "" {
	// 	log.Fatal("port string is not provided.\n")
	// }
	
	portString := "8181"

	srv := &http.Server{
		Addr:    ":" + portString,
		Handler: routes.NewRouter(),
	}

	log.Printf("listening at port: %s", portString)

	srv.ListenAndServe()

}