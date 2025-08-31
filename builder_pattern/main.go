package main

import (
	"fmt"
	pcbuilder "go-design-patterns/builder_pattern/pc_builder"
	"go-design-patterns/builder_pattern/pc_builder/pc"
)

func main() {
	pcBuilder := pc.NewPcBuilder()
	director := pcbuilder.NewDirector(pcBuilder)
	fmt.Printf("Gaming PC: %+v\n", director.BuildGamingPc())
	fmt.Printf("Office PC: %+v\n", director.BuildOfficePc())
}
