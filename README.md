# Correct Identifier

## 10 bytes or 16 characters

### 2059-5-26 1:38:27

``` 
goos: darwin
goarch: arm64
pkg: github.com/qorm/cid
BenchmarkBytesID-10                	44344735	        26.39 ns/op	1212.63 MB/s	      16 B/op	       2 allocs/op
BenchmarkStringID-10               	 3065005	       391.8 ns/op	  81.67 MB/s	     112 B/op	      18 allocs/op
BenchmarkStringIDToBytesID-10      	 6183050	       212.2 ns/op	 150.80 MB/s	      16 B/op	       1 allocs/op
BenchmarkBytesIDToStringID-10      	 3180595	       377.7 ns/op	  84.72 MB/s	     112 B/op	      17 allocs/op
BenchmarkStringIDToTimestamp-10    	13416471	        89.54 ns/op	 357.38 MB/s	       0 B/op	       0 allocs/op
BenchmarkBytesIDToTimestamp-10     	1000000000	         0.3117 ns/op	102675.54 MB/s	       0 B/op	       0 allocs/op
PASS
ok  	github.com/qorm/cid	7.534s
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

// BytesIDToTimestamp

// StringIDToTimestamp

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
```