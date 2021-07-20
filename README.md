# test-task-go
An app which gets last 100 blocks from etherscan.io and finds an address the balance of which has been changed most.

How balance change is counted: 
```
addr: ... +x1+x2-x3-x4 -> change = |x1+x2-x3-x4|
```

## Download and run

```
$ git clone https://github.com/mishkamashka/test-task-go
$ cd test-task-go/src
$ go run .
```
