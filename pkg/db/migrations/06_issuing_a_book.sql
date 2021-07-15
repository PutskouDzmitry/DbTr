-- +goose Up
CREATE TABLE "issuing a book"
(
    issuing_id INT NOT NULL
        CONSTRAINT "issuing a book_pk"
            PRIMARY KEY,
    reader_id INT NOT NULL,
    book_id INT NOT NULL,
    issue_date_of_the_book DATE NOT NULL,
    book_return_date DATE NOT NULL
);
INSERT INTO "issuing a book" VALUES (1, 1, 1, '2018-05-16', '2018-05-16');
INSERT INTO "issuing a book" VALUES (2, 2, 2, '2018-05-16', '2018-05-16');
INSERT INTO "issuing a book" VALUES (3, 3, 2, '2018-05-16', '2018-05-16');
INSERT INTO "issuing a book" VALUES (4, 1, 3, '2018-05-16', '2018-05-16');
INSERT INTO "issuing a book" VALUES (5, 4, 5, '2018-05-16', '2018-05-16');

-- +goose Down
DROP TABLE "issuing a book"