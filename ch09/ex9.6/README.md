Optimal value is GOMAXPROCS=4 on my 4-core mac
```
vmbp2012:gopl-solutions vince$ time GOMAXPROCS=1 go run ch08/ex8.5/ex8.5.go > /dev/null
real	0m1.152s
user	0m1.039s
sys	0m0.083s
vmbp2012:gopl-solutions vince$ time GOMAXPROCS=1 go run ch08/ex8.5/ex8.5.go > /dev/null
real	0m1.206s
user	0m1.052s
sys	0m0.087s
vmbp2012:gopl-solutions vince$ time GOMAXPROCS=2 go run ch08/ex8.5/ex8.5.go > /dev/null
real	0m0.914s
user	0m1.195s
sys	0m0.086s
vmbp2012:gopl-solutions vince$ time GOMAXPROCS=2 go run ch08/ex8.5/ex8.5.go > /dev/null
real	0m0.909s
user	0m1.194s
sys	0m0.089s
vmbp2012:gopl-solutions vince$ time GOMAXPROCS=3 go run ch08/ex8.5/ex8.5.go > /dev/null
real	0m0.868s
user	0m1.302s
sys	0m0.096s
vmbp2012:gopl-solutions vince$ time GOMAXPROCS=3 go run ch08/ex8.5/ex8.5.go > /dev/null
real	0m0.880s
user	0m1.331s
sys	0m0.101s
vmbp2012:gopl-solutions vince$ time GOMAXPROCS=4 go run ch08/ex8.5/ex8.5.go > /dev/null
real	0m0.858s
user	0m1.321s
sys	0m0.098s
vmbp2012:gopl-solutions vince$ time GOMAXPROCS=4 go run ch08/ex8.5/ex8.5.go > /dev/null
real	0m0.869s
user	0m1.374s
sys	0m0.096s
vmbp2012:gopl-solutions vince$ time GOMAXPROCS=5 go run ch08/ex8.5/ex8.5.go > /dev/null
real	0m0.893s
user	0m1.428s
sys	0m0.099s
vmbp2012:gopl-solutions vince$ time GOMAXPROCS=5 go run ch08/ex8.5/ex8.5.go > /dev/null
real	0m0.867s
user	0m1.308s
sys	0m0.102s
vmbp2012:gopl-solutions vince$ time GOMAXPROCS=6 go run ch08/ex8.5/ex8.5.go > /dev/null
real	0m0.886s
user	0m1.434s
sys	0m0.096s
vmbp2012:gopl-solutions vince$ time GOMAXPROCS=6 go run ch08/ex8.5/ex8.5.go > /dev/null
real	0m0.874s
user	0m1.404s
sys	0m0.100s
```
