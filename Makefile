

make: client server

client:
	go build -o ./bin/ pocok/cmd/client

server:
	go build -o ./bin/ pocok/cmd/server

clean:
	rm ./bin/*
