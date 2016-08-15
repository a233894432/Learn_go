//2016/8/12 15:35
/**
关键字 defer 允许我们进行一些函数执行完成后的收尾工作，例如：

关闭文件流：

// open a file defer file.Close() （详见第 12.2 节）

解锁一个加锁的资源

mu.Lock() defer mu.Unlock() （详见第 9.3 节）

打印最终报告

printHeader() defer printFooter()

关闭数据库链接

// open a database connection defer disconnectFromDB()

合理使用 defer 语句能够使得代码更加简洁。
*/
package main

import "fmt"

func main() {
	doDBOperations()
}

func connectToDB() {
	fmt.Println("ok, connected to db")
}

func disconnectFromDB() {
	fmt.Println("ok, disconnected from db")
}

func doDBOperations() {
	connectToDB()
	fmt.Println("Defering the database disconnect.")
	defer disconnectFromDB() //function called here with defer
	fmt.Println("Doing some DB operations ...")
	fmt.Println("Oops! some crash or network error ...")
	fmt.Println("Returning from function here!")
	return //terminate the program
	// deferred function executed here just before actually returning, even if
	// there is a return or abnormal termination before
}



/* out:
	ok, connected to db
	Defering the database disconnect.
	Doing some DB operations ...
	Oops! some crash or network error ...
	Returning from function here!
	ok, disconnected from db

 */