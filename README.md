# Golang REST API bootstrap project

## Build and run

Required `docker`, `git`

Clone this repository
```
git clone repo && cd repo
```

Build docker image
```
docker build -t github.com/sazarkin/golang-rest-api-example .
```

Run docker container

```
dokcer run --rm -p 8080:8080 github.com/sazarkin/golang-rest-api-example
```

Test endpoint

```
http -v :8080/pokemon/charizard
```
