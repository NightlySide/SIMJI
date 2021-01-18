default:
	go run github.com/markbates/pkger/cmd/pkger
	./build-linux.sh

gobuild:
	go run github.com/markbates/pkger/cmd/pkger
	GOOS=linux go build -v -o simji
