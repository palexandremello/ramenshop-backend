CREATE TABLE order_items (
    ID SERIAL PRIMARY KEY,
    order_id INT NOT NULL REFERENCES orders(id),
    dish_id INT NOT NULL REFERENCES dishes(id),
    quantity INT NOT NULL,
    updated_at TIMESTAMP DEFAULT NOW(),
    created_at TIMESTAMP DEFAULT NOW()

);