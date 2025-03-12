# QR-Code-Backend Documentation

This document provides an overview of the **QR-Code-Backend** service. The backend is built using GO, gin, and it supports basic QR code operations.

---

## Overview

The **QR-Code-Backend** is responsible for managing QR code tokens. Its primary functions include:

- **Generating QR tokens**
- **Validating QR tokens**
- **Invalidating QR tokens**
- **Deleting QR tokens**

---

## Deployment Details

- **Server IP:** 116.203.82.82
- **Port:** 8093

---

## API Endpoints

The backend provides the following endpoints:

- **/api/qr-code/generate:** Creates a new QR token.
- **/api/qr-code/validate:** Checks if a provided QR token is valid.
- **/api/qr-code/invalidate:** Marks a QR token as invalid.
- **/api/qr-code/delete:** Permanently removes a QR token from the system.

---

This concise documentation summarizes the core functions and deployment details of the **QR-Code-Backend**. For further information or code-level details, refer to the developer guides or source code repository.
