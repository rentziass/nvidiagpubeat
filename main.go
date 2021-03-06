package main

import (
	"os"

	"github.com/elastic/beats/libbeat/beat"

	"github.com/deepujain/nvidiagpubeat/beater"
)

func main() {
	err := beat.Run("nvidiagpubeat", "", beater.New)
	if err != nil {
		os.Exit(1)
	}
}
