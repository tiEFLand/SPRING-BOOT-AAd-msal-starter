OKEx open api v3 go sdk usage :
-----

### 1.Downloads or updates OKEX code's dependencies, in your command line:

```
go get -u github.com/okcoin-okex/open-api-v3-sdk/tree/master/okex-go-sdk-api
```
### 2.Write the go file. warm tips: test go file, must suffix *_test.go, eg: okex_open_api_v3_test.go
```
package gotest

import (
	"fmt"
	"github.com/okcoin-okex/open-api-v3-sdk/okex-go-sdk-api"
	"testing"
)

func TestOKExServerTime(t *testing.T) {
	serverTime, err := NewOKExClient().GetServerTime()
	