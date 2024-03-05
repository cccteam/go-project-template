# Database Migrations

The database is managed via migrations in this folder. To run migrations, you must have the [migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate) cli installed.

## Setup

The database server must be configured with a database and user with the appropriate permissions. The following is an example of how to create a database and user with the appropriate permissions.

1. Create the database
2. Create the schema
3. Create the user and ensure the username matches the schema. By default the user will use the schema matching their username. This allows for a flexible schema name without having to specify it in queries.

### Creating Migrations

```bash
# Create a new migration
migrate create -ext sql -dir schema/migrations -seq <NAME>
```

### Running Migrations

```bash
# Run all migrations
migrate -path schema/migrations -database "postgresql://user:password@localhost:5432/GoProjectTemplate?sslmode=disable&x-migrations-table=%22public%22.%22SchemaMigration%22&x-migrations-table-quoted=1" up
```

### Rolling Back Migrations

```bash
# Roll back the last migration
migrate -path schema/migrations -database "postgresql://<user>:<password>@localhost:5432/GoProjectTemplate?sslmode=disable&x-migrations-table=%22public%22.%22SchemaMigration%22&x-migrations-table-quoted=1" down
```

### Migration Naming Conventions

Migrations should be named in the following format:

```bash
<version>_<description>.sql
```

Where version is a sequential number and description is a short description of the migration.

### Migration Files

Migrations should be written in plain SQL. The following is an example of a migration (up and down):

UP:

```sql
-- TABLE: users

-- SQL is executed when this migration is applied
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE
);
```

DOWN:

```sql
DROP TABLE users;
```
