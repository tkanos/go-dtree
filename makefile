.IPHONE: test coverage

test:
	go test -race -v `go list ./... | grep -v -e /vendor/ -e /mock/`
	go list ./... | grep -v /vendor/ | grep -v /cmd/ | grep -v /handler | xargs -L1 golint -set_exit_status
	go vet `go list ./... | grep -v /vendor/ | grep -v /cmd/`

coverage:
	go test -cover `go list ./... | grep -v -e /vendor/ -e /mock/`
	go test `go list ./... | grep -v -e /vendor/ -e /mock/` -coverprofile=cover.out
	#go tool cover -html=cover.out
	go tool cover -func=cover.out