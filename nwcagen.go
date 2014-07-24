package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"runtime"
)

func evenSetByCPU(cpuNum int) []string {
	var buffer bytes.Buffer
	var result []string
	n := cpuNum
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				buffer.WriteString("1")
			} else {
				buffer.WriteString("0")
			}
		}
		result = append(result, buffer.String())
		buffer.Reset()
	}
	return result
}

func affinityEven(workerNum, cpuNum int) string {
	var buffer bytes.Buffer
	evenSet := evenSetByCPU(cpuNum)
	n := workerNum / cpuNum
	m := workerNum % cpuNum
	buffer.WriteString("worker_cpu_affinity ")
	for i := 0; i < n; i++ {
		for j, even := range evenSet {
			buffer.WriteString(even)
			if i != n-1 || j != cpuNum-1 {
				buffer.WriteString(" ")
			}
		}
	}

	if m > 0 {
		buffer.WriteString(" ")
		for i := 0; i < m; i++ {
			buffer.WriteString(evenSet[i])
			if i != m-1 {
				buffer.WriteString(" ")
			}
		}
	}

	buffer.WriteString(";")
	return buffer.String()
}

func main() {
	workerNum := flag.Int("n", runtime.NumCPU(), "worker number.")
	cpuNum := flag.Int("c", runtime.NumCPU(), "CPU number( This is a parameter for debugging).")
	bias := flag.String("b", "even", "CPU bias. default: even")
	flag.Parse()

	if *bias == "even" {
		if (*workerNum % *cpuNum) != 0 {
			log.Println("[warn]Workers are not assigned to CPUs evenly. Because worker number is indivisible by CPUs.")
		}
		fmt.Println(affinityEven(*workerNum, *cpuNum))
	} else {
		fmt.Println("not supported situation")
	}
}
