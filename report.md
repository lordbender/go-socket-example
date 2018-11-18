# Threaded Operation (Shared Memory)

seanw@LAPTOP-O33V4R1R MINGW64 ~/go/src/oop (master)
$ go run main.go 20000
2018/11/10 14:47:22 Sequential Merge Sort took 2.9913ms
2018/11/10 14:47:22 Parallel Merge Sort took 3.9778ms
2018/11/10 14:47:33 Sequential Square Matrix took 1.4840293s
2018/11/10 14:47:44 Parallel Square Matrix took 852.721ms

# Socket Operation (Distributed Memory)

// start the server on all clusters

# Setup and Clean-Up For Each Cluster

## if a server goes down, use the correct command to get it running again
rocks run host compute-0-0 "~/go/src/oop/server" &
rocks run host compute-0-1 "~/go/src/oop/server" &
rocks run host compute-0-2 "~/go/src/oop/server" &
rocks run host compute-0-3 "~/go/src/oop/server" &
rocks run host compute-0-4 "~/go/src/oop/server" &
rocks run host compute-0-5 "~/go/src/oop/server" &
rocks run host compute-0-6 "~/go/src/oop/server" &
rocks run host compute-0-7 "~/go/src/oop/server" &
rocks run host compute-0-8 "~/go/src/oop/server" &
rocks run host compute-0-9 "~/go/src/oop/server" &
rocks run host compute-0-10 "~/go/src/oop/server" &
rocks run host compute-0-11 "~/go/src/oop/server" &
rocks run host compute-0-12 "~/go/src/oop/server" &

## if a port is in use, use lsof to find and kill it.
lsof -i -P -n | grep LISTEN

## Test whether or not a specific compute cluster us up.
curl http://localhost:8080/api/v1/running
curl http://compute-0-0:8080/api/v1/running
curl http://compute-0-[0-12]:8080/api/v1/running

## find all PORTS in use on the cluster if needed
rocks iterate host "lsof -i -P -n | grep LISTEN"

// kill them
rocks iterate host "kill ???"