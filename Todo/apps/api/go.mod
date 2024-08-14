module main

go 1.22.6

replace server => /../../modules/server

replace todo_inmemory_service => /../../modules/todo_inmemory_service

require (
	server v0.0.0-00010101000000-000000000000
	todo_inmemory_service v0.0.0-00010101000000-000000000000
)

require github.com/google/uuid v1.6.0 // indirect
