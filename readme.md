# Basic Go API app w/Docker
Basic GoLang REST API using docker with database connection.

*NOTE*: This is an example and it is not production ready


_Requirements_: Docker

## Resources
 - [Your Source](#your-source)
 - [Dependencies](#dependencies)
 - [Running the API](#running-the-api)
 - [DB Administer](#running-the-api)
 - [Database](#running-the-api)

### Your Source
Modify the `docker-compose.yml` file and modify the volumes.

```
<your_source_location>:/go/src/github.com/user/app
<your_config_location>:/etc/myapi/
```

### Dependencies
Dependencies should be handled using a dependency manager like `go-deps` but for the sake of the example their are added manually and with an image re-build.

For doing so, modify the `Dockerfile` and add any `RUN go get <package>` required.


### Running the API
For running the whole project simply run the following and then enter to `http://localhost:8081`

```bash
docker-compose up
```

For building the docker image:
```bash
docker build .
```

For running the API execute:
```bash
docker run -it -p 8081:8081 -v <your_source_location>:/go/src/github.com/user/app godemo
```

This docker image contains `fresh` which will re-compile when it detects a code modification.

### DB Administer
In order to view your database go to your browser `http://localhost:8080`

### Database
The service contains a MySQL Database 