// Testing some stuff...

package main

import (
	"fmt"
	"sync"
	"time"
)

// Produces jobs to later perform (by worker())
func producer(jobs chan<- int, nums []int) {
	// Loop through all the numbers (job ids)
	for _,  n := range nums {
		// Push those job ids into the jobs channel.
		// We can do this since the producer() is call with the `go` keywork before
		// it.
		jobs <- n // ajoute ce job au channel des jobs
	}
	// We close the channel. Not sure why though.
	close(jobs)
}

// Performs the jobs created by producer()
// <-chan means channel can give value
// chan<- means channel can receive value
func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done() // will be executed when the function ends
	// `range jobs` waits for jobs to be available to process them (pop them out
	// of the queue):
	for n := range jobs {
		fmt.Printf("worker %d processing %d\n", id, n)
		time.Sleep(50 * time.Millisecond) // simulate work
		// adds the exponent result to the other channel: results
		results <- n * n
	}
}


func main() {
	const numWorkers = 4

	jobs := make(chan int, 10)
	results := make(chan int, 10)

	var wg sync.WaitGroup

	for i := 1; i <= numWorkers; i++ {
		// wg.Add(1)
		// go worker(i, jobs, results, &wg)
		wg.Go(func() { // THIS DOESNT WORK WHEN numWorkers = 4 AND THERE ARE NOT 10 ELEMENTS IN THE jobs CHANNEL
			worker(i, jobs, results, &wg)
		})
	}

	go func() {
		wg.Wait() // arrete le waitgroup, si non ca va pas marcher. ca va closer le results channel avant de range over them.
		close(results) // ET ferme le channel des results
	}()

	// genere les jobs (channel)
	go producer(jobs, []int{1,2,3,4,5,6,7,8,9,10})

	// debloque/consomme le channel de results avec `range`:
	for result := range results {
		fmt.Println("results:", result)
	}
}

