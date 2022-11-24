package utils

import (
	"fmt"
	"log"
	"os"
	"sync"
)

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func askFirst() {
	fmt.Print("Enter 'yes' to continue: ")
	var response string
	_, err := fmt.Scan(&response)
	checkErr(err)
	if response == "yes" {
		return
	}
	fmt.Println("Delete is aborted.")
	os.Exit(1)
}

type Counter struct {
	m   sync.Mutex
	val int
}

func (c *Counter) incr() {
	c.m.Lock()
	defer c.m.Unlock()
	c.val++
}

func (c *Counter) getVal() int {
	c.m.Lock()
	defer c.m.Unlock()
	return c.val
}
