# grumpy-trial
```
make
export GOPATH=$PWD/build
export PYTHONPATH=$PWD/build/lib/python2.7/site-packages
tools/grumpc hello.py > hello.go
go build -o hello hello.go
./hello
```
