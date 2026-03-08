build:
	go build -o hotreload.exe ./cmd/hotreload

run:
	go run ./cmd/hotreload --root "./testserver" --build "go build -o ./bin/server.exe ./testserver" --exec "./bin/server.exe"

clean:
	del hotreload.exe