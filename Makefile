# Make file for Copper-Monkey GoLang project

install-cobra:
	go get -u github.com/spf13/cobra

build-green:
	go build -o bin/copper-monkey-green.exe cmd/green/main.go

run-green-tests:
	go test ./cmd/green/cmd/green_test.go

# .\bin\copper-monkey-green.exe addclientpaymentrequest --Amount 19.99 --FirstName Garry --Surname Taylor --Email gpltaylor@gmail.com `
