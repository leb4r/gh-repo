package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type IssueLabel struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type Issue struct {
	Id     int64        `json:"id"`
	Title  string       `json:"title"`
	Labels []IssueLabel `json:"labels"`
}

func getOpenIssues(owner, repo string) []byte {
	resp, err := http.Get("https://api.github.com/repos/" + owner + "/" + repo + "/issues")
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return body
}

func prettyJson(data interface{}) (string, error) {
	val, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return "", err
	}
	return string(val), nil
}

func main() {
	repoPtr := flag.String("repo", "", "Public repository to query for open issues and labels.")
	flag.Parse()

	// exit if repo argument is not set
	if *repoPtr == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	splitResults := strings.Split(*repoPtr, "/")
	if len(splitResults) != 2 {
		msg := fmt.Sprintf("%s is not a valid GitHub repository identifier. Use the format of owner/repo when specifying.", *repoPtr)
		log.Fatal(msg)
	}

	owner := splitResults[0]
	repo := splitResults[1]

	// get open issues
	resp := getOpenIssues(owner, repo)

	var issues []Issue
	if err := json.Unmarshal(resp, &issues); err != nil {
		log.Fatal(err)
	}

	if prettyJson, err := prettyJson(issues); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(prettyJson)
	}
}
