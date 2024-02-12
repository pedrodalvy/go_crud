# Go CRUD

A basic CRUD application in Go

## Prerequisites

Ensure that you have the following tools installed:

- [Docker](https://www.docker.com)
- [Go](https://golang.org) 1.21.1
- [Air](https://github.com/cosmtrek/air) Live reload for Go apps

## Installation

Follow these steps to set up the project:

1. Run `docker compose up -d` to start the database and server in live reload mode

Or, if you prefer to use a local Go server, follow these steps:

1. Run `docker compose up -d postgres` to start the database
2. Install dependencies with `go mod tidy`
3. Run `air` to start the server in live reload mode or `go run main.go` to start the server without live reload

## Usage

An endpoint is available at `http://localhost:3000` for testing the server

- Health check: `/health`
- Get all products: `/api/products`
- Get a specific product: `/api/products/:id`
- Create a new product: `/api/products`
- Update a product: `/api/products/:id`
- Delete a product: `/api/products/:id`
