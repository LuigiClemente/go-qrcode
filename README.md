# QR Code API Documentation

## Overview

The **QR-Code-Backend** is responsible for managing QR code tokens. Its primary functions include:

- **Generating QR tokens**
- **Validating QR tokens**
- **Invalidating QR tokens**
- **Deleting QR tokens**

---

## Base URL

```
http://<your-server-address>/api/v1/qrcode-app
```

## Endpoints

### 1. Generate QR Code

**Endpoint:**

```
POST /generate
```

**Description:**
Generates a new QR code and returns the QR code details.

**Response Example:**

```json
{
  "id": 1,
  "token": "qr-1710328479827398",
  "url": "",
  "image": "./qr_images/qr-1710328479827398.png",
  "valid": true,
  "created_at": "2025-03-13T10:00:00Z"
}
```

---

### 2. Validate QR Code

**Endpoint:**

```
GET /validate/:token
```

**Description:**
Checks if a QR code is valid.

**Request Example:**

```
GET /validate/qr-1710328479827398
```

**Response Example:**

```json
{
  "valid": true
}
```

---

### 3. Invalidate QR Code

**Endpoint:**

```
POST /invalidate
```

**Description:**
Marks a QR code as invalid.

**Request Body:**

```json
{
  "token": "qr-1710328479827398"
}
```

**Response Example:**

```json
{
  "message": "QR code invalidated successfully"
}
```

---

### 4. Delete QR Code

**Endpoint:**

```
DELETE /qr/:token
```

**Description:**
Permanently deletes a QR code and its associated image.

**Request Example:**

```
DELETE /qr/qr-1710328479827398
```

**Response Example:**

```json
{
  "message": "QR code deleted successfully"
}
```

---

## Notes

- Ensure the API server is running before making requests.
- `token` is a unique identifier for each QR code.
- Deleting a QR code also removes its image file from the system.
