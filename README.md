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


## 1-Day：基础语法学习：
- 声明变量的几个要点
    > - 变量类型写在变量名之后
      - 编译器可以推测变量类型
      - 没有char，只有rune
      - 原生支持复数类型

- Go内建变量的类型
    > - bool, string
      - (u)int, (u)int8, (u)int16, (u)int32, (u)int64, uintptr
      -  byte, rune 
      - float32, float64, complex64, complex128
      - Go的类型转换是强制的

- 常量和枚举
    > - 未声明类型的const常量可以以任何类型使用
      - 使用const代码块进行普通枚举
      - 使用const代码块和iota公式进行自增值枚举