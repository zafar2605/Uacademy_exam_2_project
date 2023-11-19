CREATE TABLE category (
    id UUID NOT NULL PRIMARY KEY,
    title VARCHAR(64) NOT NULL,
    image VARCHAR(512),
    parent_id UUID REFERENCES category (id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);
