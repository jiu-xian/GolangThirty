package parser

import (
	"io/ioutil"
	"testing"
)

//为citylist解析器编写测试
func TestParserCityList(t *testing.T) {
	contents, err := ioutil.ReadFile("citylist_test_data.html") //读取事先保存的测试用文件
	if err != nil {
		panic(err)
	}
	result := ParserCityList(contents) //调用城市列表解析函数对测试文档进行处理，返回解析结果

	//对解析结果进行比对
	const resultSize = 470
	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d "+"requests; but have %d", resultSize, len(result.Requests))
	}
	if len(result.Items) != resultSize {
		t.Errorf("result should have %d "+"items; but have %d", resultSize, len(result.Items))
	}
	expectedUrls := []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}
	expectedCities := []string{
		"阿坝", "阿克苏", "阿拉善盟",
	}
	for i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("expected url is: #%s "+"but is: %s ", expectedUrls, result.Requests[i].Url)
		}
	}
	for i, city := range expectedCities {
		if result.Items[i] != city {
			t.Errorf("expected city is: #%s "+"but is: %s ", expectedCities, result.Items[i])
		}
	}
}
