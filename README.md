# Fibonacci Web Service

This project is a web service that calculates the Fibonacci sequence. It is built using Go and the Gin framework, and it can be run inside a Docker container.

## Prerequisites

- Go 1.23.2 or later
- Docker

## Getting Started

### Clone the Repository

```sh
git clone https://github.com/Waterdrips/fibonacci-web-service.git
cd fibonacci-web-service
```

### Install the dependencies:

```sh
make deps
```

### Build the project:

```sh
make build
```

### Test the project:

```sh
make test
```

### Format the code:

```sh
make fmt
```

### Build the docker image:

```sh
make docker-build
```

### Run the docker container:

```sh
make run-docker
```

The server will be running on `http://localhost:8080`.

## API

### `POST /`

Calculates the Fibonacci sequence.

Here is an example request and response:
```shell
curl -X POST http://localhost:8080/ -H "Content-Type: application/json" -d '{"number": 4}'

{"result":[0,1,1,2,3,5,8,13,21]}
```

#### Request body

```json
{
  "number": 10
}
```

#### Response

```json
{
  "sequence": [0, 1, 1, 2, 3, 5, 8, 13, 21, 34]
}
```

## Gracefully shutdown the server:

The server supports graceful shutdown. It will wait for ongoing requests to complete before shutting down.

## CI 

There is a CI pipeline that runs the tests and linter on every push to the main branch.

see the [GitHub Actions](.github/workflows/go.yml) file.

## Future Improvements

- Add deployment to a cloud provider or other hosting service. (for example provide yaml files for Kubernetes)
- Push the Docker image to a container registry.
- Add more tests.
- Investigate the performance of the Fibonacci calculation for large numbers.
- Add a cache to store the results of previous calculations. (If service is expected to receive many requests for the same number.)
- Add a rate limiter to prevent abuse.
- Add a logger to log requests and responses.
- Add tracing and metrics.

## Known Issues
- Large numbers for the sequence will cause the server to panic, as the slice for results will exceed the maximum 
  capacity. This could be fixed by using a different data structure or by paginating the response.
- Large numbers may cause the server to run out of memory. This could be fixed by paginating the response or by 
  calculating the sequence in chunks. Alternatively, the server could return an error if the number is too large.

## License
This project is licensed under the MIT License.