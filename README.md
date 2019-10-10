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
Running 30s test @ http://localhost:8091/65
  10 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     4.79ms    5.61ms  76.52ms   85.10%
    Req/Sec     3.11k   573.88     6.59k    69.63%
  930920 requests in 30.07s, 105.65MB read
Requests/sec:  30954.22
Transfer/sec:      3.51MB
```

* Python+Japronto
```
Running 30s test @ http://localhost:8090/65
  10 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     5.49ms    1.89ms  20.20ms   79.08%
    Req/Sec     1.83k   337.26    15.61k    87.50%
  547649 requests in 30.10s, 89.31MB read
Requests/sec:  18196.82
Transfer/sec:      2.97MB
```

### Thanks!
