CREATE TABLE orders(
    id VARCHAR(10) PRIMARY KEY,
    client_id UUID REFERENCES client (id) ON DELETE CASCADE,
    branch_id UUID REFERENCES branch (id) ON DELETE CASCADE,
    address VARCHAR(64),
    delivery_price NUMERIC,
    total_count INT,
    total_price NUMERIC,
    status VARCHAR(12),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);



CREATE OR REPLACE FUNCTION orders_function() RETURNS TRIGGER LANGUAGE PLPGSQL
    AS
$$
    DECLARE
        d_price NUMERIC;
        t_count INT;
        t_price NUMERIC;
        b_id UUID;
    BEGIN
       SELECT branch_id FROM orders INTO b_id WHERE id = new.order_id;
       SELECT deliver_price FROM branch INTO d_price WHERE id = b_id;
       SELECT SUM(sum), SUM(quantity) FROM order_products INTO t_price, t_count WHERE order_id = new.order_id;

       UPDATE orders SET
       delivery_price = d_price,
       total_count = t_count,
       total_price = t_price + d_price
       WHERE id = new.order_id;

        RETURN new;
    END;
$$;

CREATE TRIGGER orders_trigger
AFTER INSERT OR UPDATE ON order_products
FOR EACH ROW EXECUTE PROCEDURE
orders_function();