# go-mongodb-replica

This project is for testing MongoDB **replica set** behavior using a simple Go HTTP API with **GET** and **POST** endpoints. It is intended for experimentation and learning.

## Features

- Connects to a MongoDB replica set
- Basic HTTP API to interact with MongoDB
- Simple GET and POST operations for testing reads/writes

## Requirements

- Go 1.24+
- Docker & Docker Compose
- MongoDB replica set configured via `docker-compose.yml`

## Getting Started

### 1. Clone the repository

```bash
git clone https://github.com/jandranrebbalangue/go-mongodb-replica.git
cd go-mongodb-replica
```
## Running the Project

Use Docker Compose to start the MongoDB replica set and the Go app:

```bash
docker compose up -d
```
