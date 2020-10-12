package engine

//为engine包定义对象

//定义一个名为request的结构体，其中包含一个类型为string的url和一个类型为函数的parsefunc
type Request struct {
	Url        string
	ParserFunc func([]byte) ParserResult // 此函数输入一个字节切片，输出一个名为parserresult的结构体
}

//定义一个名为parserresult的结构体，其中包含一个由request结构体组成的切片变量requests，和一个任意类型切片的变量items
type ParserResult struct {
	Requests []Request
	Items    []interface{}
}

//定义一个什么都不做的函数，以用作临时调用的函数
func NilParser([]byte) ParserResult {
	return ParserResult{}
}
