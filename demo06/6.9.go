package main

import "fmt"

func main() {
	doDBOperations()
}

func connectToDB() {
	fmt.Println("链接数据库...")
}

func disconnectFronDB() {
	fmt.Println("完成操作，断开连接...")
}

func doDBOperations() {
	connectToDB()
	fmt.Println("连接中...")
	defer disconnectFronDB()
	fmt.Println("操作数据库...")
	fmt.Println("操作失败...")
	fmt.Println("重试...")

}
