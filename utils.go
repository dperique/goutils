package goutils

import (
	"fmt"
	"log"
	"os"
	"sync"
)

func CheckErrFatal(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func AskFirst() {
	fmt.Print("Enter 'yes' to continue: ")
	var response string
	_, err := fmt.Scan(&response)
	CheckErr(err)
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

func (c *Counter) Incr() {
	c.m.Lock()
	defer c.m.Unlock()
	c.val++
}

func (c *Counter) GetVal() int {
	c.m.Lock()
	defer c.m.Unlock()
	return c.val
}
