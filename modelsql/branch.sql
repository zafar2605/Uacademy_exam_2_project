CREATE TABLE branch(
    id UUID NOT NULL PRIMARY KEY,
    name VARCHAR(32),
    phone VARCHAR(32),
    photo VARCHAR(512),
    work_start_hour TIME,
    work_end_hour TIME,
    address VARCHAR(64),
    deliver_price NUMERIC DEFAULT 10000,
    active VARCHAR(12),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);
