# install dependency
install-dependency:
	go mod tidy

# run http server
run-http-server-local:
	go build -o "./cmd/todolist-http/todolist-http" ./cmd/todolist-http && ./cmd/todolist-http/todolist-http