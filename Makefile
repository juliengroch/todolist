serve:
	@(echo "-> Running app...")
	@(TODOLIST_CONF=config.json gin -i run run r)

migrate:
	@(echo "-> migrate app...")
	@(TODOLIST_CONF=config.json go run main.go m)

test:
	@(echo "-> Running unit tests...")
	@(TODOLIST_CONF_TEST=/home/julien/goWork/src/github.com/juliengroch/todolist/config_test.json go test -v ./tests/views/...)