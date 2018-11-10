# go-socket-example

# Build

go build main.go

# Run

## Shared Memory Permutations
go run main.go --size=10000 --all
go run main.go --size=10000 --nsquared
go run main.go --size=10000 --nlogn

## Distributed Memory Permutations
go run main.go --size=10000 --all --rocks
go run main.go --size=10000 --nsquared --rocks
go run main.go --size=10000 --nlogn --rocks

## Flags
    --size
    --all
    --nlogn
    --nsquared
    --rocks => if set, will attempt to distribute to the rocks cluster
# Test

\$ ./test.sh
2018/11/10 13:47:21 Sequential Merge Sort took 997.3µs
2018/11/10 13:47:21 Parallel Merge Sort took 3.0005ms
