package parser

import (
	"learngo/GolangThirty/crawler/engine"
	"regexp"
)

const CityListRe = `<a href="(http://localhost:8080/mock/www.zhenai.com/zhenghun/[a-z0-9]+)"[^>]*>([^<]+)</a>`

func ParserCityList(contents []byte) engine.ParserResult { //定义一个城市列表解析器，输入一个字节数组类型的参数，返回一个解析结果类型的结构体
	re := regexp.MustCompile(CityListRe)        //通过正则表达式进行强制查询匹配
	matches := re.FindAllSubmatch(contents, -1) //返回查找结果中所需的子关键信息

	result := engine.ParserResult{} //定一个解析结果类型的结构体变量result

	limit := 20

	for _, m := range matches { //针对查找到的每一个结果，将其赋值给变量m，此处m是一个数组
		result.Items = append(result.Items, "City: "+string(m[2])) //将m的第3个元素添加给解析结果的items变量
		result.Requests = append(result.Requests, engine.Request{  //将m的第2个元素和一个空函数作为一个新的请求任务整体添加给解析结果
			Url:        string(m[1]), //注意传输的数据类型
			ParserFunc: ParserCity,   //此处尚未定义具体的解析函数，所以用一个空函数代替nil以正常编译
		})
		limit--
		if limit == 0 {
			break
		}

	}
	return result //返回城市列表解析器的解析结果
}
