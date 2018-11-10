# go-socket-example

# Build

go build main.go

# Run

go run main.go --size=10000 --all
go run main.go --size=10000 --nsquared
go run main.go --size=10000 --nsquared

## Flags
    --size
    --all
    --nlogn
    --nsquared
# Test

\$ ./test.sh
2018/11/10 13:47:21 Sequential Merge Sort took 997.3µs
2018/11/10 13:47:21 Parallel Merge Sort took 3.0005ms
