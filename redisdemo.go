package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main(){
	/*  */ 
	conn, err := redis.Dial("tcp","127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis.Dial err = ", err)
		return
	}
	defer conn.Close() // 关闭

	_, err = conn.Do("set", "name", "tomjerry")
	if err != nil {
		fmt.Println("set err", err)
		return
	}

	r, err := redis.String(conn.Do("Get", "name"))
	if err != nil{
		fmt.Println("set err = ", err)
		return
	}
	// name对应值字符串
	fmt.Println("conn succ ...", r)
}
