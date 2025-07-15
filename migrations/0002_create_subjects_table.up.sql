CREATE TABLE subjects (
    id SERIAL PRIMARY KEY NOT NULL,
    name TEXT NOT NULL,
    class_id INTEGER NOT NULL,
    FOREIGN KEY (class_id) REFERENCES classes(id)
);