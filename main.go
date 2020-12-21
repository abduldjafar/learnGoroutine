package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strings"
	"sync"
	"time"
)

var global = map[string]int{}
var MapConchandling = sync.RWMutex{}

func GetTokens(waitGroup *sync.WaitGroup, datas []string, result chan []string) {
	defer waitGroup.Done()
	data := strings.Join(datas, " ")
	var re = regexp.MustCompile(`(?m)\w+`)

	result <- re.FindAllString(data, -1)

}

func WordCountEveryArray(waitGroup *sync.WaitGroup, data <-chan []string, dicts chan map[string]int) {
	defer waitGroup.Done()
	dataArray := <-data
	dataDicts := make(map[string]int)

	for _, key := range dataArray {
		if val, ok := dataDicts[key]; ok {
			//do something here
			dataDicts[key] = val + 1
		} else {
			dataDicts[key] = 1
		}
	}

	dicts <- dataDicts

}

func SummarizeCount(waitGroup *sync.WaitGroup, data <-chan map[string]int) {
	defer waitGroup.Done()
	dataDicts := <-data

	MapConchandling.Lock()
	for key, value := range dataDicts {
		if val, ok := global[key]; ok {
			//do something here
			global[key] = val + value
		} else {
			global[key] = value
		}
	}
	MapConchandling.Unlock()
}
func main() {
	log.Println("Readfile")
	var datas []string
	c1 := make(chan []string)
	dicts := make(chan map[string]int)

	log.Println(dicts)

	var wg sync.WaitGroup

	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	log.Println("processing file")
	start := time.Now()
	
	for scanner.Scan() {
		datas = append(datas, scanner.Text())
		if len(datas) == 5 {
			wg.Add(1)
			go GetTokens(&wg, datas, c1)

			wg.Add(1)
			go WordCountEveryArray(&wg, c1, dicts)

			wg.Add(1)
			go SummarizeCount(&wg, dicts)
			datas = []string{}

		}
	}

	wg.Wait()
	close(c1)

	log.Println(global)
	elapsed := time.Since(start)
	log.Printf("processing took %s", elapsed)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
