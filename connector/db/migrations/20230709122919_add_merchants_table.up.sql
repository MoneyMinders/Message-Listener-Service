
CREATE TABLE IF NOT EXISTS merchants (
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    category_id INT NOT NULL,
    CONSTRAINT fk_merchants_category_id FOREIGN KEY (category_id) REFERENCES categories(id)
);
