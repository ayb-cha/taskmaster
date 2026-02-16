rund:
    go run ./cmd/taskmasterd/ --conf=./config/config.yaml --log-level=debug

runctl:
    go run ./cmd/taskmasterctl/ --conf=./config/config.yaml --log-level=debug

lint:
    golangci-lint run
fmt:
    golangci-lint fmt
