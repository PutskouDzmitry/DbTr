package const_db

//Const for queries
const (
	Read                   = `SELECT books.book_id, books.name_of_book ,publishers.name_of_publisher FROM "publishers" RIGHT JOIN books on books.publisher_id=publishers.publisher_id WHERE "book_id" = $1`
	ReadBookWithJoin       = "RIGHT JOIN books on books.publisher_id=publishers.publisher_id"
	SelectBookAndPublisher = "books.book_id, books.name_of_book ,publishers.name_of_publisher"
	BookId                 = "book_id = ?"
	InsertBook             = `INSERT INTO "books" ("book_id","author_id","publisher_id","name_of_book","year_of_publication","book_volume","number") VALUES ($1,$2,$3,$4,$5,$6,$7)`
	SelectAllBooks         = `SELECT * FROM "books"`
	Update                 = `UPDATE "books" SET "name_of_book"=$1 WHERE book_id = $2`
	Delete                 = `DELETE FROM "books" WHERE book_id = $1`
	SelectFromBooksWithID  = `SELECT books.book_id, books.name_of_book ,publishers.name_of_publisher FROM "publishers" RIGHT JOIN books on books.publisher_id=publishers.publisher_id WHERE "book_id" = $1`
)
