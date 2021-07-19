-- +goose Up
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

-- +goose Down
DROP TABLE cities_of_people