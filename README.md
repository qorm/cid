# Correct Identifier

## 10 bytes (timestamp 6 bytes,random number 4 bytes)
## 16 characters (timestamp 9 characters,random number 7 characters)
### $\color{#6666cc}{0LBQ98NYA}$ $\color{#66cc66}{1VKNOJ0}$
### $\color{#6666cc}{0LBQ98NYA}$ timestamp
### $\color{#66cc66}{1VKNOJ0}$ random number

``` 
goos: darwin
goarch: arm64
pkg: github.com/qorm/cid
BenchmarkBytesID-10                   	44514168	        26.48 ns/op	1208.40 MB/s	      16 B/op	       2 allocs/op
BenchmarkStringID-10                  	 3046996	       393.0 ns/op	  81.42 MB/s	     112 B/op	      18 allocs/op
BenchmarkStringIDToBytesID-10         	 6399034	       203.7 ns/op	 157.07 MB/s	      16 B/op	       1 allocs/op
BenchmarkBytesIDToStringID-10         	 3241030	       372.4 ns/op	  85.94 MB/s	      96 B/op	      17 allocs/op
BenchmarkGetTimestampByStringID-10    	12107529	        99.58 ns/op	 321.35 MB/s	       0 B/op	       0 allocs/op
BenchmarkGetTimestampByBytesID-10     	1000000000	         0.3123 ns/op	102458.62 MB/s	       0 B/op	       0 allocs/op
PASS
ok  	github.com/qorm/cid	7.539s
```

```golang
package main
//go
import (
    "time"
    "github.com/qorm/cid"
）
func main(){
// NewBytesID
//example:  cid.NewBytesID(time.Now())
// NewStringID
//example:  cid.NewStringID(time.Now())

// GetTimestampByBytesID

// GetTimestampByStringID

// StringIDToBytesID

// BytesIDToStringID
}
```

```javascript
// js
let cid = new CID();
let id = cid.generate()
console.log(id);
console.log(cid.getTime(id))
console.log(cid.getTimestamp(id))
```