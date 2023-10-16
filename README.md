# books-api-challenge

## Tasks
1. Design and code a simple books API that exposes the following endpoints (Use
   storage/database of your choice):
   ● Description: Saves/updates the given book name and release date in the database.
   Request: PUT /books/<bookname> { “releaseDate”: “DD-MM-YYYY” }
   Response: 204 No Content
   Note: DD-MM-YYYY must be a date before the today date.
   ● Description: Gets all books ordered by release date.
   Request: GET /books?order=(asc|desc)
   Response: 200 OK
2. Using the previous base code, write the same functionality with a lambda.
3. Produce 2 system diagrams of your solutions deployed to AWS. The first one running on a
   EC2 and the second one with the lambda.
   NOTE: The solution must have tests and must be runnable locally.


### Technical details
- [Go 1.21.0](https://go.dev/): Go version 1.21.0
- [go-chi](https://github.com/go-chi/chi): lightweight router

### Requirements
- Go version 1.21.0

### How to run the application:
Clone the repository and move to the project folder:
```
git clone https://github.com/juanmabaracat/books-challenge.git
cd books-challenge
```
Run the application:
```
go run cmd/main.go
```
Run all tests (root folder):
```
go test ./...
```

## Examples
### Update a book
### PUT `/books/The%20Alchemist`

| Code | Description  |
|------|--------------|
| 201  | book created |
| 204  | book updated |
| 400  | Bad request  |
| 500  | Server error |

```
curl --location --request PUT 'localhost:8080/books/The%20Alchemist' \
--header 'Content-Type: application/json' \
--data '{"release_date": "15-10-1988"}'
```

#### Request body example
```
{"release_date": "15-10-1988"}
```

### List books
### PUT `/books/`

| Code | Description  |
|------|--------------|
| 200  | books list   |
| 400  | Bad request  |
| 500  | Server error |

```
curl --location 'localhost:8080/books?order=desc'
```

#### Response body example
```
[
    {
        "Name": "Sapiens",
        "ReleaseDate": "04-03-2014"
    },
    {
        "Name": "The Alchemist",
        "ReleaseDate": "15-10-1988"
    }
]
```