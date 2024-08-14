module common

go 1.22.6

replace todo_service => /../todo_service

replace todo_inmemory_service => /../todo_inmemory_service

require todo_service v0.0.0-00010101000000-000000000000

require todo_inmemory_service v0.0.0-00010101000000-000000000000

require github.com/google/uuid v1.6.0 // indirect
