module todo

go 1.22.5

replace todo_service => ../modules/todo_service

replace todo_inmemory_service => ../modules/todo_inmemory_service

require (
	todo_inmemory_service v0.0.0-00010101000000-000000000000
	todo_service v0.0.0-00010101000000-000000000000
)

require github.com/google/uuid v1.6.0 // indirect
