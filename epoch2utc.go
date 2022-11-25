package main

import (
    "encoding/json"
    "time"
    "bufio"
    "fmt"
    "math"
    "os"
    "log"
)

func main() {
    args := os.Args
    filePath := args[1]

    // Attempt to open file
    file, err := os.Open(filePath)
    if err != nil {
        log.Fatal(err)
    }

    // A map to store JSON data temporarily
    var data map[string]interface{}

    // Read file line by line
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        // Unmarshal JSON data into the map
        err := json.Unmarshal([]byte(scanner.Text()), &data)
        if err != nil {
            log.Println(err)
            continue
        }

        // Convert Epoch time to a human readable UTC format
        timestamp := data["timestamp"].(float64)
        s, ms := math.Modf(timestamp)
        time := time.Unix(int64(s), 0).UTC()
        utcTime := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d.%d", time.Year(), time.Month(), time.Day(), time.Hour(), time.Minute(), time.Second(), int64(math.Round(ms*1e6)))

        // Update timestamp, convert back to JSON and print line to stdout
        data["timestamp"] = utcTime
        newJson, err := json.Marshal(data)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Println(string(newJson))
	  }
}
