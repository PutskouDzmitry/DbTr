-- +goose Up
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

-- +goose Down
DROP TABLE occupation