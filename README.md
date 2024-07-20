# Phonebook Web Server API in Go
A phonebook application in Go, MongoDB and Docker. It provides an API for managing contacts, including creating, reading, updating, and deleting contact entries. It also includes some tests for adding, searching and deleting a contact. The tests are built using Docker as well.

## Prerequisites
1. Docker and Docker Compose
2. Curl for API testing (alternatively, you can use Postman)
3. Go 1.22.5 (for running tests manually)


## How to Setup
1. Clone the repository
2. **Server Setup:** From an elevated terminal directed to the repository's directory, build and start the application server and MongoDB containers. Test container will be built and run as well, then exit.
    ```sh
    docker-compose up --build
    ```
    The application will be available at http://localhost:8080.
3. **Client Setup:** From an elevated terminal, test the API endpoints below, using Curl, for example.


## API Endpoints

#### 1) Get Contacts
* **Endpoint**: `GET /contacts`
* **Query Parameters**: 
    1. `limit` - The maximum contact amount to retreive, an integer between 1 and 10.
    2. `pages`  - The amount of pages to display the contacts in.
* **Description**: 
         Fetch a list of contacts. Note that if the total users fetched (within the limit) is larger than the amount of pages provided, each contact will be displayed in a page.
* **Examples**:
    ```sh
    curl -X GET "http://localhost:8080/contacts?limit=10&pages=1"
    ```
    Should return up to 10 existing contacts in one page structure.
    
    ```sh
    curl -X GET "http://localhost:8080/contacts?limit=14&pages=1"
    ```
    Should result in an error due to an invallid limit. 
    
#### Search Contacts
* **Endpoint**: `GET /contacts/search`
* **Query Parameters**: 
   1. `q` - Query to be searched within the database. Query could be of first or last name, phone number, etc.
   2. `limit` - The maximum contact amount to retreive, an integer between 1 and 10.
* **Description**: Search for contacts based on a query. The search is optimized due to indices created on the DB.
* **Examples**:
    ```sh
    curl -X GET "http://localhost:8080/contacts/search?q=Dana&limit=10"
    ```
    
#### Add Contact
* **Endpoint**: `POST /contacts`
* **Description**: Add a contact to the phonebook. All fields are neccessary.
* **Request Body**:
    ```
    {
      "first_name": "Tomer",
      "last_name": "Chen",
      "phone": "012332101",
      "address": "6 Dizingoff St"
    }
    ```
* **Examples**:
    ```
    curl -X POST http://localhost:8080/contacts -H "Content-Type: application/json" -d "{\"first_name\": \"Idan\", \"last_name\": \"David\", \"phone\": \"12121212\", \"address\": \"111 Dror St\"}"
    ```
    Should result in a success.
    
    ```
    curl -X POST http://localhost:8080/contacts -H "Content-Type: application/json" -d "{\"first_name\": \"Idan\", \"last_name\": \"David\", \"address\": \"111 Dror St\"}"
    ```
    Will fail due to missing phone number field.

#### Edit Contact
* **Endpoint**: `PUT /contacts/:id`
* **Description**: Edit an existing contact based on its ID. All fields are neccessary. If the contact doesn't exist, or if the given ID is in a wrong format, an error will be returned.
* **Request Body**:
    ```
    {
      "first_name": "Jane",
      "last_name": "Rojers",
      "phone": "0987654321",
      "address": "456 Marganit St"
    }
    ```
* **Examples**:
    ```
    curl -X PUT http://localhost:8080/contacts/<CONTACT_ID> -H "Content-Type: application/json" -d "{\"first_name\": \"Jane\", \"last_name\": \"Doe\", \"phone\": \"9876543210\", \"address\": \"456 Mainss St\"}
    ```
    
#### Delete Contact
* **Endpoint**: `DELETE /contacts/:id`
* **Description**: Delete an existing contact based on its ID. If the contact doesn't exist, or if the given ID is in a wrong format, an error will be returned.
* **Examples**:
    ```
    curl -X DELETE http://localhost:8080/contacts/<CONTACT_ID>
    ```
## Tests
**Test Suite**: 
located in the tests/ directory. To run the tests, build and execute the test container, which is done automatically when you build the entire setup. When tests are done, the test container is exited.
The tests cover the following scenarios:
*  **Search Contact**
* **Add Contact**
* **Delete Contact**

## Additional Information and General Notes
* **Database Initialization**: The database and collections are initialized automatically on startup.
* **Logs**: Application logs can be viewed using docker-compose logs.
* For more details on each component, refer to the documentation: 
  * [Gin]
  * [MongoDB]
  * [Docker]
  * [Docker Compose]


   [MongoDB]: <https://www.mongodb.com/docs/drivers/go/current/fundamentals/>
   [Docker]: <https://docs.docker.com/>
   [Gin]: <https://gin-gonic.com/>
   [Docker Compose]: <https://docs.docker.com/compose/>
   
