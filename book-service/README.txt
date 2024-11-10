POST /books
Content-Type: application/json
{
  "title": "Go Programming",
  "author_id": 1,
  "category_id": 1,
  "stock_id": 1,
  "description": "Learn Go programming."
}

Respnse :
{
  "id": 1,
  "title": "Go Programming",
  "author_id": 1,
  "category_id": 1,
  "stock_id": 1,
  "description": "Learn Go programming.",
  "created_at": "2024-11-10T00:00:00Z",
  "updated_at": "2024-11-10T00:00:00Z"
}

GET /books
Response:
[
    {
      "id": 1,
      "title": "Go Programming",
      "author_id": 1,
      "category_id": 1,
      "stock_id": 1,
      "description": "Learn Go programming."
    }
]

PUT /books/1
Content-Type: application/json
{
  "title": "Advanced Go Programming",
  "author_id": 1,
  "category_id": 1,
  "stock_id": 1,
  "description": "Advanced techniques in Go programming."
}

Response:
{
    "id": 1,
    "title": "Advanced Go Programming",
    "author_id": 1,
    "category_id": 1,
    "stock_id": 1,
    "description": "Advanced techniques in Go programming.",
    "updated_at": "2024-11-10T00:00:00Z"
}

DELETE /books/1
HTTP/1.1 204 No Content

POST /books/borrow
Content-Type: application/json

{
  "user_id": 1,
  "book_id": 1,
  "due_date": "2024-11-20"
}

Response :
{
    "id": 1,
    "user_id": 1,
    "book_id": 1,
    "due_date": "2024-11-20",
    "returned": false
}

PUT /books/return/1
{
  "id": 1,
  "user_id": 1,
  "book_id": 1,
  "due_date": "2024-11-20",
  "returned": true
}