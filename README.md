## Integration test setup with ginkgo:
https://www.ardanlabs.com/blog/2019/03/integration-testing-in-go-executing-tests-with-docker.html

## Mocking http api call:
https://www.alexhyett.com/mock-api-calls-wiremock/#adding-mappings-using-files


##Upgrade go version
Currently my machine has GO version 1.16. When I was trying to install zap dependency, It gives version  mismatch error as Zap needs at least 1.18 version as it only support last 2 version of Go
```
➜  rateService git:(bdd-test-setup) ✗ go get -u go.uber.org/zap
# go.uber.org/atomic
../../go/pkg/mod/go.uber.org/atomic@v1.10.0/error.go:55:12: x.v.CompareAndSwap undefined (type Value has no field or method CompareAndSwap)
../../go/pkg/mod/go.uber.org/atomic@v1.10.0/error.go:61:24: x.v.Swap undefined (type Value has no field or method Swap)
../../go/pkg/mod/go.uber.org/atomic@v1.10.0/string.go:58:12: x.v.CompareAndSwap undefined (type Value has no field or method CompareAndSwap)
../../go/pkg/mod/go.uber.org/atomic@v1.10.0/string.go:64:12: x.v.Swap undefined (type Value has no field or method Swap)
note: module requires Go 1.18
```

### command to upgrade
```
brew install go@1.18
brew link go@1.18
brew link --overwrite go@1.18
and finally to keep the path
echo 'export PATH="/usr/local/opt/go@1.18/bin:$PATH"' >> ~/.zshrc
```

