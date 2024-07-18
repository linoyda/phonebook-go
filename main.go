package main

import (
    "log"
    "os"
	
    "github.com/gin-gonic/gin"
	
    "phonebook/config"
    "phonebook/handlers"
)

func main() {
    config.ConnectDatabase()
	
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
	
    r := gin.Default()
	
    // Trust localhost proxy requests only
    r.ForwardedByClientIP = true
    r.SetTrustedProxies([]string{"127.0.0.1"})

    r.GET("/contacts", handlers.GetContacts)
    r.GET("/contacts/search", handlers.SearchContacts)
    r.POST("/contacts", handlers.AddContact)
    r.PUT("/contacts/:id", handlers.EditContact)
    r.DELETE("/contacts/:id", handlers.DeleteContact)

    err := r.Run(":" + port)
    if err != nil {
    	log.Fatal("FATAL: Failed to start server with error: ", err)
    }
}
