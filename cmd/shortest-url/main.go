package main

import (
	"fmt"

	"github.com/cmczk/shortest-url/internal/config"
)

func main() {
	cfg := config.MustLoad()

	fmt.Println(cfg)
}
