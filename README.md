# Benchmark for fun

The purpose of this benchmark is compare the Request per seconds of 2 language+frameworks: `Go+Gin` and `Python+Japronto`

### Techs
* Go language + Gin https://github.com/gin-gonic/gin 
* Python + Japronto https://github.com/squeaky-pl/japronto
* Wrk benchmark tool https://github.com/wg/wrk

### Running Go+Gin

* Runs at 8091

```
$ go build ginmain.go
$ ./ginmain
```

### Running Python+Japronto

* Runs at 8090

``` 
$ pip3 install japronto --user
$ python3 moes-app.py
```

### Benchmarking

* For Go+Gin, issue the following command:
```
$ wrk -c100 -t10 -d30s -T1s http://localhost:8091/65
```

* For Python+Japronto, issue the following command:
```
$ wrk -c100 -t10 -d30s -T1s http://localhost:8090/65
```

### Results @ my machine: (Dell i5 8gen + 16GB Ram)

* Go+Gin
```
Running 10s test @ http://localhost:8091/65
  10 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     0.92ms    1.10ms  20.89ms   89.43%
    Req/Sec    13.81k     3.44k   21.08k    64.60%
  1377075 requests in 10.04s, 156.28MB read
Requests/sec: 137183.40
Transfer/sec:     15.57MB
```

* Python+Japronto
```
Running 10s test @ http://localhost:8090/65
  10 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     2.24ms    1.14ms  21.18ms   89.82%
    Req/Sec     4.67k     1.26k    5.79k    75.50%
  466592 requests in 10.09s, 76.09MB read
Requests/sec:  46220.64
Transfer/sec:      7.54MB
```

### Thanks!
