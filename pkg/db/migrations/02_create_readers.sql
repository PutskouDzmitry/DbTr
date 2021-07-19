-- +goose Up
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

-- +goose Down
DROP TABLE readers
