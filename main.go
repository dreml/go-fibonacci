package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server is running")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	params, ok := r.URL.Query()["count"]
	if !ok {
		fmt.Fprint(w, "Count parameter is missing")
		return
	}
	count, err := strconv.Atoi(params[0])
	if err != nil {
		fmt.Fprintf(w, "Count is invalid")
		return
	}

	for i := range fibonacci(count) {
		fmt.Fprintln(w, i)
	}
}

func fibonacci(count int) chan int {
	ch := make(chan int)
	go func(ch chan int) {
		prev := 0
		cur := 1
		for i := 0; i <= count; i++ {
			fib := prev + cur

			ch <- fib

			prev = cur
			cur = fib
		}
		close(ch)
	}(ch)

	return ch
}
