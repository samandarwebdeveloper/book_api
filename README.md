# Books API in Go with Fiber and Docker Microservice


## Docker
Clone this repository and run:
```
docker-compose up
```

You can then hit the following endpoints:

| Method | Route         | Body                                                                                              |
| ------ | ------------- | ------------------------------------------------------------------------------------------------- |
| GET    | /books        |                                                                                                   |
| GET    | /books/:id    |                                                                                                   |
| POST   | /books        | `{"title": "Atomic habits", "author": "James Clear", "stock": 10, "price": 99, "category_id": 2 }`|
| DELETE | /books/:id    |                                                                                                   |
| PUT    | /books/:id    | `{"title": "Atomic habits", "author": "James Clear", "stock": 10, "price": 99, "category_id": 2 }`|

