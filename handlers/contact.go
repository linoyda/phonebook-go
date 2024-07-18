package handlers

import (
    "context"
	"log"
    "strconv"
    "time"
    "net/http"
	
    "phonebook/config"
    "phonebook/models"
	
    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo/options"
)

// Handles a GET request to return a list of contacts from contacts, based on pagination and limit parameters.
func GetContacts(c *gin.Context) {
    var contacts []models.Contact

    // Parse and validate query parameters
    limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
    if err != nil || limit <= 0 {
        c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid limit parameter"})
        return
    }
    page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
    if err != nil || page <= 0 {
        c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid page parameter"})
        return
    }
	
    offset := (page - 1) * limit
    collection := config.DB.Collection("contacts")
    ctx, cancel := context.WithTimeout(c, 10*time.Second)
    defer cancel()

    // Set up options for pagination
    opts := options.Find().SetLimit(int64(limit)).SetSkip(int64(offset))
    cursor, err := collection.Find(ctx, bson.M{}, opts)
    if err != nil {
        log.Println("Error fetching contacts:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch contacts"})
        return
    }
    defer cursor.Close(ctx)

    // Iterate through the cursor and decode contacts
    for cursor.Next(ctx) {
        var contact models.Contact
        if err := cursor.Decode(&contact); err != nil {
            log.Println("Error decoding contact:", err)
            c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to decode contact"})
            return
        }
        contacts = append(contacts, contact)
    }

    // Check for cursor errors
    if err := cursor.Err(); err != nil {
        log.Println("Cursor error:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"message": "Cursor error"})
        return
    }

    // Respond with the list of contacts
    c.JSON(http.StatusOK, gin.H{"message": "Contacts fetched successfully", "data": contacts})
}

// Todo
func SearchContacts(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message": "SearchContacts - TBD"})
}

// Handles a POST request for adding a new contact.
func AddContact(c *gin.Context) {
    var contact models.Contact

    // Bind the request payload to a Contact struct
    if err := c.ShouldBindJSON(&contact); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input"})
        return
    }

    // Generate a new ID for the contact
    contact.ID = primitive.NewObjectID()

    collection := config.DB.Collection("contacts")
    ctx, cancel := context.WithTimeout(c, 10*time.Second)
    defer cancel()

    // Add the new contact
    _, err := collection.InsertOne(ctx, contact)
    if err != nil {
        log.Println("Error adding contact:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to add contact"})
        return
    }

    // Respond with success and the added contact
    c.JSON(http.StatusCreated, gin.H{"message": "Contact added successfully", "data": contact})
}

// Todo
func EditContact(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message": "EditContact - TBD"})
}

// Todo
func DeleteContact(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message": "DeleteContact - TBD"})
}
