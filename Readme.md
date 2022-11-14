### About

Library golang sederhana untuk get data dari REST API Showroom.

### Install

```sh
go get github.com/nekoding/showroom
```

### Usage

```go
package main

import (
	"fmt"

	"github.com/nekoding/showroom"
)

func main() {
	shaniRoomId := "318207"
	result := showroom.GetLiveinfo(shaniRoomId)
	fmt.Println(result.LiveStatus)
}

```