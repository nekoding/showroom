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
