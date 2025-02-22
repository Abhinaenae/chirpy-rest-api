<div align="left">
    <img src="https://github.com/user-attachments/assets/acf2178e-8b45-425d-87fa-a8a32dcf55d2" width="40%" align="left" style="margin-right: 15px"/>
    <div style="display: inline-block;">
        <h2 style="display: inline-block; vertical-align: middle; margin-top: 0;">CHIRPY-REST-API</h2>
        <p>
    <em><code>❯ A lightweight Twitter-like social network backend written in Go</code></em>
</p>
        <p>
    <img src="https://img.shields.io/github/license/abhinaenae/chirpy-rest-api?style=flat-square&logo=opensourceinitiative&logoColor=white&color=E92063" alt="license">
    <img src="https://img.shields.io/github/last-commit/abhinaenae/chirpy-rest-api?style=flat-square&logo=git&logoColor=white&color=E92063" alt="last-commit">
    <img src="https://img.shields.io/github/languages/top/abhinaenae/chirpy-rest-api?style=flat-square&color=E92063" alt="repo-top-language">
    <img src="https://img.shields.io/github/languages/count/abhinaenae/chirpy-rest-api?style=flat-square&color=E92063" alt="repo-language-count">
</p>
        <p>Built with the tools and technologies:</p>
        <p>
<img src="https://img.shields.io/badge/Go-00ADD8.svg?style=flat-square&logo=Go&logoColor=white" alt="Go">
<img src="https://img.shields.io/badge/PostgreSQL-316192.svg?style=flat-square&logo=PostgreSQL&logoColor=white" alt="PostgreSQL">
<img src="https://img.shields.io/badge/Docker-2496ED.svg?style=flat-square&logo=Docker&logoColor=white" alt="Docker">
<img src="https://img.shields.io/badge/Goose-Active-brightgreen.svg" alt="Goose">
<img src="https://img.shields.io/badge/sqlc-Enabled-blue.svg" alt="sqlc">


</p>
    </div>
</div>
<br clear="left"/>

## 🚀 Overview

Chirpy is a social network similar to Twitter, designed to be lightweight and scalable. This repository contains Chirpy's backend server, built with Go. The project was created as an educational exercise to learn about building web servers, authentication, and API development using Go.

## ✨ Features

- 🏗 **User Authentication**: Supports user registration, login, token refresh, and token revocation.
- 📢 **Chirp Management**: Users can create, retrieve, and delete chirps (posts).
- 📊 **Metrics & Health Checks**: Provides system health status and operational metrics.
- 🔐 **JWT-Based Authentication**: Implements secure token-based authentication.

## 📂 Project Structure

```sh
└── chirpy-rest-api/
    ├── LICENSE
    ├── README.md
    ├── go.mod
    ├── go.sum
    ├── internal/
    │   ├── auth/        # Authentication handlers
    │   ├── database/    # Database interaction layer
    ├── sql/
    │   ├── queries/     # SQL query definitions
    │   └── schema/      # Database schema migrations
    ├── src/
    │   ├── chirps.go    # Chirp handling logic
    │   ├── main.go      # Application entry point
    │   ├── metrics.go   # Metrics and monitoring endpoints
    │   ├── refresh.go   # Refresh token endpoints
    │   ├── user.go      # User management
    │   ├── readiness.go # Health Check
    ├── sqlc.yaml
    └── .env
```

## 🛠 Getting Started

### 📋 Prerequisites

Ensure you have the following installed:

- [Go](https://go.dev/dl/) (1.22+ recommended)
- [SQLC](https://sqlc.dev/)
- [Goose](https://github.com/pressly/goose)
- [PostgreSQL](https://www.postgresql.org/) (or modify for SQLite)

### 📥 Installation

Clone the repository:

```sh
git clone https://github.com/abhinaenae/chirpy-rest-api.git
cd chirpy-rest-api
```

Install dependencies:

```sh
go mod tidy
```

### 🚀 Running the Server

```sh
go run src/main.go
```

## 🔗 API Endpoints

### 🔍 Health Check

#### `GET /api/healthz`
Returns the current status of the system.

### 📈 Metrics

#### `GET /admin/metrics`
Returns system metrics.

#### `GET /api/reset`
Resets system metrics.

### 🔑 Authentication

#### `POST /api/login`
Authenticates a user.

#### `POST /api/refresh`
Issues a new access token using a refresh token.

#### `POST /api/revoke`
Revokes a refresh token.

### 👤 Users

#### `POST /api/users`
Creates a new user.

#### `PUT /api/users`
Updates user information.

### 📝 Chirps

#### `POST /api/chirps`
Creates a new chirp.

#### `GET /api/chirps`
Retrieves all chirps.

#### `GET /api/chirps/{chirpId}`
Retrieves a specific chirp by ID.

#### `DELETE /api/chirps/{chirpId}`
Deletes a chirp.

## 🤝 Contributing

Contributions are welcome! Please fork the repository, create a feature branch, and submit a PR.

## 📜 License

This project is licensed under the MIT License.

## 🙌 Acknowledgments

- Inspired by Twitter’s API design.
- Built using Go, PostgreSQL, SQLC, and Goose

