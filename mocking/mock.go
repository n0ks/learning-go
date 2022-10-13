package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const start = 3

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := start; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}

	fmt.Fprint(out, "Go!")
}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}

	Countdown(os.Stdout, sleeper)
}
