# GOLANG Coroutines, Channels and Socket Examples

# Build
    In order to build this code, you MUST put it in the correct location. GOLANG requires that 
    code be placed in one of two specific locations in order to leverage the sub packages features,
    which I have used heavily.

## Steps to build

1. unshar turnin compression somewhere
2. if go is installed on your system, you will already have a directory at ~/go/src/
    1. This is the root for all go programs for your profile.
3. cp the oop directory in the shar to the following path 
    1. ~/go/src/oop
    2. Such that the main.go is at ~/go/src/oop/main.go
4. cd ~/go/src/oop
5. go build main.go

# Flags
    --size
    --all
    --nlogn
    --nsquared
    --rocks => if set, will attempt to distribute to the rocks cluster

# Test

## Test Execution
    $ chmod 777 test.sh
    $ ./test.sh
## Expected Test Output
    2018/11/10 13:47:21 Sequential Merge Sort took 997.3µs
    2018/11/10 13:47:21 Parallel Merge Sort took 3.0005ms

# Run Sequential (Shared Memory -> Multi Threaded (Goroutines and Buffered Channels))

## Shared Memory Flag Permutations
1. go run main.go --size=10000 --all
2. go run main.go --size=10000 --nsquared
3. go run main.go --size=10000 --nlogn

## Distributed Memory Flag Permutations

#NOTE: Server must be started seperatly to run these commands see Below Section on Running Distributed Memory Server#
1. go run main.go --size=10000 --all --rocks
2. go run main.go --size=10000 --nsquared --rocks
3. go run main.go --size=10000 --nlogn --rocks

# Running Distributed Memory

1. go run server.go
2. go run main.go --size=10000 --nlogn --rocks1

