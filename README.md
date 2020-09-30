# 去浪三十日（GoLangThirty）


## 小白零基础自学Golang入门尝试，欢迎各位大神指导


## PreDay：开始之前的准备：

- 本地安装go1.15，配置环境和依赖（编译器：vscode ；module on模式）
    > - 官网下载go1.15安装包进行安装,安装完成后检查是否安装成功
    > - ` $ go version `
    > - 检查go工作环境
    > - ` $ go env `
    > - 更改国内代理GOPROXY和GO111MODULE两项参数
    > - ` $ go env -w GOPROXY=https://goproxy.cn,direct `
    > - ` $ go env -w GO111MODULE=on `
    > - 下载一个goimports的工具包
    > - ` $ go get -v golang.org/x/tools/cmd/goimports `

- 在GitHub上创建项目库，并和本地项目库同步，GitHub使用教程参考：[廖雪峰的git教程](https://www.liaoxuefeng.com/wiki/896043488029600)
- 编写Hello World ！
- README书写，markdown的语法参考:[markdown菜鸟教程](https://www.runoob.com/markdown/md-tutorial.html)

- **心得**
    > - 万事开头难，准备工作太繁琐，对于新手来说，能顺利输出“helloworld”就是一件挺不容易的事情
    > - 如果有条件，尽量还是不要自学，**大神指路、如有神助**，很多自己花一两个小事搞不明白的事情，大神也就一两句话就让你峰回路转了
    > - 个人比较追求完美，所以在README的编写上，也希望能够做的越来越细致 *（虽然这会浪费很多时间）* 
    > - 通过检索相关资料，给自己定了一个30天的学习计划，希望可以坚持下来


## Day-1：基础语法学习：(*代码示例和注释见./Day-1/basic.go*)
1. 声明变量的几个要点
    > - 变量类型写在变量名之后
    > - 默认值声明/声明+赋值/初始值声明/简短声明/
    > - 块内声明/函数内声明/包内声明
    > - 编译器可以推测变量类型
    > - 没有char，只有rune
    > - 原生支持复数类型

2. Go内建变量的类型
    > - bool, string
    > - (u)int, (u)int8, (u)int16, (u)int32, (u)int64, uintptr
    > -  byte, rune 
    > - float32, float64, complex64, complex128
    > - Go的类型转换是强制的

3. 常量和枚举
    > - 未声明类型的const常量可以以任何类型使用
    > - 使用const代码块进行普通枚举
    > - 使用const代码块和iota公式进行自增值枚举

4. 条件语句
    > - if……/if……else……/if……else if……else……/基本用法
    > - if语句内可以赋值变量，但此变量仅限于该if代码块内使用
    > - switch...case...语句的用法
    > - switch后可以不加条件语句
    > - switch语句中不需要break
    > - for循环的基本用法
    > - for省略初始条件、结束条件和递增表达式的用法

5. 函数
    > - 函数名在前，返回值类型在后
    > - 可定义多个返回值
    > - 函数的参数也可以是函数
    > - 没有可选参数，只有可变长参数
    > - 可直接在函数体内写匿名函数

6. 指针
    > - go不能对指针进行运算
    > - go的参数传递都是值传递
    > - 注意指针、地址、值之间的关系和差别



## Day-2: Golang的内建容器（数组array、切片slice、映射map）：(*代码示例和注释见./Day-2/container.go*)
1. 数组：array
    > - 数组的定义（var/:=/[3]int/[...]int)
    > - 数组是值类型，函数调用的是原始数组的拷贝
    > - [3]int 和 [5]int是两种不同类型的数组，不能相互赋值

2. 切片：slice
    > - slice的特性  
    > > - slice本身没有值，它是对底层数组的一个view  
    > > - slice的表示方法  
    > > - reslice是对同一个底层数组的不同view  
    > > - slice的扩展,可以向后扩展，不能向前扩展,向后扩展时不能超越底层数组
    > - slice的内部结构（ptr/len/cap的概念的用法)  
    > > - ptr(pointer):指向slice的第一个元素  
    > > - len(length):该slice的长度  
    > > - cap(capacity):该slice从第一个元素到原始数组最后一个元素的容量
    > - slice的操作  
    > > - 创建slice  
    > > - 增加slice的元素  
    > > - 删除slice的元素  
    > > - slice的拷贝  

3. 映射：map
    > - map的定义：map[k]v, map[k1]map[k2]v, make(map[string]int)
    > - map的操作:  
    > > - 获取元素：m[key]  
    > > - 获取元素时，如果key不存在，返回value类型的初始值  
    > > - 用 value， ok := m[key] 来判断是否存在对应的key  
    > > - 使用delete删除一个key，使用m[k]=v增加元素  
    > > - 使用range遍历key，或者遍历key，value对，但是不保证遍历顺序，可将key赋值给一个slice进行排序遍历  
    > > - 使用len获得元素的个数  
    > - map的key的类型  
    > > - map使用哈希表，必须可以比较相等  
    > > - 除了slice,map,function的内建类型都可以作为key  
    > > - struct类型中如果不包括上述字段，也可以作为key  
    > - 例题：最长不重复字符串  

4. 字符串处理
    > - rune相当于其他语言里的char
    > - 使用range遍历position，rune对，注意position并不是逐一递增的
    > - 使用utf8.runeCountInString()获得字符数量
    > - 使用len（）获得字节长度
    > - 使用[]byte()获得字节slice
    > - 使用[]rune()获得字符slice
    > - 常用字符串处理：  
    > > - Fields,Split,Join;  
    > > - Cointains, Index;  
    > > - ToLower, ToUpper;  
    > > - Trim, TrimRight, TrimLeft  

5. 结构体：struct
    > - golang只有封装，没有类和继承
    > - struct类型的定义：type name struct {a string, b int , c ptr, d structtypename, ...}
    > - struct变量的创建：  
    > > - var name structtypename  
    > > - name := structtypename{}  
    > > - name.a = /  name.d.a =  
    > > - 可以使用自定义的工厂函数创建变量  
    > - 为struct类型定义方法，可用于所有该类型的结构体变量  
    > > - struct的方法是值传递  
    > > - 可以使用指针类型作为接受者，这样才能改变结构体内的值  
    > > - nil的指针也可以传递  
    > > - 编译器可以自行根据需要传递指针或值  

6. 包、封装和扩展:(关于包的封装和扩展，还需要多加练习以熟练掌握)
    > - 每个目录只能有一个包，但可以有多个文件
    > - main包是程序的入口
    > - 命名采取CamelCase驼峰命名法
    > - 函数名、变量名、方法名等，首字母大写表示public，首字母小写表示private
    > - 为结构定义的方法必须放在同一个包内，但可以是不同的文件
    > - 如果原有的包中的某个结构不包括所需要的方法，可以对其进行扩展  
    > > - 使用组合的方式扩展已有类型  
    > > - 定义别名的方式扩展已有类型  
    > > - 使用内嵌的方式扩展已有类型  





