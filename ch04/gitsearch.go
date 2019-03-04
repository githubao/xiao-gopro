// github search
// author: baoqiang
// time: 2019/3/4 下午8:08
package ch04

import (
	"time"
	"net/url"
	"strings"
	"net/http"
	"fmt"
	"encoding/json"
	"github.com/labstack/gommon/log"
)

const IssueURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string    `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string // in Markdown format
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

func RunSearchIssues() {
	res, err := SearchIssues([]string{"golang", "docker"})
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(render.Render(res))

	//data,_ := json.MarshalIndent(res,"","\t")
	//fmt.Println(string(data))

	fmt.Printf("%d Issues\n", res.TotalCount)
	for _, item := range res.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}

}

func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))

	//resp, err := http.Get(IssueURL+"?q="+q)
	//if err != nil{
	//	return nil,err
	//}

	req, err := http.NewRequest("GET", IssueURL+"?q="+q, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/vnd.github.v3.text-match+json")
	resp, err := http.DefaultClient.Do(req)

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}

	resp.Body.Close()

	return &result, nil
}
