# Rate

Welcome to the repository of exchange rate project!

## Description



## Features


1. **Database Integration**: All application data is stored in a database. The service includes functionality to handle database migrations upon startup.
2. **Docker Support**: The repository includes a Dockerfile and Docker Compose configuration to facilitate running the system in Docker containers.

## Setup and Usage

### Prerequisites

- Docker
- Docker Compose

### Running the Application

1. **Clone the Repository**

    ```bash
    git clone https://github.com/sergey4qb/rate.git
    cd rate
    ```
2. **Set Environment Variables**

    Before running the application, go to `rate/configs` and configure the environment file. Add the following environment variables:

    - `SMTP_HOST`: The SMTP server host.
    - `SMTP_PORT`: The SMTP server port.
    - `SMTP_USERNAME`: The username for the SMTP server (if required).
    - `SMTP_PASSWORD`: The password for the SMTP server (if required).
    - `SMTP_FROM`: The email address used in the "From" field of the notification emails.
    - `SMTP_SUBJECT`: The subject of the notification emails. Default is "Subject: Daily exchange rates\n".
    - `GIN_MODE`: The mode in which the Gin web framework runs. Default is `release`.
    - `HTTP_PORT`: The port on which the HTTP server runs. Default is `8081`.
    - `SCHEDULED_TIME`: The time at which scheduled tasks (like sending daily notifications) run. Format is `HH:MM`.

    Additionally, set up the following PostgreSQL environment variables:

    - `POSTGRES_USER`: The PostgreSQL database user. Default is `myuser`.
    - `POSTGRES_PASSWORD`: The PostgreSQL database password. Default is `mypassword`.
    - `POSTGRES_DB`: The PostgreSQL database name. Default is `mydatabase`.
    - `POSTGRES_HOST`: The PostgreSQL database host. Default is `db`.
    - `POSTGRES_PORT`: The PostgreSQL database port. Default is `5432`.




3. **Build and Run with Docker Compose**

    ```bash
    docker-compose up --build
    ```

   This will start the service along with the necessary database.

### Database Migrations

Database migrations are executed automatically when the service starts. This ensures the database schema is up-to-date with the current version of the application.
