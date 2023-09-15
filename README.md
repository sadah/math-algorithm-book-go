# math-algorithm-book-go

This repository is personal practices of [問題解決のための「アルゴリズム×数学」が基礎からしっかり身につく本](https://www.amazon.co.jp/dp/4297125218).

The original support site is here.

* https://github.com/E869120/math-algorithm-book

The tasks in AtCoder are here.

* https://atcoder.jp/contests/math-and-algorithm


## Require

atcoder-tools

* https://kyuridenamida.github.io/atcoder-tools/
* https://github.com/kyuridenamida/atcoder-tools/issues/263


## Usage

### Make

```
Usage:
        make build DIR=<build_directory>        Run test scripts (Specify the directory).
        make test DIR=<test_directory>          Run test scripts (Specify the directory).
        make submit DIR=<submit_directory>      Run test scripts (Specify the directory).
```

```
❯ make build DIR=./src/001
Build directory: ./src/001
[Main Program]
compile command:  go build -o main main.go
Compiling... (command: `go build -o main main.go`)
```

```
❯ make test DIR=./src/001
make: Circular test <- test dependency dropped.
Test directory: ./src/001
# in_1.txt ... PASSED 1 ms
# in_2.txt ... PASSED 1 ms
# in_3.txt ... PASSED 1 ms
Passed all test cases!!!
```

### Manual

Compile

```
atcoder-tools compile
```

Test

```
atcoder-tools test
```

Submit

```
atcoder-tools submit -u
```
