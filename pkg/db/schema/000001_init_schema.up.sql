CREATE TABLE IF NOT EXISTS "author"(
    "author_id" bigserial NOT NULL PRIMARY KEY,
    "full_name" varchar NOT NULL,
    "nick" varchar NOT NULL,
    "speciality" varchar NOT NULL
);

CREATE TABLE IF NOT EXISTS "book"(
    "book_id" bigserial PRIMARY KEY,
    "author_id" bigint not null,
    "name" varchar NOT NULL DEFAULT '',
    "genre" varchar not null DEFAULT '',
    "code_isbn" varchar NOT NULL
);

CREATE TABLE IF NOT EXISTS "reader"(
	"reader_id" SERIAL PRIMARY KEY,
	"full_name" VARCHAR NOT NULL
);

CREATE TABLE IF NOT EXISTS "reader_book"(
	"reader_id" INT NOT NULL,
	"book_id" INT NOT NULL
);

ALTER TABLE "book" ADD FOREIGN KEY ("author_id") REFERENCES "author" ("author_id");
ALTER TABLE "reader_book" ADD FOREIGN KEY("reader_id") REFERENCES "reader" ("reader_id");
ALTER TABLE "reader_book" ADD FOREIGN KEY("book_id") REFERENCES "book" ("book_id");