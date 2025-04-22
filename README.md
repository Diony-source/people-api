# PeopleHub API

A production-ready REST API for managing people using Go, PostgreSQL, and Clean Architecture.

## Features

- GET /people – list all people
- POST /people – create person
- DELETE /people/{id} – delete person
- PATCH /people/{id} – update person
- GET /people/search?name=X – search people by name
- GET /people/stats – count total people

## Tech Stack

- Go 1.23.4
- PostgreSQL
- sqlx
- Clean Architecture
- .env config
- RESTful API

## Running Locally

```bash
go mod tidy
psql -U postgres -d person_api -f schema.sql
go run main.go
