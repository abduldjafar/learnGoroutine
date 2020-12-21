package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strings"
	"time"
)

var global = map[string]int{}

func GetTokens(datas []string) []string {

	data := strings.Join(datas, " ")
	var re = regexp.MustCompile(`(?m)\w+`)

	return re.FindAllString(data, -1)

}

func WordCountEveryArray(dataArray []string) map[string]int {

	dataDicts := make(map[string]int)

	for _, key := range dataArray {
		if val, ok := dataDicts[key]; ok {
			dataDicts[key] = val + 1
		} else {
			dataDicts[key] = 1
		}
	}

	return dataDicts

}

func SummarizeCount(dataDicts map[string]int) {

	for key, value := range dataDicts {
		if val, ok := global[key]; ok {
			//do something here
			global[key] = val + value
		} else {
			global[key] = value
		}
	}
}

func main() {
	log.Println("Readfile")
	var datas []string
	dicts := map[string]int{}

	log.Println(dicts)

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
		if len(datas) == 5{

			stage1 := GetTokens(datas)

			stage2 := WordCountEveryArray(stage1)

			SummarizeCount(stage2)

			datas = []string{}

		} else {

		}

	}

	log.Println(global)
	elapsed := time.Since(start)
	log.Printf("processing took %s", elapsed)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
