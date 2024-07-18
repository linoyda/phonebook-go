package main

import (
    "github.com/gin-gonic/gin"
    "phonebook/config"
    "phonebook/handlers"
)

func main() {
    r := gin.Default()
	
    // Trust localhost proxy requests only
    r.ForwardedByClientIP = true
    r.SetTrustedProxies([]string{"127.0.0.1"})

    config.ConnectDatabase()

    r.GET("/contacts", handlers.GetContacts)
    r.GET("/contacts/search", handlers.SearchContacts)
    r.POST("/contacts", handlers.AddContact)
    r.PUT("/contacts/:id", handlers.EditContact)
    r.DELETE("/contacts/:id", handlers.DeleteContact)

    r.Run()
}
