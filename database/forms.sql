CREATE DATABASE qrcode_app;
-- QrCode Table
CREATE TABLE qr_codes (
    id SERIAL PRIMARY KEY,
    token TEXT UNIQUE,
    url TEXT,
    image TEXT,
    valid boolean,
    created_at TIMESTAMP DEFAULT NOW()
);

