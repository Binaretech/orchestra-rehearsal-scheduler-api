# Orchestra Rehearsal Scheduler API

Welcome to the Orchestra Rehearsal Scheduler API! This project is designed to help orchestras manage their rehearsal schedules efficiently. The API provides endpoints for managing users, sections, instruments, and more.

## Table of Contents

- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
  - [Running the Application](#running-the-application)
- [Configuration](#configuration)
- [Database Migrations](#database-migrations)
- [API Documentation](#api-documentation)
- [Contributing](#contributing)
- [License](#license)

## Getting Started

### Prerequisites

Before you begin, ensure you have the following installed on your machine:

- [Docker](https://www.docker.com/get-started)
- [Go](https://golang.org/dl/) (version 1.23 or higher)

### Installation

1. Create a `.env` file in the root directory and add the following environment variables:

```env
PORT=8080
DATABASE_HOST=db
DATABASE_USER=postgres
DATABASE_PASS=secret
DATABASE_NAME=scheduler
DATABASE_PORT=5432
```

### Running the Application

1. Start the application using Docker Compose:

```bash
docker-compose up --build
```

2. The API server will be running on `http://localhost:8080`.

## Configuration

The application can be configured using environment variables. The following variables are available:

- `PORT`: The port on which the API server will run (default: 8080).
- `DATABASE_HOST`: The hostname of the PostgreSQL database (default: db).
- `DATABASE_USER`: The username for the PostgreSQL database (default: postgres).
- `DATABASE_PASS`: The password for the PostgreSQL database (default: secret).
- `DATABASE_NAME`: The name of the PostgreSQL database (default: scheduler).
- `DATABASE_PORT`: The port on which the PostgreSQL database is running (default: 5432).

## Database Migrations

Database migrations are handled using [Goose](https://github.com/pressly/goose). Goose is available inside the `app` container. To run migrations, follow these steps:

1. Ensure the `app` and `db` containers are running:

```bash
docker-compose up --build
```

2. Access the `app` container:

```bash
docker-compose exec app bash
```

3. Run the migrations:

```bash
goose up
```

This will apply all pending migrations to the database.

## Contributing

We welcome contributions to the project! To contribute, please follow these steps:

1. Fork the repository.
2. Create a new branch for your feature or bugfix.
3. Make your changes and commit them with a clear message.
4. Push your changes to your fork.
5. Create a pull request to the main repository.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
