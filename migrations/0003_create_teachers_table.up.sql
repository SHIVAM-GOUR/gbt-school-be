CREATE TABLE teachers (
    id SERIAL PRIMARY KEY NOT NULL,
    first_name TEXT NOT NULL,
    last_name TEXT,
    email TEXT,
    phone_number TEXT,
    date_of_birth DATE,
    gender TEXT,
    hire_date DATE,
    address TEXT
);