

CREATE TABLE IF NOT EXISTS messages (
    id SERIAL PRIMARY KEY,
    message JSONB NOT NULL,
    transaction_date TIMESTAMP DEFAULT current_timestamp,
    app_id INT NOT NULL,
    category_id INT NOT NULL,
    transaction_id VARCHAR(255),
    merchant VARCHAR(255),
    CONSTRAINT fk_messages_app_id FOREIGN KEY (app_id) REFERENCES apps(id),
    CONSTRAINT fk_messages_category_id FOREIGN KEY (category_id) REFERENCES categories(id)
);
