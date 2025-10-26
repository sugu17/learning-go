package exercises

import (
	"fmt"
	"maps"
	"math"
	"runtime"
	"sync"
)

var getSqrtLookup = sync.OnceValue(func() map[int]int {
	return buildSqrtLookup(10_00_000)
})

func buildSqrtLookup(maxNum int) map[int]int {
	numWorkers := runtime.NumCPU()
	results := make(map[int]int, maxNum)
	var mu sync.Mutex
	var wg sync.WaitGroup

	chunkSize := maxNum / numWorkers

	wg.Add(numWorkers)
	for i := range numWorkers {
		go func(workerID int) {
			defer wg.Done()

			start := workerID * chunkSize
			end := start + chunkSize
			if workerID == numWorkers-1 {
				end = maxNum
			}

			localMap := make(map[int]int, end-start)
			for num := start; num < end; num++ {
				localMap[num] = int(math.Sqrt(float64(num)))
			}

			mu.Lock()
			maps.Copy(results, localMap)
			mu.Unlock()
		}(i)
	}

	wg.Wait()
	return results
}

func Exercise_12_3() {
	lookup := getSqrtLookup()
	fmt.Printf("Built lookup table with %d entries\n", len(lookup))
	fmt.Printf("Sample: sqrt(100)=%d, sqrt(10000)=%d, sqrt(99999)=%d\n",
		lookup[100], lookup[10000], lookup[99999])
}
