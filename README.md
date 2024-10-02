# Funny movies server

## Introduction

A server serving a web application for sharing Youtube videos.

### Key features

- User registration and login
- Sharing YouTube videos
- Viewing a list of shared videos
- Real-time notifications for new video shares 

### Technologies

- Programming Language: [Golang](https://go.dev/)
- Framework: [Echo](https://echo.labstack.com/)
- ORM: [GORM](https://gorm.io/index.html)
- DBMS: [PostgreSQL](https://www.postgresql.org/)
- Deployment: [Vercel](https://vercel.com/), [Render](https://render.com/), [Docker](https://www.docker.com/)

## Prerequisites

- [Go](https://golang.org/doc/install) ^1.22
- [Docker](https://docs.docker.com/install/) ^24.0
- [Docker Compose](https://docs.docker.com/compose/install/) ^2.28.0

## Installation & Configuration

After cloning repository, follow the commands below:
1. Copy `env.sample` file to `env.` or run this command:
   ```
   cp .env.sample .env
   ```
2. Install dependencies:
   ```
   make depends
   ```
3. Run docker:
   ```
   make docker.run
   ```
4. Run migration and seed data:
   ```
   make migrate
   ```

## Running the Application

Before running the development server or running unit tests, please run docker first:
   ```
   make docker.run
   ```

Start the development server:
   ```
   make start
   ```

For run unit tests:
   ```
   make test
   ```

## Troubleshooting

The `make` commands in the instructions above just work with Linux or MacOS. If you use Window, you can install WSL to run `make` command or run the corresponding subcommands in `Makefile`.

