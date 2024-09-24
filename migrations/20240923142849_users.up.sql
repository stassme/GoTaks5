CREATE TABLE users (
                      id SERIAL PRIMARY KEY,
                      email VARCHAR(255) NOT NULL,
                      password VARCHAR(255) NOT NULL,
                      deleted_at TIMESTAMP,
                      created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                      updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);