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

CREATE TABLE "readers"
(
    reader_id INT NOT NULL
        CONSTRAINT readers_pk
            PRIMARY KEY ,
    occupation_id INT NOT NULL,
    city_id INT NOT NULL,
    reader_name CHAR (50) NOT NULL,
    reader_surname CHAR (50) NOT NULL,
    debtor BOOLEAN NOT NULL
);
INSERT INTO "readers" VALUES (1, 1, 1, 'Dima', 'Putkou', FALSE);
INSERT INTO "readers" VALUES (2, 2, 2, 'Ivan', 'Ivanov', TRUE);
INSERT INTO "readers" VALUES (3, 2, 1, 'Danik', 'Domaskanou', TRUE);
INSERT INTO "readers" VALUES (4, 5, 3, 'Artem', 'Menshikou', TRUE);
INSERT INTO "readers" VALUES (5, 4, 5, 'Nikita', 'Miladouski', TRUE);

CREATE TABLE "authors"
(
    author_id INT NOT NULL
        CONSTRAINT authors_pk
            PRIMARY KEY ,
    name_of_author CHAR(50) NOT NULL,
    surname CHAR(50) NOT NULL
);
INSERT INTO "authors" VALUES (1, 'Alexander','Pushkin');
INSERT INTO "authors" VALUES (2, 'Lev','Tolstoy');
INSERT INTO "authors" VALUES (3, 'Alexander','Blok');
INSERT INTO "authors" VALUES (4, 'Nikolay','Nekrasov');
INSERT INTO "authors" VALUES (5, 'Anton','Chekhov');

CREATE TABLE "occupation"
(
    occupation_id INT NOT NULL
        CONSTRAINT occupation_pk
            PRIMARY KEY ,
    name_of_occupation CHAR(50) NOT NULL
);
INSERT INTO "occupation" VALUES (1, 'student');
INSERT INTO "occupation" VALUES (2, 'schoolboy');
INSERT INTO "occupation" VALUES (3, 'worker');
INSERT INTO "occupation" VALUES (4, 'pensioner');
INSERT INTO "occupation" VALUES (5, 'vip');

CREATE TABLE "cities_of_people"
(
    city_id INT NOT NULL
        CONSTRAINT "cities_of_people_pk"
            PRIMARY KEY,
    name_of_city CHAR(50) NOT NULL
);
INSERT INTO "cities_of_people" VALUES (1, 'Gomel');
INSERT INTO "cities_of_people" VALUES (2, 'Minsk');
INSERT INTO "cities_of_people" VALUES (3, 'Vitebsk');
INSERT INTO "cities_of_people" VALUES (4, 'Moscow');
INSERT INTO "cities_of_people" VALUES (5, 'Peter');

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

CREATE TABLE "publishers"
(
    publisher_id INT NOT NULL
        CONSTRAINT publishers_pk
            PRIMARY KEY ,
    city_id INT NOT NULL,
    name_of_publisher CHAR(50) NOT NULL
);
INSERT INTO "publishers" VALUES (1, 1, 'Moscow');
INSERT INTO "publishers" VALUES (2, 2, 'Vitebsk');
INSERT INTO "publishers" VALUES (3, 3, 'Gomel');
INSERT INTO "publishers" VALUES (4, 4, 'Peter');
INSERT INTO "publishers" VALUES (5, 5, 'Peter');

CREATE TABLE "cities of publishers"
(
    city_id INT NOT NULL
        CONSTRAINT "cities of publishers_pk"
            PRIMARY KEY ,
    name_of_city CHAR(50) NOT NULL
);
INSERT INTO "cities of publishers" VALUES (1, 'Moscow');
INSERT INTO "cities of publishers" VALUES (2, 'Vitebsk');
INSERT INTO "cities of publishers" VALUES (3, 'Gomel');
INSERT INTO "cities of publishers" VALUES (4, 'Peter');
INSERT INTO "cities of publishers" VALUES (5, 'Grodno');