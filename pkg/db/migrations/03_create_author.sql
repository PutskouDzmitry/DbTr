-- +goose Up
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

-- +goose Down
DROP TABLE authors