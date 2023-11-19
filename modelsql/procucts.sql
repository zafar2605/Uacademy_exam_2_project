CREATE TABLE products(
    id VARCHAR(10) PRIMARY KEY,
    title VARCHAR,
    description VARCHAR,
    photos VARCHAR(512),
    price NUMERIC,
    category_id UUID REFERENCES category (id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);