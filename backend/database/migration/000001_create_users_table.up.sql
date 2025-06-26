CREATE TABLE IF NOT EXISTS users (
    user_id CHAR(32) PRIMARY KEY,
    session_id CHAR(32),
    email VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(100) NOT NULL,
    username VARCHAR(100) NOT NULL UNIQUE,
    firstname VARCHAR(100) NOT NULL,
    lastname VARCHAR(100) NOT NULL,
    birthdate DATETIME NOT NULL,
    role VARCHAR(50) NOT NULL,
    avatar TEXT,
    creation_date DATETIME DEFAULT CURRENT_TIMESTAMP,
    description VARCHAR(255)
);