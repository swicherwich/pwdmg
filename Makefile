build:
	go build -o bin/pwdmg cmd/pwdmg/main.go

test:
	go test -v ./...