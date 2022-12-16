Correct Identifier

10 bytes or 16 characters

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

//# BytesIDToStringID
}
```

```javascript
// js
let cid = new CID();
let id = cid.generate()
document.querySelector('#cid').innerHTML = id;
console.log(cid.getTime(id))
```