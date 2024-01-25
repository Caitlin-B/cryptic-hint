CREATE SCHEMA IF NOT EXISTS public;

CREATE TABLE IF NOT EXIST "Clue_Types" (
    id SERIAL NOT NULL PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT NOT NULL
);

CREATE TABLE IF NOT EXIST "Indicator_Words" (
    id SERIAL NOT NULL PRIMARY KEY,
    word TEXT NOT NULL,
    clue_type TEXT NOT NULL REFERENCES "Clue_Types"(name),
    description TEXT NOT NULL REFERENCES "Clue_Types"(description)
);