Just a simple REST Go Apps with Mux

go run main.go in the folder and you get several endpoints
1. GET http://localhost:8000/api/books --> Find All
2. GET http://localhost:8000/api/books/1 --> Find One
3. POST http://localhost:8000/api/books --> Create
   Payload
   {
        "isbn" : "232342324",
        "title" : "Java Programming",
        "author" : {
            "firstname" : "Carol",
            "lastname" : "William"
        }
    }
4. PUT http://localhost:8000/api/books/2 --> Update
   Payload
   {
        "isbn" : "123456",
        "title" : "Title",
        "author" : {
            "firstname" : "Firstname",
            "lastname" : "Lastname"
        }
    }
5. DELETE http://localhost:8000/api/books/2 --> Delete