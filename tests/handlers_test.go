package tests

import (
    "bytes"
    "context"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "phonebook/config"
    "phonebook/handlers"
    "phonebook/models"
    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "testing"
)

func setupRouter() *gin.Engine {
    r := gin.Default()
    r.GET("/contacts/search", handlers.SearchContacts)
    r.POST("/contacts", handlers.AddContact)
    r.DELETE("/contacts/:id", handlers.DeleteContact)
    return r
}

func setupTestDB() *mongo.Collection {
    config.ConnectDatabase()
    return config.DB.Collection("contacts")
}

// Tests the GET /contacts/search endpoint for searching contacts.
func TestSearchContacts(t *testing.T) {
    router := setupRouter()
    collection := setupTestDB()

    contact := models.Contact{
        FirstName: "Yael",
        LastName:  "Levi",
        Phone:     "2222222222",
        Address:   "1 Histadrut St",
    }
    _, err := collection.InsertOne(context.Background(), contact)
    if err != nil {
        t.Fatalf("Failed to insert test contact: %v", err)
    }

    req, _ := http.NewRequest("GET", "/contacts/search?q=Levi&limit=10", nil)
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)
}

// Tests the POST /contacts endpoint for adding a new contact.
func TestAddContact(t *testing.T) {
    router := setupRouter()
    collection := setupTestDB()

    contact := models.Contact{
        FirstName: "Tomer",
        LastName:  "Chen",
        Phone:     "1010101010",
        Address:   "6 Dizingoff St",
    }
    contactJSON, _ := json.Marshal(contact)

    req, _ := http.NewRequest("POST", "/contacts", bytes.NewBuffer(contactJSON))
    req.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusCreated, w.Code)

    // Verify the contact was added
    var result models.Contact
    err := collection.FindOne(context.Background(), bson.M{"phone": contact.Phone}).Decode(&result)
    if err != nil {
        t.Fatalf("Failed to fetch added contact: %v", err)
    }

    assert.Equal(t, contact.FirstName, result.FirstName)
    assert.Equal(t, contact.LastName, result.LastName)
    assert.Equal(t, contact.Phone, result.Phone)
    assert.Equal(t, contact.Address, result.Address)
}

// Tests the DELETE /contacts/:id endpoint for deleting a contact.
func TestDeleteContact(t *testing.T) {
    router := setupRouter()
    collection := setupTestDB()

    // Setup test data
    contact := models.Contact{
        FirstName: "John",
        LastName:  "Doe",
        Phone:     "1234567890",
        Address:   "123 Main St",
    }
    contact.ID = primitive.NewObjectID()
    
    _, err := collection.InsertOne(context.Background(), contact)
    if err != nil {
        t.Fatalf("Failed to insert test contact: %v", err)
    }

    req, _ := http.NewRequest("DELETE", "/contacts/"+contact.ID.Hex(), nil)
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)

    // Verify the contact was deleted
    count, err := collection.CountDocuments(context.Background(), bson.M{"_id": contact.ID})
    if err != nil {
        t.Fatalf("Failed to count documents after delete: %v", err)
    }

    assert.Equal(t, int64(0), count)
}

