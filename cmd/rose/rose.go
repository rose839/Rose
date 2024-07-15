package main

import (
	"os"
	"runtime"
	"time"

	"github.com/rose839/Rose/internal/rose"
	"golang.org/x/exp/rand"
)

func main() {
	rand.Seed(uint64(time.Now().UTC().UnixNano()))
	if len(os.Getenv("GOMAXPROCS")) == 0 {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}

	rose.NewApp("rose").Run()
}
