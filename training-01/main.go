package main

import (
	"fmt"
	"os"
)

func main() {
	/**
	 */
	fmt.Println("通过go build进行编译成二进制程序")
	fmt.Println(`
1 入门
1.1 hello world
	通过go build进行二进制编译
	go的代码是通过包来组织的，go中的包类似java中的静态类，不需要创建，导入就可以直接使用，比如上面的 import "fmt"，导入的fmt就直接调用Println方法了
	go的包包含一个或多个go文件
	每个源文件都用 package xxxx 声明打头
	package main 是个特殊的包，相当与java的 public static void main(String[] args)，其中的 func main(){}就是启动函数
	包导入 import xxxx，导入方式和java一样，除了导入包，还可以导入包里的函数，变量，常量，类型
	go fmt的作用等价与gofmt，是用于强制格式化当前目录或者制定目录下的文件，go代码的格式化要求严格
	go中 main.go比较特别，相当与java中的public static void main(){}
	gofmt是一个go自带的格式化工具
	goimports是一个第三方导入声明的插入/移除工具，通过 go get golang.org/x/tools/cmd/goimports获得

	编译
	go build helloworld.go
	执行后可以编译成一个可执行文件
	go语言不需要用分号结尾，这点和python和js一致
	go语言对格式要求很严格，但一般来说，java的习惯是可以直接套用再go语言里的，但c的{}习惯会在go语言里比较难受。
	`)
	/**
	 */
	fmt.Println(`
1.2 命令行参数
	os包提供了Args的变量来提供命令行参数的访问
	import (
		"fmt"
		"os"
	)
	fmt.Println(len(os.Args)) 	//获取参数的个数
	fmt.Println(os.Args[0]) 	//获取第0个参数

	`)
	fmt.Println("os.ARgs 长度: ", len(os.Args))
	for i := 0; i <= len(os.Args)-1; i++ {
		fmt.Printf("参数%d: %s", i, os.Args[i])
	}
	fmt.Println(`
	`)
}
