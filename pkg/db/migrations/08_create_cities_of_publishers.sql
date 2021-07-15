-- +goose Up
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

-- +goose Down
DROP TABLE "cities of publishers"