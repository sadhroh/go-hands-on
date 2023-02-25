// Blueshift
package main

// You may add any imports here, if you wish, but only from the 
// standard library

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

type Score struct {
     country string
     score int
}

func ReadScores() ([]Score, error) { 
    // Read `country: score` pairs from infile and return a list of pairs
    // Stop when `infile` reaches EOF
    var result []Score

    reader := bufio.NewReader(os.Stdin)
    for {
        line, err := reader.ReadString('\n')
        if err == io.EOF {
            return result, nil
        }
        
        if err != nil {
            return result, err
        }
        
        fields := strings.Split(line, ":")
        score, err := strconv.Atoi(strings.TrimSpace(fields[1]))
        if err != nil {
            return result, err
        }
        
        item := Score {
            country: strings.TrimSpace(fields[0]),
            score: score,
        }
        result = append(result, item)
    }
    return result, nil
}

func NoCenturyCountryCount(scores []Score) (int) {
    noCenturyCount := 0
    for i:=0; i<len(scores); i++ {
        country := scores[i].country
        if len(country) == 0 {
            // This country was already processed earlier
            continue
        } 

        centuries := 0
        for j:=i+1; j<len(scores); j++ {
            if scores[j].country == country {
                if scores[j].score >= 100 {
                    centuries++
                }

                // disable checking of this same country again
                scores[j].country = ""
            }
        }
        if centuries == 0 {
            noCenturyCount++
        }
    }
    return noCenturyCount;
}

func main() {
    scores, err := ReadScores()
    if err != nil {panic(err)}
    result := NoCenturyCountryCount(scores)
    fmt.Println(result)
}
