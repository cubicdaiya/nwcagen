package main

import (
	"bytes"
	"flag"
	"fmt"
	"runtime"
)

func affinityEven(n int) string {
	var buffer bytes.Buffer
	buffer.WriteString("worker_cpu_affinity ")
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				buffer.WriteString("1")
			} else {
				buffer.WriteString("0")
			}
		}
		if i != n-1 {
			buffer.WriteString(" ")
		}
	}
	buffer.WriteString(";")
	return buffer.String()
}

func main() {
	cpuNum := runtime.NumCPU()
	workerNum := flag.Int("n", cpuNum, "worker number.")
	bias := flag.String("b", "even", "cpu bias. default: even")
	flag.Parse()

	if *bias == "even" && *workerNum == cpuNum {
		fmt.Println(affinityEven(*workerNum))
	} else {
		fmt.Println("not supported situation")
	}
}
