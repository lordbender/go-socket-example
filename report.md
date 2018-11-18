# Analysis of O(n*log(n)) Shared Memory

# Analysis of O(n*log(n)) Distributed Memory

# Analysis of O(n^2) Shared Memory

# Analysis of O(n^2) Distributed Memory

# Sequential Shared Memory Runs
#Host environment: cisatlas.ccec.unf.edu#

## O(n*log(n)) Runs Outcomes

### 10000
    Report for: Linear MergeSort
            Hostname          : cisatlas.ccec.unf.edu
            Complexity        : O(n*log(n))
            Duration of Run   : 4.261944ms
            Size of test set  : 10000

### 100000
    Report for: Linear MergeSort
            Hostname          : cisatlas.ccec.unf.edu
            Complexity        : O(n*log(n))
            Duration of Run   : 40.692877ms
            Size of test set  : 100000


### 1000000
    Report for: Linear MergeSort
            Hostname          : cisatlas.ccec.unf.edu
            Complexity        : O(n*log(n))
            Duration of Run   : 429.373408ms
            Size of test set  : 1000000

## O(n^2) Run Outcomes

### 10,000
Report for: Sequential Square Matrix
        Hostname          : cisatlas.ccec.unf.edu
        Complexity        : O(n^2)
        Duration of Run   : 781.520774ms
        Size of test set  : 10000
### 20,000
Report for: Sequential Square Matrix
        Hostname          : cisatlas.ccec.unf.edu
        Complexity        : O(n^2)
        Duration of Run   : 3.010648463s
        Size of test set  : 20000
### 40,000
Report for: Sequential Square Matrix
        Hostname          : cisatlas.ccec.unf.edu
        Complexity        : O(n^2)
        Duration of Run   : 12.521956485s
        Size of test set  : 40000

# Goroutines with Channels, Threaded Shared Memory Runs
#Host environment: cisatlas.ccec.unf.edu#

## O(n*log(n)) Run Outcomes

### 10,000
    Report for: Shared Memory, Threaded Parallel MergeSort
            Hostname          : cisatlas.ccec.unf.edu
            Complexity        : O(n*log(n))
            Duration of Run   : 18.741988ms
            Size of test set  : 10000

### 100,000
    Report for: Shared Memory, Threaded Parallel MergeSort
            Hostname          : cisatlas.ccec.unf.edu
            Complexity        : O(n*log(n))
            Duration of Run   : 162.636596ms
            Size of test set  : 100000

### 1,000,000
    Report for: Shared Memory, Threaded Parallel MergeSort
            Hostname          : cisatlas.ccec.unf.edu
            Complexity        : O(n*log(n))
            Duration of Run   : 1.760324075s
            Size of test set  : 1000000

## O(n^2) Run Outcomes

### 10,000
Report for: Shared Memory, Threaded Parallel Square Matrix
        Hostname          : cisatlas.ccec.unf.edu
        Complexity        : O(n^2)
        Duration of Run   : 1.050560723s
        Size of test set  : 10000

### 20,000
Report for: Shared Memory, Threaded Parallel Square Matrix
        Hostname          : cisatlas.ccec.unf.edu
        Complexity        : O(n^2)
        Duration of Run   : 4.572166647s
        Size of test set  : 20000

### 40,000
Report for: Shared Memory, Threaded Parallel Square Matrix
        Hostname          : cisatlas.ccec.unf.edu
        Complexity        : O(n^2)
        Duration of Run   : 14.21036863s
        Size of test set  : 40000

# Client Server, HTTP Distributed Memory Runs
#Host environment: uranus.ccec.unf.edu Beowulf Cluster compute-0-[0-12] and root -> 127.0.0.1#

### 10000
### 100000
### 1000000

## O(n^2) Run Outcomes

### 10000
### 100000
### 1000000

# Atlas Environment Load at Test time
There was another user on during my testing, but they were consistently idle.

[n00599835@cisatlas oop]$ top
top - 11:55:55 up 431 days,  2:32,  1 user,  load average: 0.00, 0.00, 0.00
Tasks: 1400 total,   4 running, 1392 sleeping,   4 stopped,   0 zombie
Cpu(s):  0.6%us,  3.3%sy,  0.0%ni, 96.1%id,  0.0%wa,  0.0%hi,  0.0%si,  0.0%st
Mem:  132281152k total,  8094148k used, 124187004k free,   335656k buffers
Swap:  4194300k total,        0k used,  4194300k free,  3323276k cached

  PID USER      PR  NI  VIRT  RES  SHR S %CPU %MEM    TIME+  COMMAND
36221 n0087198  20   0  134m 2972 2140 S  5.2  0.0   1212:53 master
36259 n0087198  20   0  134m 2976 2140 S  5.2  0.0   1213:48 master
36333 n0087198  20   0  134m 2976 2140 S  5.2  0.0   1215:21 master

# Uranus Environment Load at Test time