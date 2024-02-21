CREATE SCHEMA IF NOT EXISTS public;

CREATE TABLE IF NOT EXISTS "Clue_Types" (
    name TEXT NOT NULL PRIMARY KEY,
    description TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS "Indicator_Words" (
    word TEXT NOT NULL PRIMARY KEY,
    clue_type TEXT NOT NULL REFERENCES "Clue_Types"(name),
    hint TEXT
);

ALTER TABLE "Clue_Types" ADD COLUMN example TEXT;