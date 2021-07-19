-- +goose Up
CREATE TABLE "books"
(
    book_id INT NOT NULL,
    author_id INT NOT NULL,
    publisher_id INT NOT NULL,
    name_of_book CHAR(50) NOT NULL,
    year_of_publication date NOT NULL,
    book_volume INT NOT NULL,
    number INT NOT NULL
);
CREATE UNIQUE index books_book_id_uindex
    ON "books" (book_id);
ALTER TABLE "books"
    ADD CONSTRAINT books_pk
        PRIMARY KEY(book_id);
INSERT INTO "books" VALUES (1,1,1, 'Van Helsing', '2018-05-16',10, 4);
INSERT INTO "books" VALUES (2,2,4, 'Romeo and Juliet', '2018-05-16',10, 4);
INSERT INTO "books" VALUES (3,5,2, 'Three Musketeers', '2018-05-16',10, 4);
INSERT INTO "books" VALUES (4,3,1, 'Captains daughter', '2018-05-16',10, 4);
INSERT INTO "books" VALUES (5,4,3, 'Dubrovsky', '2018-05-16',10, 4);

-- +goose Down
DROP TABLE books
