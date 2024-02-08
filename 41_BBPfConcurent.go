package main

import (
	"fmt"
	"math"
	"math/big"
	"sync"
	"time"
)

// case 41: // -- AMFBBPfConcurentA
func BBPfConcurent(selection int) {
	usingBigFloats = false
	workerCount := 4
	prescribedIterations := 8

	fmt.Println("Enter the desired number of workers (recomend 2 to 16, 4 is optimal) ")
	fmt.Scan(&workerCount)

	fmt.Println("Enter your number of iterations (recomend 2 to 16, 8 being optimal) ")
	fmt.Scan(&workerCount)

	start := time.Now()

	pi := PiMultiThread(workerCount, prescribedIterations)

	t := time.Now()
	elapsed := t.Sub(start)
	TotalRun := elapsed.String()
	fmt.Println("elapsed time was: ", TotalRun)

	piAsBF := new(big.Float)
	piAsBF = big.NewFloat(pi) // pi is being cast to big from float64
	printResultStatsLong(piAsBF, prescribedIterations, "BBPfConcurent", workerCount, TotalRun, selection)
}

// global structure
type Float64 struct {
	value float64
	lock  sync.RWMutex
}

// helper functions:
func (f *Float64) Inc(diff float64) float64 {
	f.lock.Lock()
	ret := f.value + diff
	f.value = ret
	f.lock.Unlock()
	return ret
}
func (f *Float64) Get() float64 {
	f.lock.RLock()
	ret := f.value
	f.lock.RUnlock()
	return ret
}
func partialSum(kStart int, kOffset int, amount int) (sum float64) {
	for k := float64(kStart); k < float64(kStart+kOffset*amount); k += float64(kOffset) {
		sum += 1 / math.Pow(16, k) * (4/(8*k+1) - 2/(8*k+4) - 1/(8*k+5) - 1/(8*k+6))
	}

	return
}
func PiMultiThread(workers int, iteration int) float64 {
	wg := sync.WaitGroup{}
	ret := Float64{}

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go (func(kStart int) {
			ret.Inc(partialSum(kStart, workers, iteration))
			// fmt.Println(ret)
			wg.Done()
			// fmt.Println(ret) // Rick's progress report, one line for each worker ?
			// fmt.Printf("\nWorker number %d came up with: ", workers) // Rick's progress report
			// fmt.Print(ret)
			// fmt.Print("\n")
		})(i)
		fmt.Printf("\nWorkers numbering %d came up with: ", workers) // Rick's progress report
		fmt.Print(ret)
		fmt.Print("\n")
	}
	wg.Wait()
	return ret.Get()
	// adapted by Richard Woolley
} // end of func BBPfConcurent() set // -- AMFBBPfConcurentB
