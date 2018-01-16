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
$ ./bin/k9server
```

## k9ctl CLI Help
```
$ ./bin/k9ctl --help
CLI for interacting with the k9server api

Usage:
  k9ctl [command]

Available Commands:
  help        Help about any command
  zones       Manage zones

Flags:
  -h, --help   help for k9ctl

Use "k9ctl [command] --help" for more information about a command.
```

## k9ctl CLI List all zones
```
$ ./bin/k9ctl zones list
ID                                    |NAME
23ae79af-7652-4776-915a-8501de481c85  |zone 1
d5045e80-b602-4e7c-a04a-7a1a6f4efa32  |zone 2
54e15b6c-65c4-4f05-b5b7-97df6f59e1aa  |zone 3
82f1944e-87f1-4736-bf33-226232eb3569  |zone 4
```

## k9ctl CLI List all zones (as json)
```
$ ./bin/k9ctl zones list -o json
[
  {
    "name": "zone 1",
    "id": "23ae79af-7652-4776-915a-8501de481c85"
  },
  {
    "name": "zone 2",
    "id": "d5045e80-b602-4e7c-a04a-7a1a6f4efa32"
  },
  {
    "name": "zone 3",
    "id": "54e15b6c-65c4-4f05-b5b7-97df6f59e1aa"
  },
  {
    "name": "zone 4",
    "id": "82f1944e-87f1-4736-bf33-226232eb3569"
  }
]
```

## k9ctl CLI Create a new zone
```
$ ./bin/k9ctl zones create myZone
ID                                    |NAME
a603abc4-246f-4069-8851-c89549d66547  |myZone
```

## k9ctl CLI Delete a zone
```
$ ./bin/k9ctl zones delete 468be827-dc13-4e15-87f4-a556181d6d42
```
