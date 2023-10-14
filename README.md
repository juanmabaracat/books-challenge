# books-challenge

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