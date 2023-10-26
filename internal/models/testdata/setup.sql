-- Create a new database called 'DatabaseName'
CREATE TABLE snippets (
    id INTEGER NOT NULL PRIMARY KEY,
    title VARCHAR NOT NULL,
    content TEXT NOT NULL,
    created TIMESTAMP NOT NULL,
    expires TIMESTAMP NOT NULL
);

CREATE INDEX idx_snippets_created ON snippets(created);

CREATE TABLE users (
    id INTEGER NOT NULL PRIMARY KEY,
    name VARCHAR NOT NULL,
    email VARCHAR NOT NULL,
    hashed_pasword VARCHAR NOT NULL,
    created TIMESTAMP NOT NULL
);

ALTER TABLE users ADD CONSTRAINT users_uc_email UNIQUE(email);

INSERT INTO users (id, name, email, hashed_pasword, created) VALUES (
    1,
    'nicolas',
    'ggdnicolas@gmail.com',
    'password',
    '2023-01-01 10:00:00'
);