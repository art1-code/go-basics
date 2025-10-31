package main

import (
	"my_project_math/types"
)

func main() {
	c := &types.Calculadora{}
	m := types.NewMediator(c)
	m.Execute()
}
