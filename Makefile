build: 
	go build -o server main.go

run: build
	./server

watch:
	ulimit -n 10000
	go run main.go migrate-seed
	reflex -s -r '\.go$$' make run

migrate:
	ulimit -n 10000
	go run main.go migrate
	reflex -s -r '\.go$$' make run

migrate-seed:
	ulimit -n 10000
	go run main.go migrate-seed
	reflex -s -r '\.go$$' make run