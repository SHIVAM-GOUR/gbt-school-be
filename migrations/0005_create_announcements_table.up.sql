CREATE TABLE announcements (
    id SERIAL PRIMARY KEY NOT NULL,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    published_date DATE,
    start_date DATE,
    end_date DATE,
    audience TEXT,
    attachment_url TEXT,
    is_active BOOLEAN,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);