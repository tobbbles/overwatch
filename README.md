# Overwatch

A simple Golang REST API that consume a third-party resource, persists it's data into our own datastore, and serves it
through a RESTful API.

Currently the service uses SQLite3 for demonstration purposes.


## Running

### Configuring

The service is configured via three enviornment variables:

- `ADDR` - The Address the service will run on (defaults to `localhost:3000`.)
- `SQLITE_PATH` - The full path to where the SQLite3 file will be (default to `./overwatch.db`.)
- `INTERVAL` - The interval in minutes to fetch data from the Overwatch API (defaults to 30 minutes.)



## Development

### Building
This project uses Go Modules for depencency management and vendoring. Dependencies will be resovled and fetched at build time - freely develop in the application and compile with

```shell script
$ go build . 
```

### Testing
You can run all of the unit tests within the service using Go's own unit test runner.
```shell script
$ go test -v ./...
```

### Local environment
The current employed development workflow revolves around Docker and utilising docker-compose. You can bring up a full local development environment with:

```shell script
$ docker-compose up --build
```

## Deployment

This service was designed to be deployed with [kustomize](https://github.com/kubernetes-sigs/kustomize) and using a Docker container registry to hold built images. You'd ideally contain all manifests and patches in a seperate repo to orchestrate the deployment of the service to a Kubernetes cluster. This repo itself contains a `doc` directory which contains manifests for demonstration purposes only.
