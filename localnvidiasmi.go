package main

import (
	"fmt"
)

func main() {
	header := "utilization.gpu [%], utilization.memory [%], memory.total [MiB], memory.free [MiB], memory.used [MiB], temperature.gpu, pstate"
	gpu := [4]string{"1 %, 10 %, 6082 MiB, 6082 MiB, 0 MiB, 29, P8",
                    	"2 %, 20 %, 6082 MiB, 6082 MiB, 0 MiB, 32, P8",
			"3 %, 30 %, 6082 MiB, 6082 MiB, 0 MiB, 31, P8",
			"4 %, 40 %, 6082 MiB, 6082 MiB, 0 MiB, 27, P8"}
	
	fmt.Println(header)
	for _, gpuUtil := range gpu {
	   fmt.Println(gpuUtil)	
	}
}

