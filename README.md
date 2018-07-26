# TEST PERFORMANCE GO

What doing here is to test performance script golang and efficiency statement we use for daily coding.


## IF vs SWITCH

to run file, going to directory `switchvsif` and run this command
```
go test -bench . -benchtime 30s
```
the result will you get

```
go test -bench . -benchtime 30s
1
goos: darwin
goarch: amd64
pkg: github.com/suryatresna/goperformance/switchvsif
BenchmarkSwitchString-4   	100
10000
1000000
100000000
3000000000
3000000000	        17.5 ns/op
1
BenchmarkIfString-4       	100
10000
1000000
100000000
3000000000
3000000000	        15.2 ns/op
1
BenchmarkSwitchInt-4      	100
10000
1000000
100000000
5000000000
5000000000	        13.0 ns/op
1
BenchmarkIfInt-4          	100
10000
1000000
100000000
2000000000
2000000000	        12.9 ns/op
PASS
ok  	github.com/suryatresna/goperformance/switchvsif	194.933s
```

What going on here, we can see when use `switch` for validate string is more lower than use `if`, but it's different when use for numeric `(int)` that's fast than `if`.