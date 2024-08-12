module todo_inmemory_service

go 1.22.5

replace todo_service => ../todo_service

require (
	github.com/google/uuid v1.6.0
	todo_service v0.0.0-00010101000000-000000000000
)
