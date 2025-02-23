<div align="left">
    <img src="https://github.com/user-attachments/assets/43b48ae3-da26-43fa-af44-150159f1cc92" width="40%" align="left" style="margin-right: 15px"/>
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

Chirpy is a social network similar to Twitter, designed to be lightweight and scalable. This repository contains Chirpy's backend server, which was built with Go. The project was created to learn about building web servers, authentication, and API development using Go.

## ✨ Features

- 🏗 **User Authentication**: Supports user registration, login, token refresh, secure password storage via hashing, and token revocation.
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
    ├── Dockerfile
    └── .env
```

## 🛠 Getting Started

### 📋 Prerequisites

Ensure you have the following installed:

- [Docker]
- [Go](https://go.dev/dl/) (1.22+ recommended)
- [PostgreSQL](https://www.postgresql.org/) (or modify for SQLite)

### 📥 Installation

1. Clone the repository:

```sh
git clone https://github.com/abhinaenae/chirpy-rest-api.git
cd chirpy-rest-api
```

2. Modify `.env` file accordingly.
- Create a secret for your server using `openssl rand -base64 64` and store it as `JWT_SECRET=<secret>`
- Ensure you have a postgreSQL database that can be connected from the `DB_URL` string

3. Docker build and run at root of the application:
```Docker
docker build -t chirpy .
docker run -p 8080:8080 chirpy
```

## 🔗 API Endpoints

### 🔍 Health Check

#### `GET /api/healthz`
Returns the current status of the system.

### 📈 Metrics

#### `GET /admin/metrics`
Returns system metrics.

#### `GET /admin/reset`
Deletes all users from the database, as well as their chirps and refresh tokens.

### 🔑 Authentication

#### `POST /api/login`
Authenticates a user by including request body:
```JSON
{
  "email": "abhi@test.com",
  "password": "password"
}
```
and returns a JSON Response with a JWT token expiring in 1 hour and a refresh token expiring in 60 days.
```JSON
{
  "id": "134c96d5-39df-423f-818d-2abe58ab8bc5",
  "created_at": "2025-02-22T15:36:25.733924Z",
  "updated_at": "2025-02-22T15:36:25.733924Z",
  "email": "abhi@test.com",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJjaGlycHkiLCJzdWIiOiIxMzRjOTZkNS0zOWRmLTQyM2YtODE4ZC0yYWJlNThhYjhiYzUiLCJleHAiOjE3NDAzMjM4NDMsImlhdCI6MTc0MDMyMDI0M30.YqUUylkxWxpoNYw1ir0NqFD8EtCikAnRUh_vECUZDLI",
  "refresh_token": "de64e7fe8999069b460e7e4576285b3717962a6ca5b22a36ece9ad2713aa6357"
}
```

#### `POST /api/refresh`
Issues a new JWT access token using a refresh token.
- Must send a refresh token as a bearer token

#### `POST /api/revoke`
Revokes a refresh token.
- Must send a refresh token as a bearer token

### 👤 Users

#### `POST /api/users`
Creates a new user.
Expected request body:
```JSON
{
  "email": "abhi@example.com",
  "password": "password"
}
```
Expected response body:
```JSON
{
  "id": "c3bd0bbf-2db5-4124-a316-a2b85e8f230a",
  "created_at": "2025-02-23T08:22:02.55794Z",
  "updated_at": "2025-02-23T08:22:02.55794Z",
  "email": "abhi@example.com"
}
```


#### `PUT /api/users`
Updates user information via a JWT token sent to the api as a bearer token.
Expected request body:
```JSON
{
  "email": "abhi@example.com",
  "password": "password"
}
```
Expected response body:
```JSON
{
  "id": "134c96d5-39df-423f-818d-2abe58ab8bc5",
  "created_at": "2025-02-22T15:36:25.733924Z",
  "updated_at": "2025-02-23T08:24:52.703069Z",
  "email": "abhi@test.com"
}
```

### 📝 Chirps

#### `POST /api/chirps`
Creates a new chirp.


#### `GET /api/chirps`
Retrieves all chirps.
Expected response body:
```JSON
[
  {
    "ID": "00e60731-3699-4d3e-983f-85a78b8131d3",
    "CreatedAt": "2025-02-22T15:37:33.065483Z",
    "UpdatedAt": "2025-02-22T15:37:33.065483Z",
    "Body": "Let’s just say I know a guy... who knows a guy... who knows another guy.",
    "UserID": "134c96d5-39df-423f-818d-2abe58ab8bc5"
  },
  {
    "ID": "4c19fd40-93ba-49a8-b66d-6695518946fe",
    "CreatedAt": "2025-02-22T15:38:29.063462Z",
    "UpdatedAt": "2025-02-22T15:38:29.063462Z",
    "Body": "Test2",
    "UserID": "134c96d5-39df-423f-818d-2abe58ab8bc5"
  }
]
```

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

