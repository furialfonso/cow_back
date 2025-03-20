# cow_api

It is a 3-tier based architecture with dependency injection.

**Author**
  - *Andres Felipe Alfonso Ortiz*

**Technologies**
  - *Golang*: programming language.
  - *Mysql*: data persistence.
  - *Gin*: framework for rest applications.
  - *Mokery*: automatic mocks for unit tests.
  - *Dig*: automatic dependency injection.
  - *Docker*: application's contenerization.

```mermaid
graph TB
    User((External User))
    KeycloakSystem[("Keycloak<br>Authentication System")]

    subgraph "Shared Wallet Service"
        subgraph "API Layer"
            GinAPI["API Server<br>Gin Framework"]
            Router["URL Router<br>Gin Router"]
            
            subgraph "API Handlers"
                PingHandler["Ping Handler<br>Go"]
                BudgetHandler["Budget Handler<br>Go"]
                TeamHandler["Team Handler<br>Go"]
            end
        end

        subgraph "Domain Layer"
            subgraph "Services"
                BudgetService["Budget Service<br>Go"]
                TeamService["Team Service<br>Go"]
                UserService["User Service<br>Go"]
            end
        end

        subgraph "Infrastructure Layer"
            subgraph "Database"
                DBConnection["Database Connection<br>MySQL"]
                ReadConn["Read Connection<br>MySQL"]
                WriteConn["Write Connection<br>MySQL"]
            end

            subgraph "External Clients"
                KeycloakClient["Keycloak Client<br>REST"]
                RestClient["REST Client<br>Go HTTP"]
            end

            subgraph "Cache System"
                CacheClient["Cache Client<br>Go"]
                UserCache["User Cache<br>Go"]
            end

            subgraph "Repositories"
                BudgetRepo["Budget Repository<br>Go"]
                TeamRepo["Team Repository<br>Go"]
                CacheRepo["Cache Repository<br>Go"]
                KeycloakRepo["Keycloak Repository<br>Go"]
            end

            ConfigManager["Config Manager<br>YAML"]
            JobHandler["Job Handler<br>Go"]
        end

        MariaDB[("MariaDB<br>Database")]
    end

    %% External connections
    User -->|"HTTP Requests"| GinAPI
    GinAPI -->|"Authenticates"| KeycloakSystem

    %% API Layer connections
    GinAPI -->|"Routes"| Router
    Router -->|"Handles"| PingHandler
    Router -->|"Handles"| BudgetHandler
    Router -->|"Handles"| TeamHandler

    %% Handler to Service connections
    BudgetHandler -->|"Uses"| BudgetService
    TeamHandler -->|"Uses"| TeamService

    %% Service to Repository connections
    BudgetService -->|"Uses"| BudgetRepo
    TeamService -->|"Uses"| TeamRepo
    UserService -->|"Uses"| KeycloakRepo

    %% Repository to Infrastructure connections
    BudgetRepo -->|"Reads/Writes"| DBConnection
    TeamRepo -->|"Reads/Writes"| DBConnection
    DBConnection -->|"Read Operations"| ReadConn
    DBConnection -->|"Write Operations"| WriteConn
    ReadConn -->|"Queries"| MariaDB
    WriteConn -->|"Updates"| MariaDB

    %% External client connections
    KeycloakClient -->|"Authenticates"| KeycloakSystem
    KeycloakClient -->|"Uses"| RestClient
    KeycloakRepo -->|"Uses"| KeycloakClient

    %% Cache connections
    CacheClient -->|"Manages"| UserCache
    CacheRepo -->|"Uses"| CacheClient
    JobHandler -->|"Updates"| UserCache

    %% Configuration
    ConfigManager -.->|"Configures"| GinAPI
    ConfigManager -.->|"Configures"| DBConnection
    ConfigManager -.->|"Configures"| KeycloakClient
```

**Run unit tests**
  - execute tests
  ```
    export CONFIG_DIR=$(pwd)/pkg/config && export SCOPE=local && go test -v ./... -covermode=atomic -coverprofile=coverage.out -coverpkg=./... -count=1
  ```
  - Look result in html
  ```
    go tool cover -html=coverage.out
  ```
**Gin**
  - Documentation
    - https://gin-gonic.com/docs/quickstart/

**Mokery**
  - Documentacion
    - https://vektra.github.io/mockery/latest/
  - Instalacion 
    - mac
    ```
      brew install mockery
    ```
    - windows
    ```
    docker pull vektra/mockery
    ```
  - Crear mocks
    - Mac:
    ```
      mockery --all --disable-version-string
    ```
    - Windows:
    ```
      docker run -v $PWD:/src -w /src vektra/mockery --all
    ```
  - Sort app
    ```
      fieldalignment -fix ./...
    ```
    ```
      gofumpt -l -w .
    ```
  
**Dig**
  - Documentation
    - https://ruslan.rocks/posts/golang-dig
    - https://www.golanglearn.com/golang-tutorials/golang-dig-a-better-way-to-manage-dependency/

**Mysql**
  - Added DB with user configuration test_R and test_W.
  - The Gpool is removed since it causes it to be slower and GO already manages the automatic pool of connections.
    - https://koho.dev/understanding-go-and-databases-at-scale-connection-pooling-f301e56fa73

**Start Aplication**
  - Execute the next command for start the application.
  ```
    docker-compose up -d
  ```
**Config project**
  - For unit test
  ```
    "go.testEnvVars": {
          "CONFIG_DIR": "${workspaceRoot}/pkg/config",
          "SCOPE":"local"
      },
  ```
  - Environment vs-code
  ```
    "SCOPE": "local",
    "PORT": "8080",
    "CONFIG_DIR": "${workspaceRoot}/pkg/config",
    "GIN_MODE":"release",
  ```

**Utils**
- https://www.youtube.com/watch?v=Ms5RKs8TNU4&t=1504s