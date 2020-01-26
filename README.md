# Requirements

Before starting to work with the app, you need to execute the following:

```bash
$ go mod init go-rest-docker
```

# Running

## From command line

To run app from the command line, please execute the following:

```bash
$ cd src/
$ go run go-rest.go
```

## Docker

To run image on Docker, please execute the following:

```bash
$ docker build -t go-rest-docker .
$ docker run -it -p 8080:8080 go-rest-docker
```

# Testing

To test the app, execute the following:

```bash
$ curl localhost:8080
```
