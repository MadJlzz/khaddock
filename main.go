package main

import (
	"fmt"

	"github.com/MadJlzz/khaddock/engine"
	"github.com/MadJlzz/khaddock/sailor"
)

func main() {
	var sailorSpec engine.SailorConfiguration
	err := engine.Unmarshal("examples/vm.yaml", &sailorSpec)
	if err != nil {
		panic(err)
	}
	fmt.Println(sailorSpec)

	for k, v := range sailorSpec.Specs.SysctlConfiguration {
		fmt.Printf("Setting key %s with value %s\n", k, v)
		err = sailor.Set(k, v)
		if err != nil {
			panic(err)
		}
	}

}
