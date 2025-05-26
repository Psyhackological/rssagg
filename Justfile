alias r := run
alias cr := container_run

run:
  go build && ./rssagg

# Get a Go package and clean up modules
goget package:
    go get {{package}}
    go mod tidy
    go mod vendor

# Load environment variables from .env file
set dotenv-load

# Container runtime (docker or podman)
runtime := env_var_or_default("CONTAINER_RUNTIME", "podman")
postgres_container := env_var_or_default("POSTGRES_CONTAINER", "my-postgres")
postgres_port := env_var_or_default("POSTGRES_PORT", "5432")
postgres_volume := env_var_or_default("POSTGRES_VOLUME", "postgres_data")

container_run:
    {{runtime}} run -d \
        --name {{postgres_container}} \
        --env-file .env \
        -p {{postgres_port}}:5432 \
        -v {{postgres_volume}}:/var/lib/postgresql/data \
        docker.io/library/postgres:15

# Stop the container
stop:
    {{runtime}} stop {{postgres_container}}

# Force use Docker
docker-run:
    docker run -d \
        --name {{postgres_container}} \
        --env-file .env \
        -p {{postgres_port}}:5432 \
        -v {{postgres_volume}}:/var/lib/postgresql/data \
        docker.io/library/postgres:15

# Force use Podman
podman-run:
    podman run -d \
        --name {{postgres_container}} \
        --env-file .env \
        -p {{postgres_port}}:5432 \
        -v {{postgres_volume}}:/var/lib/postgresql/data \
        docker.io/library/postgres:15
