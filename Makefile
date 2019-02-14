clean:
	rm bin/*
	rm *.so

build:
	go test
	mkdir -p bin/
	go build -o bin/leagopher