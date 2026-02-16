rund:
    go run ./cmd/taskmasterd/ --conf=./config/taskmaster.yaml --log-level=debug

runctl:
    go run ./cmd/taskmasterctl/ --log-level=debug

lint:
    golangci-lint run
fmt:
    golangci-lint fmt
