# Threaded Operation (Shared Memory)

seanw@LAPTOP-O33V4R1R MINGW64 ~/go/src/oop (master)
$ go run main.go 20000
2018/11/10 14:47:22 Sequential Merge Sort took 2.9913ms
2018/11/10 14:47:22 Parallel Merge Sort took 3.9778ms
2018/11/10 14:47:33 Sequential Square Matrix took 1.4840293s
2018/11/10 14:47:44 Parallel Square Matrix took 852.721ms

# Socket Operation (Distributed Memory)

// start the server on all clusters

# For Each Cluster
lsof -i -P -n | grep LISTEN
rocks run host compute-0-0 "~/go/src/oop/server"
rocks iterate host "~/go/src/oop/server &"
curl http://localhost:8080/api/v1/running