package handlers

import (
    "context"
	"log"
    "strconv"
	"encoding/json"
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

    // Parse and validate query parameters. Limit is the max amount of users to retrieve.
	// totalPages is the amount of pages to display in total. 
	// If contacts / totalPage < 1, we'll show exactly one user per page (less than the totalPages required)
    limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
    if err != nil || limit < 1 || limit > 10{
        c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid limit parameter"})
        return
    }
    totalPages, err := strconv.Atoi(c.DefaultQuery("pages", "1"))
    if err != nil || totalPages < 1 {
        c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid pages parameter"})
        return
    }
	
	collection := config.DB.Collection("contacts")
    ctx, cancel := context.WithTimeout(c, 10*time.Second)
    defer cancel()

	// Fetch all contacts with the given limit
    opts := options.Find().SetLimit(int64(limit))
    cursor, err := collection.Find(ctx, bson.M{}, opts)
    if err != nil {
        log.Println("Error fetching contacts:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch contacts"})
        return
    }
    defer cursor.Close(ctx)

    // Iterate through the cursor and decode contacts into the struct type.
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
	
	// Calculate the number of contacts per page if the contacts fetched are larger than 0.
    totalContacts := len(contacts)
    if totalContacts <= 0 {
        c.JSON(http.StatusOK, gin.H{"data": []models.Contact{}})
        return
    }

    contactsPerPage := (totalContacts + totalPages - 1) / totalPages

    // Split contacts into pages
    pages := make([]map[string]interface{}, 0, totalPages)
    for i := 0; i < totalContacts; i += contactsPerPage {
        end := i + contactsPerPage
        if end > totalContacts {
            end = totalContacts
        }
        page := map[string]interface{}{
            "pageIndex": i / contactsPerPage + 1,
            "contacts":  contacts[i:end],
        }
        pages = append(pages, page)
    }

    response := gin.H{
        "pages":      pages,
        "totalPages": len(pages),
    }

    // Marshal the response to pretty JSON and write the response
    jsonResponse, err := json.MarshalIndent(response, "", "    ")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to format JSON response"})
        return
    }

    c.Writer.Header().Set("Content-Type", "application/json")
    c.Writer.WriteHeader(http.StatusOK)
    c.Writer.Write(jsonResponse)
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
