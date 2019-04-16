package parser

import (
	"crawler/engine"
	"crawler/model"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const profileRePurple = `<div class="m-btn purple" data-v-ff544c08>([^<]*)</div>`

const incomeRe = `<div class="m-btn purple" data-v-ff544c08>月收入:([^<]+)</div>`
const workplaceRe = `<div class="m-btn purple" data-v-ff544c08>工作地:([^<]+)</div>`
const heightRe = `<div class="m-btn purple" data-v-ff544c08>([^<]*)cm</div>`
const weightRe = `<div class="m-btn purple" data-v-ff544c08>([^<]*)kg</div>`
const constellationRe = `<div class="m-btn purple" data-v-ff544c08>([^<]+座)[^<]*</div>`
const ageRe = `<div class="m-btn purple" data-v-ff544c08>([0-9]+)岁</div>`

const profileRePink = `<div class="m-btn pink" data-v-ff544c08>([^<]*)</div>`

const nationRe = `<div class="m-btn pink" data-v-ff544c08>([^<]+)族</div>`
const nativePlaceRe = `<div class="m-btn pink" data-v-ff544c08>籍贯:([^<]+)</div>`

// [^<]+ / [^<]* means matching will not stop only when you meet a '<', before '^' everything is legal
// [^<]+座 means
// 1. matching will not end only when you meet a '^'
// 2. the string will be illegal besides it contains a '座'

func extractString(contexts []byte, re *regexp.Regexp) string {
	matches := re.FindSubmatch(contexts)

	if len(matches) >= 2 {
		return string(matches[1])
	} else {
		return ""
	}
}

// another way to find useful message
func getCarStatus(contents [][][]byte) string { // if find keyword '车', return this field
	for _, m := range contents {
		for i := 1; i < len(m); i++ {
			carStatus := string(m[i])
			if strings.Contains(carStatus, "车") {
				return carStatus
			}
		}
	}
	return "car condition not found "
}

func getHouseStatus(contents [][][]byte) []string { // if find keyword '住'/'房', return this field
	var houseStatus []string
	for _, m := range contents {
		for i := 1; i < len(m); i++ {
			str := string(m[i])
			if strings.ContainsAny(str, "住|房") {
				houseStatus = append(houseStatus, str)
			}
		}
	}

	if houseStatus != nil {
		return houseStatus
	}else {
		return []string{"house condition not found "}
	}
}

func Profile(contents []byte, name string, gender string) engine.ParserRequests {

	profile := model.Profile{}
	profile.Name = name
	profile.Gender = gender

	rePurple := regexp.MustCompile(profileRePurple)
	matchPurple := rePurple.FindAllSubmatch(contents, -1)

	userNotExist := engine.ParserRequests{
		Items: []interface{}{profile},
	}

	rePink := regexp.MustCompile(profileRePink)
	matchPink := rePink.FindAllSubmatch(contents, -1)

	profile.Car = getCarStatus(matchPink)

	profile.House = getHouseStatus(matchPink)

	if len(matchPurple) == 0 {
		return userNotExist
	}

	ageRe := regexp.MustCompile(ageRe)
	age, err := strconv.Atoi(extractString(contents, ageRe))
	if err != nil {
		fmt.Printf("%s's age miss : %v\n", name, err)
	}
	profile.Age = age

	heightRe := regexp.MustCompile(heightRe)
	height, err := strconv.Atoi(extractString(contents, heightRe))
	if err != nil {
		fmt.Printf("%s's height miss : %v\n", name, err)
	}
	profile.Height = height

	weightRe := regexp.MustCompile(weightRe)
	weight, err := strconv.Atoi(extractString(contents, weightRe))
	if err != nil {
		fmt.Printf("%s's weight miss : %v\n", name, err)
	}
	profile.Weight = weight

	nativePlaceRe := regexp.MustCompile(nativePlaceRe)
	nativePlace := extractString(contents, nativePlaceRe)
	profile.NativePlace = nativePlace

	nationRe := regexp.MustCompile(nationRe)
	nation := extractString(contents, nationRe)
	profile.Nation = nation

	incomeRe := regexp.MustCompile(incomeRe)
	income := extractString(contents, incomeRe)
	profile.Income = income

	workplaceRe := regexp.MustCompile(workplaceRe)
	workplace := extractString(contents, workplaceRe)
	profile.WorkPlace = workplace

	constellationRe := regexp.MustCompile(constellationRe)
	constellation := extractString(contents, constellationRe)
	profile.Constellation = constellation

	// i don't wanna lose those regexp, so keep it even though those code is redundant

	/**
	here is a print for test
	 */
	//for _, m := range matches {
	//	for _, n := range m {
	//		fmt.Printf("%s\n",n)
	//	}
	//}

	marriage := string(matchPurple[0][1])
	profile.Marriage = marriage

	occupation := string(matchPurple[len(matchPurple)-2][1])
	profile.Occupation = occupation

	education := string(matchPurple[len(matchPurple)-1][1])
	profile.Education = education

	result := engine.ParserRequests{
		Items: []interface{}{profile},
	}

	return result
}
