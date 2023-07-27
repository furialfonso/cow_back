# DockerGoProject

It is a 3-tier based architecture with dependency injection.

**Author**
  - *Andres Felipe Alfonso Ortiz*

**Technologies**
  - *Golang*: programming language.
  - *Mysql*: data persistence.
  - *Gin*: framework for rest applications.
  - *Mokery*: automatic mocks for unit tests.
  - *Dig*: automatic dependency injection.

**Run unit tests**
```
  export CONFIG_DIR=${workspaceRoot}/DockerGoProject/pkg/config && export SCOPE=local && go test -v ./... -covermode=atomic -coverprofile=coverage.out -coverpkg=./... -count=1
```

**Look result in html**
```
  go tool cover -html=coverage.out
```
**Gin**
  - Documentation
    - https://gin-gonic.com/docs/quickstart/
**Mokery**
  - Documentacion
    - https://vektra.github.io/mockery/installation/#homebrew
  - Instalacion mac
    ```
      brew install mockery
    ```
  - Crear mocks
    ```
      mockery --all --disable-version-string
    ```
**Dig**
  - Documentation
    - https://ruslan.rocks/posts/golang-dig
    - https://www.golanglearn.com/golang-tutorials/golang-dig-a-better-way-to-manage-dependency/

**Mysql**
  - Added DB with user configuration test_R and test_W.
  - The Gpool is removed since it causes it to be slower and GO already manages the automatic pool of connections.
    - https://koho.dev/understanding-go-and-databases-at-scale-connection-pooling-f301e56fa73

**Docker**
  - Commands for containerized database
  ```
    docker container run \
      -dp 3306:3306 \
      --name cow_db \
      --env MARIADB_USER=andre \
      --env MARIADB_PASSWORD=Admin123 \
      --env MARIADB_ROOT_PASSWORD=Admin123 \
      --env MARIADB_DATABASE=cow_db \
      mariadb:jammy
  ```