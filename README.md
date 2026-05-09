# prusk-sql

Repo for my sql cmd utility written in Go

This utility is designed to connect to SQL Server instances and execute queries. It supports connection testing, retrieving server information, listing databases, and executing custom queries.

## Features

- Test SQL Server connections with configurable timeouts
- Retrieve server name and version information
- List all databases in the instance
- Execute custom SQL queries with formatted output
- Environment variable support for secure password handling
- Debug mode for troubleshooting

## Security

**Important**: Never use default passwords in production. The utility supports environment variables for secure password handling.

## Usage

### Basic Connection Test

Connect to SQL Server and retrieve basic information:

```bash
prusk-sql -server 127.0.0.1 -port 1433 -user sa -password YourPassword
```

### Using Environment Variables

Set the password securely using environment variables:

```bash
export SQL_PASSWORD=YourSecurePassword
prusk-sql -server 127.0.0.1 -port 1433 -user sa
```

### List Databases

Add the `-databases` flag to list all databases:

```bash
prusk-sql -server 127.0.0.1 -port 1433 -user sa -password YourPassword -databases
```

### Execute Custom Queries

Use the `-query` flag to execute custom SQL queries:

```bash
prusk-sql -server 127.0.0.1 -port 1433 -user sa -password YourPassword \
  -query "SELECT name, database_id FROM sys.databases WHERE database_id > 4"
```

### Connection Timeout

Set a custom connection timeout (default 30s):

```bash
prusk-sql -server 127.0.0.1 -port 1433 -user sa -password YourPassword -timeout 10s
```

### Show Version

Display version information:

```bash
prusk-sql -version
```

## Command Line Options

- `-server`: Database server address (default: "localhost")
- `-port`: Database port (default: 1433)
- `-user`: Database username (default: "sa")
- `-password`: Database password (can also use SQL_PASSWORD env var)
- `-databases`: List all databases in the instance
- `-query`: Custom SQL query to execute
- `-timeout`: Connection timeout (default: 30s)
- `-debug`: Enable debug output
- `-version`: Show version information