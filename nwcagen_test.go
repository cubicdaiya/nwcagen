package main

import (
	"testing"
)

func Test_affinityEven(t *testing.T) {
	actual := "worker_cpu_affinity 1000 0100 0010 0001;"
	expected := affinityEven(4, 4)
	if actual != expected {
		t.Errorf("got %s\nwant %s", actual, expected)
	}
}
