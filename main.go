package main

import (
	"fmt"

	"github.com/Dhliv/Gometry/src/models"
)

func main() {
	p1 := models.NewPoint(3, 30)
	p2 := models.NewPoint(3, 30)
	fmt.Println(p1.Equal(p2))
}
