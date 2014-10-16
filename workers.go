package main

import "fmt"
import "time"
import "math/rand"

func worker(id int, randomSleep bool, jobs <-chan int, results chan<- int) {
  for j := range jobs {
    fmt.Println("worker", id, "processing job", j)

    delayFactor := 1
    if randomSleep {
      delayFactor = rand.Intn(4)
    }

    time.Sleep(time.Second * time.Duration(delayFactor))
    results <- j * 2
  }
}

func main() {
  noOfJobs := 9
  randomDelayFlag := false

  jobs := make(chan int, 100)
  results := make(chan int, 100)

  for w := 1; w <= 3; w++ {
    go worker(w, randomDelayFlag, jobs, results)
  }

  for j := 1; j <= noOfJobs; j++ {
    jobs <- j
  }
  close(jobs)

  for a := 1; a <= noOfJobs; a++ {
    <-results
  }

}
