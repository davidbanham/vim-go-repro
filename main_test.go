package main

import (
	"flag"
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	flag.Parse()
	os.Exit(m.Run())
}

func TestExample(t *testing.T) {
	fmt.Println("Test is passing")
}
