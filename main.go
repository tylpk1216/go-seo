package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/headzoo/surf/browser"
	"gopkg.in/headzoo/surf.v1"
)

type Rule struct {
	Search  string `json:search`
	Pattern string `json:pattern`
}

type Parameters struct {
	Agent           string `json:agent`
	SleepSecs       int    `json:clickSleepSecs`
	NextPagePattern string `json:nextPagePattern`
}

type List struct {
	Arg   Parameters `json:arg`
	Items []Rule     `json:items`
}

func clickPage(bow *browser.Browser, arg Parameters, r Rule, rex *regexp.Regexp, index int) (string, error) {
	err := bow.Open(r.Search)
	if err != nil {
		return "", err
	}

	var nextPages []string

	links := bow.Links()
	url := ""

	for _, v := range links {
		s := v.Asset.Url().String()

		if strings.Contains(s, r.Pattern) && url == "" {
			url = s
		}

		if rex.MatchString(s) {
			nextPages = append(nextPages, s)
		}
	}

	if url == "" {
		index++
		if index <= len(nextPages) {
			// click next page
			r.Search = nextPages[index-1]

			// sleep a while
			time.Sleep(time.Duration(arg.SleepSecs*1000) * time.Millisecond)

			return clickPage(bow, arg, r, rex, index)
		}

		msg := fmt.Sprintf("not found, tried %d pages(%d)", index-1, len(nextPages))
		return msg, nil
	}

	bow.SetUserAgent(arg.Agent)
	time.Sleep(time.Duration(arg.SleepSecs*1000) * time.Millisecond)

	return "clicked", bow.Open(url)
}

func main() {
	// get config setting
	exe, err := os.Executable()
	if err != nil {
		panic(err)
	}

	exePath := filepath.Dir(exe)

	jsonFile := fmt.Sprintf("%v%v%v", exePath, string(os.PathSeparator), "\\config.json")
	jsonBlob, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		panic(fmt.Sprintf("Can't find %s", jsonFile))
	}

	var list List

	if err = json.Unmarshal(jsonBlob, &list); err != nil {
		panic(err)
	}

	fmt.Printf("page pattern : %s \n\n", list.Arg.NextPagePattern)

	// prepare shared variables
	bow := surf.NewBrowser()
	rex, _ := regexp.Compile(list.Arg.NextPagePattern)

	// click page from items
	for _, v := range list.Items {
		res, err := clickPage(bow, list.Arg, v, rex, 1)

		fmt.Printf("search  : %v \n", v.Search)
		fmt.Printf("pattern : %v \n", v.Pattern)
		fmt.Printf("result  : %v \n", res)
		fmt.Printf("error   : %v \n\n", err)
	}
}
