package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
)

func fetchCSV(url string) ([]map[string]string, error) {
	// fetch url
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// check for 200 status code
	if resp.StatusCode != http.StatusOK {
		return nil, err
	}

	// read the csv
	reader := csv.NewReader(resp.Body)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	// convert to slice of maps using header row as keys
	headers := records[0]
	var results []map[string]string

	for _, record := range records[1:] {
		result := make(map[string]string)
		for i, value := range record {
			// ignore status field as its always released due to our search query
			if headers[i] != "Status" {
				result[headers[i]] = value
			}
		}
		results = append(results, result)
	}

	return results, nil
}

func main() {
	// fetch esu file
	esu, err := fetchCSV("https://cve.tuxcare.com/els/download-csv?os=250b68a8-d847-467a-b1a2-51c1917f1164&status=4761a92a-acb2-412a-aa23-0e86510efa78&orderBy=cve-asc")
	if err != nil {
		fmt.Println("Error fetching ESU CSV:", err)
		return
	}

	// fetch fips file
	fips, err := fetchCSV("https://cve.tuxcare.com/els/download-csv?os=e89ddb3c-c0b3-454f-8257-ee6d6505f1b8&status=4761a92a-acb2-412a-aa23-0e86510efa78&orderBy=cve-asc")
	if err != nil {
		fmt.Println("Error fetching FIPS CSV:", err)
		return
	}

	// merge both files
	merged := append(esu, fips...)

	// sort merged data by cve
	sort.Slice(merged, func(i, j int) bool {
		return merged[i]["CVE"] < merged[j]["CVE"]
	})

	// convert to json
	marshalled, err := json.Marshal(merged)
	if err != nil {
		fmt.Println("Error converting to JSON:", err)
		return
	}

	// pretty print the json instead of piping to jq
	var prettified bytes.Buffer
	err = json.Indent(&prettified, marshalled, "", "  ")
	if err != nil {
		fmt.Println("Error indenting JSON:", err)
		return
	}
	fmt.Println(prettified.String())
}
