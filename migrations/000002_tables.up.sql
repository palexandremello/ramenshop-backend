CREATE TABLE IF NOT EXISTS tables (
    id BIGSERIAL primary key,
    capacity  int not null,
    is_available varchar(255) not null,
    created_at TIMESTAMP default now()
);