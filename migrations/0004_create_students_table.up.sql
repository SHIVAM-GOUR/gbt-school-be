CREATE TABLE students (
    student_id SERIAL PRIMARY KEY NOT NULL,
    roll_number TEXT,
    first_name TEXT NOT NULL,
    last_name TEXT,
    date_of_birth DATE,
    gender TEXT,
    email TEXT,
    phone_number TEXT,
    address TEXT,
    enrollment_date DATE,
    class_id INTEGER NOT NULL,
    status TEXT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    FOREIGN KEY (class_id) REFERENCES classes(id)
);