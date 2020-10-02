package maps

import "fmt"

//++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
//映射map
func mapTest() {
	m1 := map[string]string{
		"name":     "jiuxian",
		"learning": "golang",
		"feeling":  "notbad",
		"purpose":  "blockchain",
	}
	var m2 map[string]int      //此处的zerovalue是nil
	m3 := make(map[string]int) //此处的zerovalue是0

	fmt.Println(m1)
	fmt.Println(m2)
	fmt.Println(m3)

	a := m1["name"]
	fmt.Println(a)

	b := m1["nime"]
	fmt.Println(b) //此处由于map中无nime字段，故打印出value类型的空值：0

	//遍历map元素
	for k, v := range m1 {
		fmt.Println(k, v)
	}

	//删除字段
	delete(m1, "name")

	//查找字段
	if v, ok := m1["name"]; ok {
		fmt.Println(v, ok)
	} else {
		fmt.Println("您要找的值不存在", "m1的元素数量为： ", len(m1))
	}
}

func NonRepeating(s string) int {
	lastoccurred := make(map[rune]int)
	start := 0
	maxlength := 0
	for i, v := range []rune(s) {
		if lasti, ok := lastoccurred[v]; ok && lasti >= start {
			start = lasti + 1
		}
		if i-start+1 > maxlength {
			maxlength = i - start + 1
		}

		lastoccurred[v] = i
	}
	return maxlength
}

func Test() {
	fmt.Println("MapTestResult:")
	mapTest()

	fmt.Println(NonRepeating("abcdacdebcd"))
	fmt.Println(NonRepeating("abcdefg"))
	fmt.Println(NonRepeating(""))
	fmt.Println(NonRepeating("我爱family！"))
	fmt.Println(NonRepeating("一二三三二一"))
	fmt.Println("----------------------------------")
}
