# DockerGoProject

# Correr test unitarios

export CONFIG_DIR=/Users/andortiz/Documents/GO/src/Examples/DockerGoProject/pkg/config && export SCOPE=local && go test -v ./... -covermode=atomic -coverprofile=coverage.out -coverpkg=./... -count=1

# Mirar resultado

go tool cover -html=coverage.out

# Correr tests para crear automaticamente

mockery --all --disable-version-string

# Notas

Se agrega DB con configuracion por usuario test_R y test_W
