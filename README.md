# What the Bill?

Website for viewing the latest bills being passed.

## Database

Create new migration
```bash
migrate create -ext sql -dir db/migrations -seq create_users_table
```

Apply migrations
```bash
migrate -source file://db/migrations -database postgres://localhost:5432/database up 
```