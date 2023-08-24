CREATE TABLE order_table (
    order_id INT REFERENCES orders(id),
    table_id INT REFERENCES tables(id),
    PRIMARY KEY(order_id, table_id),
    created_at TIMESTAMP DEFAULT NOW()
);