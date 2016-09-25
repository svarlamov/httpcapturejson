# httpcapturejson
### Install
`go get github.com/svarlamov/httpcapturejson`
cd into install path
`go run server.go`
### Test
`curl -X POST -d "{\"this\":\"is json\"}" -H "Content-Type: application/json" localhost:3000/hello/resource`
