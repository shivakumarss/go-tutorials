package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)

	go func() {

		for {
			k := <-ch
			if k == "end" {
				fmt.Println("about to break the loop")
				break
			}

			fmt.Println(k)

		}

		//for m := range ch {
		//	fmt.Println("processed:", m)
		//}
	}()

	ch <- "cmd.1"
	ch <- "cmd.2"
	ch <- "cmd.4"
	ch <- "cmd.3" //won't be processed
	ch <- "cmd.5" //won't be processed
	ch <- "end"   //won't be processed

	//time.Sleep(time.Duration(10) * time.Second)
	//time.Sleep(time.Duration(10000) )
	fmt.Println("Done")
}
