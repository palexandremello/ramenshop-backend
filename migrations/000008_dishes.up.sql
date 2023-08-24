CREATE TABLE dishes (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    DESCRIPTION text,
    photo_id INT references photos(id),
    price DECIMAL(10, 2) NOT NULL,
    available BOOLEAN DEFAULT TRUE,
    type dishtype NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);