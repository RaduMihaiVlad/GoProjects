package main

import (
    "encoding/csv"
    "fmt"
    "log"
    "os"
    "bufio"
    "strings"
    "time"
    "flag"
)

func readCsvFile(filePath string) [][]string {
    f, err := os.Open(filePath)
    if err != nil {
        log.Fatal("Unable to read input file " + filePath, err)
    }
    defer f.Close()

    csvReader := csv.NewReader(f)
    records, err := csvReader.ReadAll()
    if err != nil {
        log.Fatal("Unable to parse file as CSV for " + filePath, err)
    }

    return records
}

func getInput(C chan string) {
    reader := bufio.NewReader(os.Stdin)
    text, _ := reader.ReadString('\n')

    C <- text
}

func main() {

    /*
        This task I read from numbers.csv some math operations and their answer. The user has to 
        answer these questions in a given time(which can be modified from the limit's flag). 
        If the user runs out of time, the program stops and he receive his score.
    */

	records := readCsvFile("numbers.csv")

    score := 0
    maxScore := len(records)

    inputChan := make(chan string)

    duration := flag.Int("limit", 2, "limit of time for user")
    flag.Parse()

    timer := time.NewTimer(time.Duration(*duration) * time.Second)

    for _, record := range(records) {
        fmt.Print(record[0] + " = ")
        go getInput(inputChan)
        select {
            case <- timer.C:
                fmt.Printf("Time expired. You scored %d out of %d", score, maxScore)
                return
            
            case text := <- inputChan:
                if text[:(len(text) - 2)] == strings.Trim(record[1], " \n") {
                    score ++
                }    
        }
    }

    fmt.Printf("You scored %d out of %d", score, maxScore)
}
