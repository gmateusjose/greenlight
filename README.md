# Greenlight

## The go.mod file

the `go.mod` file defines a module path, which is an identifier that will be used
as the root import path for your packages. It's a good practice to make the
module path unique for you and your project. Maybe you could base it in a url
that you own

## Directories

- `bin`: contains our compiled application binaries
- `cmd/api`: contains application-specific code for our api
- `internal`: contains any code which isn't application-specific and can be reused
- `migrations`: contains sql migration files for our database
- `remote`: contains configuration files and setup scripts for prod server
- `go.mod`: contains project dependencies
- `Makefile`: contains recipes for automating common admin tasks

Important: Any packages which live in `internal` can only be imported by code
inside our greenlight project directory.

## Endpoints

Method | URL Pattern | Handler | Action
--- | --- | --- | --- |
GET | v1/healthcheck | healthcheckHandler | Show application information
GET | v1/movies | listMoviesHandler | Show the details of all movies
POST | v1/movies | createMovieHandler | Create a new movie
GET | v1/movies/:id | showMovieHandler | Show the details of a specific movie
PUT | v1/movies/:id | editMovieHandler | Update the details of a specific movie
DELETE | v1/movies/:id | deleteMovieHandler | Delete a specific movie

## Routing

`http.ServeMux` is quite limited in terms of its functionality. In particular it
doesn't allow you to route requests to different handlers based on the request
method (GET, POST, etc.), nor does it provide support for clean URLs with
interpolated parameters.

So we'll going to integrate the `httprouter` package with our application. If
you're building a REST API for public consumption, then it is a solid choice.

## Databases

### Restart the database
```
$ sudo service postgresql restart
```

### Metadata
```
database: 'greenlight'
password: 'pa55word'
```

### Driver
```
data source name (DNS): 'postgres://greenlight:pa55word@localhost/greenlight'
```

### Migrations

Using migration tool: [migrate CLI](https://github.com/golang-migrate/migrate)

Create pair of migration files
```
migrate create -seq -ext=.sql -dir=./migrations $migration-name
```

Execute migrations
```
migrate -path=./migrations -database=$database-dsn up
```

# TO FIX

```
http: error: ConnectionError: HTTPConnectionPool(host='localhost', port=4000): Max retries exceeded with url: /v1/movies (Caused by NewConnectionError('<urllib3.connection.HTTPConnection object at 0x7fcd8eda3100>: Failed to establish a new connection: [Errno 111] Connection refused')) while doing GET request to URL: http://localhost:4000/v1/movies
```
