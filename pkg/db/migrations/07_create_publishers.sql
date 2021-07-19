-- +goose Up
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

-- +goose Down
DROP TABLE publishers