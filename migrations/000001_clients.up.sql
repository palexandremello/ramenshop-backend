CREATE TABLE IF NOT EXISTS clients (
    id BIGSERIAL primary key,
    name TEXT not null,
    gender int not null,
    age int not null,
    email VARCHAR(259) null,
    created_at TIMESTAMP default now()
);