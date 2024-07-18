package handlers

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

// Todo
func GetContacts(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message": "GetContacts - TBD"})
}

// Todo
func SearchContacts(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message": "SearchContacts - TBD"})
}

// Todo
func AddContact(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message": "AddContact - TBD"})
}

// Todo
func EditContact(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message": "EditContact - TBD"})
}

// Todo
func DeleteContact(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message": "DeleteContact - TBD"})
}
