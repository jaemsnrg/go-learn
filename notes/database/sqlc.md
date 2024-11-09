# SQLC - SQL Compiler

commands:
sqlc init                    # Create a new sqlc.yaml config file
sqlc generate               # Generate Go code from SQL queries
sqlc verify                 # Verify SQL queries are valid
sqlc version               # Show sqlc version information


SQLC is a powerful tool that generates type-safe Go code from SQL queries. It offers several advantages:

## Key Features
- Generates Go code directly from SQL queries
- Provides compile-time type safety
- No runtime reflection or query building
- Supports PostgreSQL, MySQL, and SQLite
- Validates SQL queries during code generation

## Benefits
- Catches SQL errors at compile time rather than runtime
- Eliminates boilerplate database code
- Maintains type safety between database and Go code
- Better performance than ORMs due to direct SQL usage
- Easier to maintain as SQL queries are in .sql files

## Workflow
1. Write SQL queries in .sql files
2. Define database schema
3. Run sqlc generate command
4. Use generated Go code in your application

SQLC is particularly useful for applications like our simple_bank project where type safety and performance are important considerations.
