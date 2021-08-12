# books-grpc

gRPC Implementation using Go.

# How To Use

1. There are two branch in this project:

- choose `local-storage` for gRPC implementation with local storage.
- choose `master` for gRPC implementation with actual database.

2. Clone this repository into your local machine.

3. If `master` branch is cloned, create a new database called `booksdb` then create a new table called `books`.

```sql
CREATE DATABASE booksdb;
```

```sql
USE booksdb;
-- create new table
CREATE TABLE books(
    id VARCHAR(255) PRIMARY KEY,
    title VARCHAR(255),
    author VARCHAR(255),
    is_read BOOLEAN
);
```

4. Configure the database connection in `.env` file, the example can be seen in `.env.example`.

5. Start the gRPC server with `go run main.go`.
