package parser

import (
	"fmt"
	"learnGo/crawier/engine"
	"learnGo/crawier/model"
	"regexp"
	"strings"
)

var profileRe = regexp.MustCompile(`<div class="des f-cl" data-v-3c42fade>(.*?)</div>[ \f\n\r\t\v]+<div class="actions"`)

//var nameRe = regexp.MustCompile(`<h1 class="nickName" data-v-5b109fc3>(.*?)</h1>`)

func ParseProfile(contents []byte, name string) engine.ParseResult {

	matchs := profileRe.FindAllSubmatch(contents, -1)
	profile := model.Profile{}

	var profileMsg []string

	for _, m := range matchs {
		//fmt.Printf("%s", m[1])
		profileMsg = strings.Split(string(m[1]), "|")
	}

	//nameMatch := nameRe.FindAllSubmatch(contents, -1)

	// string(nameMatch[0][1])
	profile.Name = name

	for i, v := range profileMsg {
		fmt.Println(i, v)
		switch i {
		case 0:
			profile.Hokou = v
		case 1:
			profile.Age = v
		case 2:
			profile.Education = v
		case 3:
			profile.Marriage = v
		case 4:
			profile.Height = v
		case 5:
			profile.Income = v
		}
	}

	result := engine.ParseResult{
		Items: []interface{}{profile},
	}

	return result
}
