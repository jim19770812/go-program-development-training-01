package main

import "fmt"

func main() {
	 fmt.Println("1 入门")
	 fmt.Println("1.1 hello world")
	 /**
	 通过go build进行二进制编译
	 go的代码是通过包来组织的，go中的包类似java中的静态类，不需要创建，导入就可以直接使用，比如上面的 import "fmt"，导入的fmt就直接调用Println方法了
	  */
	 fmt.Println("通过go build进行编译成二进制程序")
	 /**
	 go的包包含一个或多个go文件
	 每个源文件都用 package xxxx 声明打头
	 package main 是个特殊的包，相当与java的 public static void main(String[] args)，其中的 func main(){}就是启动函数
	  */
	 fmt.Println("")

}
