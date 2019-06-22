package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"strings"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for _, path := range os.Args[1:] {
		wg.Add(1)
		path := path
		go func() {
			defer wg.Done()
			myIPParser(path)
		}()
	}
	wg.Wait()

}

func myIPParser(fileName string) {

	fileContents, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Print(err)
	}
	var collectedIps []net.IP
	listOfWords := strings.Split(string(fileContents), " ")
	for _, word := range listOfWords {
		if ip := net.ParseIP(word); ip != nil {
			collectedIps = append(collectedIps, ip)
		}
	}
	fmt.Printf("%v: %v\n", fileName, collectedIps)

}
