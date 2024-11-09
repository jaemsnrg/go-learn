# ACID Principles in Database Transactions

ACID is an acronym that stands for Atomicity, Consistency, Isolation, and Durability. These principles ensure reliable processing of database transactions.

## Atomicity
- All operations in a transaction succeed or they all fail
- Transaction is treated as a single unit
- If any part fails, the entire transaction is rolled back
- Example: In a money transfer, both debit and credit must succeed or neither happens

## Consistency 
- Database must be in a valid state before and after transaction
- All rules, constraints, and triggers are enforced
- Data remains consistent and valid across the database
- Example: Account balances must remain valid after transfers

## Isolation
- Multiple transactions can occur concurrently without interfering
- Changes from incomplete transactions remain isolated
- Prevents dirty reads, non-repeatable reads, and phantom reads
- Example: Two simultaneous transfers don't corrupt account balances

## Durability
- Once a transaction is committed, it stays committed
- Changes survive system crashes or power failures
- Usually implemented through transaction logs
- Example: Confirmed transfers persist even if system crashes

These principles are especially important for financial transactions like the ones in our simple_bank database, where data integrity is critical.
