package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`        // "released" as key
	Color  bool `json:"color,omitempty"` // "color" as key, omit if empty
	Actors []string
	// the left-trailing comma is so as not to make the compiler think we want "omitempty" as key
	TimesWatched int    `json:",omitempty"` // struct Field as key, omit if empty
	DevComment   string `json:"-"`          // omit
	// they right-trailing comma is so as not to make the compiler thing we want to omit
	HyphenKey bool `json:"-,"` // you want "-" as key
}

// empty detection, works based on JavaScript's falsy values

var movies = []Movie{{Title: "Casablanca", Year: 1942, TimesWatched: 0, Actors: []string{"Humphrey Bogart", "Ingird Bergman"}, DevComment: "testing"}}

// var moviesJSONDecoded []Movie
var moviesJSONDecoded = []Movie{}

func JSON() {
	// data, _ := json.Marshal(movies)
	dataIndent, _ := json.MarshalIndent(movies, "", "  ")
	// fmt.Printf("%s\n", dataIndent)

	// fmt.Printf("%v\n", moviesJSONDecoded) // []

	json.Unmarshal(dataIndent, &moviesJSONDecoded)
	// fmt.Printf("%v\n", moviesJSONDecoded)

	envJson, _ := json.MarshalIndent(os.Environ(), "", "  ")
	fmt.Printf("%s\n", envJson)
}

/* ----------- */

const IssuesURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

func SearchGithubIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))

	resp, err := http.Get(IssuesURL + "?q=" + q)

	if err != nil {
		return nil, err
	}

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
