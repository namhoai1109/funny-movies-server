# Funny movies server

## Prerequisites

- [Go](https://golang.org/doc/install) ^1.22
- [Docker](https://docs.docker.com/install/) ^18.09.2
- [Docker Compose](https://docs.docker.com/compose/install/) ^1.23.2

## Getting started

1. Initialize the app for the first time:
   ```
   make provision
   ```
2. Generate swagger API docs:
   ```
   make specs
   ```
3. Run the development server:
   ```
   make start
   ```

The application runs as an HTTP server at port 8000. 

## Deployment

