## Description

Application for get books with such operations as in table below:


|             Path            | Method | Description                           | Body example                                                                                                                                                                                                                     |
|:---------------------------:|--------|---------------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| /books                   | GET    | get all books                      | ```[{"BookId":1,"AuthorId":2,"PublisherId":1,"NameOfBook":"Belka","YearOfPublication":"2020-10-10", "BookVolume":20, "Number":1},{"BookId":2,"AuthorId":1,"PublisherId":4,"NameOfBook":"Strelka","YearOfPublication":"2021-12-21", "BookVolume":220, "Number":11},{"BookId":2,"AuthorId":3,"PublisherId":4,"NameOfBook":"Space","YearOfPublication":"2010-10-10", "BookVolume":202, "Number":11}]``` |
| /books                   | POST   | create new book                    |                                                                                                                                                                                                                                  |
| /books/{id}              | GET    | get book by the id                 | ```{"BookId":1,"AuthorId":2,"PublisherId":1,"NameOfBook":"Belka","YearOfPublication":"2020-10-10", "BookVolume":20, "Number":1}```                                                                                                                                  |
| /books/{id}/{unit_price} | PUT    | update book's price by the id |                                                                                                                                                                                                                                  |
| /books/{id}              | DELETE | delete book by the id              |                                                                                                                                                                                                                                  |

## Usage 
1. Run server on port `8080`
	`go run cmd/main.go`
2.  Open URL
`http://localhost:8080/`

## Usage unit tests
To run unit tests type:
`go test ./...`
