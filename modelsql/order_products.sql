CREATE TABLE order_products(
    id UUID NOT NULL PRIMARY KEY,
    order_id VARCHAR(12) REFERENCES orders (id),
    discount_type VARCHAR(12),
    discount_amount NUMERIC,
    product_id VARCHAR(12) REFERENCES products (id),
    quantity INT,
    price NUMERIC,
    sum NUMERIC,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);


CREATE OR REPLACE FUNCTION order_products_function() RETURNS TRIGGER LANGUAGE PLPGSQL
    AS
$$
    DECLARE
        p_price NUMERIC;
        d_amount NUMERIC = 0;
    BEGIN
        SELECT price FROM products INTO p_price WHERE id = new.product_id;

        IF new.discount_type = 'fix' THEN
            d_amount = new.quantity * p_price - 10000;
        ELSIF new.discount_type = 'percentage' THEN
            d_amount = new.quantity * p_price - (new.quantity * p_price)/10;
        END IF;

        UPDATE order_products SET
        price = p_price,
        sum = new.quantity * p_price,
        discount_amount = d_amount
        WHERE id = new.id;

        RETURN new;
    END;
$$;

CREATE TRIGGER order_products_trigger
AFTER INSERT ON order_products
FOR EACH ROW EXECUTE PROCEDURE
order_products_function();