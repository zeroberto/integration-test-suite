# integration-test-suite
Project with routines that assist in the execution of integration tests.

## Requirements
To use all the basic features of this project, it is necessary to have the docker installed, as well as the docker-compose.

## Import
Import this project into your go project using the following statement:

```go
import "github.com/zeroberto/integration-test-suite"
```

Run the command below to download the lib:

```bash
go get github.com/zeroberto/integration-test-suite
```

## Usage
Create a docker-compose file containing your infrastructure services:

```yml
# Example of docker-compose.yml
version: "3.8"

services:
  db-test:
    image: postgres
    environment: 
      POSTGRES_USER: test
      POSTGRES_PASSWORD: test
      POSTGRES_DB: test
    volumes:
      - ./resources/sql:/docker-entrypoint-initdb.d
    ports: 
      - 65432:5432
    networks: 
      - test-network
  
networks:
  test-network:
    name: test-network
    driver: bridge
```

Get your infrastructure up before running your tests. At the end of the tests, your infra must be dropped. Your tests must be performed between these different moments. Below is a complete example of how your test file might look partially:

```go
// Example of test file
import (
  "os"
  "testing"
  "time"

  infra "github.com/zeroberto/integration-test-suite"
)

const (
  host                  string = "localhost"
  serverPort            string = "7777"
  dbType                string = "postgres"
  dataSourceName        string = "postgres://test:test@localhost:65432/test?sslmode=disable"
  dockerComposeFileName string = "docker-compose.yml"
)

func TestFindUserInfo(t *testing.T) {
  // Your test code
}

func TestMain(m *testing.M) {
  setup()
  code := m.Run()
  teardown()
  os.Exit(code)
}

func setup() {
  infra.DownInfra(dockerComposeFileName)
  infra.UpInfra(dockerComposeFileName)
  go initServer()
  validateStructure(dockerComposeFileName)
}

func teardown() {
  infra.DownInfra(dockerComposeFileName)
}

func initServer() {
  // Start your server in a personalized way
}

func validateStructure(dockerComposeFileName string) {
  // Validate that all the conditions of your infra are ok. This is just an example.
  for {
    if infra.CheckPortIsOpen(host, serverPort) && infra.CheckDBConnection(dbType, dataSourceName) == nil {
      return
    }
    time.Sleep(100 * time.Millisecond)
  }
}
```

## License

[MIT](LICENSE) License
