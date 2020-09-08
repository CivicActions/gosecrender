package main

import (
	"fmt"

	"github.com/tom-camp/gossptk/internal/pkg/opencontrol"
)

var oc opencontrol.OpenControl

func main() {
	oc.LoadOpenControl()
	fmt.Println(oc.Metadata)
}
