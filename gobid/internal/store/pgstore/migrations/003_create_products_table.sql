-- Write your migrate up statements here
CREATE TABLE IF NOT EXISTS products (
    id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),

    seller_id UUID NOT NULL REFERENCES users (id),
    product_name TEXT NOT NULL,
    description TEXT NOT NULL,
    baseprice FLOAT NOT NULL,
    auction_end TIMESTAMPTZ NOT NULL,
    is_sold BOOLEAN NOT NULL DEFAULT false
);

---- create above / drop below ----
DROP TABLE IF EXISTS products;

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
