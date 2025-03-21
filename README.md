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
    KeycloakAuth((Keycloak Auth))

    subgraph "Shared Wallet Service"
        subgraph "API Layer"
            APIServer["API Server<br>Gin/Go"]
            
            subgraph "API Handlers"
                PingHandler["Ping Handler<br>Go"]
                BudgetHandler["Budget Handler<br>Go"]
                TeamHandler["Team Handler<br>Go"]
            end

            subgraph "URL Router"
                GinRouter["Router<br>Gin"]
            end
        end

        subgraph "Use Cases"
            BudgetUseCase["Budget Use Case<br>Go"]
            TeamUseCase["Team Use Case<br>Go"]
            UserUseCase["User Use Case<br>Go"]
        end

        subgraph "Domain Layer"
            BudgetDomain["Budget Domain<br>Go"]
            TeamDomain["Team Domain<br>Go"]
            UserDomain["User Domain<br>Go"]
        end

        subgraph "Infrastructure"
            subgraph "Database"
                MariaDB[("MariaDB<br>MariaDB")]
                DBReader["DB Reader<br>Go SQL"]
                DBWriter["DB Writer<br>Go SQL"]
            end

            subgraph "Cache System"
                CacheClient["Cache Client<br>Go"]
                InMemoryCache[("In-Memory Cache<br>Go Sync.Map")]
            end

            subgraph "External Clients"
                KeycloakClient["Keycloak Client<br>Go"]
                RestClient["REST Client<br>Go"]
            end

            subgraph "Jobs"
                JobHandler["Job Handler<br>Go"]
                CacheLoader["Cache Loader<br>Go"]
            end
        end
    end

    %% External connections
    User -->|"HTTP Requests"| APIServer
    APIServer -->|"Authenticates"| KeycloakAuth
    KeycloakClient -->|"Fetches Users"| KeycloakAuth

    %% API Layer connections
    APIServer -->|"Routes"| GinRouter
    GinRouter -->|"Handles"| PingHandler
    GinRouter -->|"Handles"| BudgetHandler
    GinRouter -->|"Handles"| TeamHandler

    %% Handler to Use Case connections
    BudgetHandler -->|"Uses"| BudgetUseCase
    TeamHandler -->|"Uses"| TeamUseCase
    JobHandler -->|"Uses"| UserUseCase

    %% Use Case to Domain connections
    BudgetUseCase -->|"Uses"| BudgetDomain
    TeamUseCase -->|"Uses"| TeamDomain
    UserUseCase -->|"Uses"| UserDomain

    %% Domain to Infrastructure connections
    BudgetDomain -->|"Reads/Writes"| DBReader
    BudgetDomain -->|"Reads/Writes"| DBWriter
    TeamDomain -->|"Reads/Writes"| DBReader
    TeamDomain -->|"Reads/Writes"| DBWriter
    UserDomain -->|"Caches"| CacheClient

    %% Infrastructure internal connections
    DBReader -->|"Queries"| MariaDB
    DBWriter -->|"Updates"| MariaDB
    CacheClient -->|"Stores"| InMemoryCache
    JobHandler -->|"Manages"| CacheLoader
    CacheLoader -->|"Updates"| CacheClient
    KeycloakClient -->|"Uses"| RestClient

    %% Style definitions
    classDef container fill:#e6e6e6,stroke:#333,stroke-width:2px
    classDef component fill:#fff,stroke:#333,stroke-width:1px
    classDef external fill:#d7d7d7,stroke:#333,stroke-width:2px
    
    class APIServer,Database container
    class PingHandler,BudgetHandler,TeamHandler,GinRouter,BudgetUseCase,TeamUseCase,UserUseCase,KeycloakClient,RestClient,JobHandler,CacheLoader component
    class User,KeycloakAuth external
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