# cubegorithm

Writing bunch of code trying to solve rubiks cube most optimally and learning golang at the same time.



## golangci-lint

Install:

```bash
# binary will be $(go env GOPATH)/bin/golangci-lint
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.54.2

golangci-lint --version
```


Run:
```bash
golangci-lint run
```


## Test

```bash
go test -v ./...
```