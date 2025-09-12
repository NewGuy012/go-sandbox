# Go Sandbox

A couple projects written in Go to explore its capabilities.

## Description

- go-proverbs-greetings: generate a random Go proverb in a greeting.
- sample-rest-api-gin: a sample RESTful web service API with Go and the Gin Web Framework.

## Getting Started

### Dependencies

* go: https://go.dev/doc/install

### Executing go-proverbs-greetings program

1. Check that you have go installed
```
go version
```
2. Change directory into module
```
cd go-proverbs-greetings
```
3. Run code
```
go run .
```

### Executing sample-rest-api-gin program

1. Check that you have go installed
```
go version
```
2. Change directory into module
```
cd sample-rest-api-gin
```
3. Run code
```
go run .
```
4. From a command line window, use curl to make a request to your running web service
```
curl http://localhost:8080/albums
curl http://localhost:8080/albums/2
```

## Authors

Moses Yoo, juyoung.m.yoo at gmail dot com.

## Version History

* 0.1
    * Initial Release

## License

This project is licensed under the BSD 3-Clause license. See the LICENSE.md file for details.

## Acknowledgments

These projects were inspired by the Go tutorials: https://go.dev/doc/tutorial/
