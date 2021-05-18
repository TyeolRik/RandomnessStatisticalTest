package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"runtime"

	"github.com/ivpusic/grpool"
)

func MaxParallelism() int {
	maxProcs := runtime.GOMAXPROCS(0)
	numCPU := runtime.NumCPU()
	if maxProcs < numCPU {
		return maxProcs
	}
	return numCPU
}

func main() {
	maxProcessor := MaxParallelism()
	runtime.GOMAXPROCS(maxProcessor)
	fmt.Println("MaxParallelism: ", maxProcessor)

	// number of workers, and size of job queue
	pool := grpool.NewPool(100, 100)
	defer pool.Release()

	// how many jobs we should wait
	pool.WaitCount(100)

	// submit one or more jobs to pool
	for i := 0; i < 100; i++ {
		newRandomInt, _ := rand.Int(rand.Reader, big.NewInt(1000000))
		count := i
		pool.JobQueue <- func() {
			// say that job is done, so we can know how many jobs are finished
			defer pool.JobDone()

			fmt.Println("hello ", count, " ", newRandomInt)
		}
	}

	// wait until we call JobDone for all jobs
	pool.WaitAll()
}
