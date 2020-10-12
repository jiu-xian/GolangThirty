package parser

import (
	"io/ioutil"
	"testing"
)

func TestParserCity(t *testing.T) {
	contents, err := ioutil.ReadFile("city_test_data.html")
	if err != nil {
		panic(err)
	}

	result := ParserCity(contents)
	const resultSize = 20

	if len(result.Items) != resultSize {
		t.Errorf("expected %d "+"users but have %d", resultSize, len(result.Items))
	}
	if len(result.Requests) != resultSize {
		t.Errorf("expected %d "+"requests but have %d", resultSize, len(result.Requests))
	}

	expectedUrls := []string{
		"http://album.zhenai.com/u/1693313153",
		"http://album.zhenai.com/u/1567399184",
		"http://album.zhenai.com/u/1926955946",
	}
	expectedUsers := []string{
		"User: 醉思恋人", "User: 懂得珍惜才配拥有", "User: styi",
	}
	for i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("expected url is: #%s "+"but is: %s ", url, result.Requests[i].Url)
		}
	}
	for i, user := range expectedUsers {
		if result.Items[i] != user {
			t.Errorf("expected user is: #%s "+"but is: %s ", user, result.Items[i])
		}
	}
}
