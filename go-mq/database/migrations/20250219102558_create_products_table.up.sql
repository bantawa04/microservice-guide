CREATE TABLE IF NOT EXISTS products
(
    id          UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    name        VARCHAR(100) NOT NULL,
    description TEXT NULL,
    price       DECIMAL(10,2) NOT NULL,
    user_id     UUID NOT NULL,
    created_at  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP,
    deleted_at  TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE INDEX idx_products_user_id ON products(user_id);