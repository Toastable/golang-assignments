module main

go 1.22.6

replace web_server => ../../modules/web_server
replace todo_service => /../../modules/todo_service


require(
    web_server v0.0.0-00010101000000-000000000000
)

require (
	github.com/google/uuid v1.6.0 // indirect
    todo_service v0.0.0-00010101000000-000000000000 // indirect
)
