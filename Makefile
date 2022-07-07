PROJECT = usercontrol

DIR_SERVER = src/server
DIR_TCP_SERVER = $(DIR_SERVER)/tcp

test_tcp_server:
	@echo "Testing TCP Server"
	go test -v $(PROJECT)/$(DIR_TCP_SERVER)