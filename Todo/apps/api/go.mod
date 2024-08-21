module main

go 1.22.6

replace common => /../../modules/common

replace todo_service => /../../modules/todo_service

replace todo_inmemory_service => ../../modules/todo_inmemory_service

require (
	common v0.0.0-00010101000000-000000000000
	todo_inmemory_service v0.0.0-00010101000000-000000000000
)

require (
    todo_service v0.0.0-00010101000000-000000000000 // indirect
	github.com/google/uuid v1.6.0 // indirect
)
