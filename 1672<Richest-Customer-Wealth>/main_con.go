package main

import (
	"sync"
)

func maximumWealth_concurrency(accounts [][]int) int{
  biggest := 0;
  // make a channel buffer
  ch := make(chan int, len(accounts));
  wg := sync.WaitGroup{};
  
  for _,half := range accounts{
    half := half;
    //sum :=0;

    wg.Add(1);
    go func ()  {
      defer wg.Done()

      sum :=0;
      for _,num := range half{
        sum += num;
      } 

      ch <- sum;
    }()
  }

  wg.Wait();
  close(ch);

  for sum := range ch{
    if sum > biggest{
      biggest = sum;
    }
  }

  return biggest;
}


