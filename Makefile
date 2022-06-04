golangci := go run github.com/golangci/golangci-lint/cmd/golangci-lint@v1.46.2
benchstat := go run golang.org/x/perf/cmd/benchstat@v0.0.0-20220411212318-84e58bfe0a7e
benchart := go run github.com/storozhukBM/benchart@v1.0.0

bench_hash:
	go test -timeout 3h -count=5 -run=xxx -bench=BenchmarkHash ./... | tee hash_stat.txt
	$(benchstat) hash_stat.txt
	$(benchstat) -csv hash_stat.txt > hash_stat.csv
	$(benchart) 'Hash;xAxisType=log;yAxisType=log' hash_stat.csv hash_stat.html
	open hash_stat.html

test:
	go test ./...

lint:
	$(golangci) run
