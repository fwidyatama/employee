CREATE TABLE employees
(
    id              SERIAL PRIMARY KEY,
    first_name      TEXT NOT NULL,
    last_name       TEXT,
    email           TEXT  NOT NULL,
    hire_date       DATE NOT NULL
);