CREATE DATABASE qrcode_app;

-- User Table
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(64)  NOT NULL,
    last_name VARCHAR(64)  NOT NULL,
    email TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL,
    role TEXT DEFAULT 'user',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


-- QrCode Table
CREATE TABLE qr_codes (
    id SERIAL PRIMARY KEY,
    token TEXT UNIQUE,
    url TEXT,
    image TEXT,
    valid boolean,
    created_at TIMESTAMP DEFAULT NOW()
);

