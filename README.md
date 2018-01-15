# k9
k9 demo project to show implementation of goswagger.io

This project exists out of 2 components, connected to a rest api 
- k9server (API Server)
- k9ctl (CLI)

## prerequisite
```
$ go get -u github.com/golang/dep/cmd/dep
$ go get -u github.com/go-swagger/go-swagger/cmd/swagger
```

## Install Dependencies (https://github.com/golang/dep)
This command will download all go dependencies and install them to /vendor dir
```
$ dep ensure
```

## Generate swagger code from swagger file
```
$ make k9server-gen-openapi
```

## Building from source
```
$ make
```

## Run the k9server
```
./bin/k9server
```

## k9ctl CLI Help
```
./bin/k9ctl --help
```

## k9ctl CLI Get all zones
```
./bin/k9ctl zones list
```

## k9ctl CLI Get all zones (as json)
```
./bin/k9ctl zones list -o json
```

## k9ctl CLI Create a new zone
```
./bin/k9ctl zones create myZone
```

## k9ctl CLI Create a new zone
```
./bin/k9ctl zones delete 468be827-dc13-4e15-87f4-a556181d6d42
```
