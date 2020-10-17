package parser

import (
	"io/ioutil"
	"learngo/GolangThirty/crawler/model"
	"testing"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data.html")
	if err != nil {
		panic(err)
	}
	result := ParseProfile(contents, "无戏配角小短腿i")
	if len(result.Items) != 1 {
		t.Errorf("Items should contain 1 "+"element; but was %v", result.Items)
	}

	profile := result.Items[0].(model.Profile)
	//profile.Name = "无戏配角小短腿i"

	expected := model.Profile{
		Name:       "无戏配角小短腿i",
		Gender:     "男",
		Age:        78,
		Height:     30,
		Weight:     238,
		Xinzuo:     "金牛座",
		Income:     "1-2000元",
		Marriage:   "离异",
		Occupation: "销售",
		Education:  "初中",
		Hokou:      "西安市",
		House:      "无房",
		Car:        "有车",
	}

	if profile != expected {
		t.Errorf("expected %v; but was %v", expected, profile)
	}
}
