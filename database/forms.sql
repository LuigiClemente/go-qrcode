CREATE DATABASE qrcode_app;

-- User Table
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(64) NOT NULL,
    password VARCHAR(128) NOT NULL,
    email VARCHAR(128) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

-- QrCode Table
CREATE TABLE qr_codes (
    id SERIAL PRIMARY KEY,
    token TEXT,
    url TEXT,
    image TEXT,
    valid boolean,
    created_at TIMESTAMP DEFAULT NOW()
);

