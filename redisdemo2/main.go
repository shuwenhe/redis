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

	_, err = conn.Do("HSet", "user01", "name", "john")
	if err != nil {
		fmt.Println("hset err", err)
		return
	}

	_, err = conn.Do("HSet", "user01", "age", 18)
	if err != nil {
		fmt.Println("hset err", err)
		return
	}

	r1, err := redis.String(conn.Do("HGet","user01", "name"))
	if err != nil{
		fmt.Println("hget err = ", err)
		return
	}

	r2, err := redis.Int(conn.Do("HGet","user01", "age"))
	if err != nil{
		fmt.Println("hget err = ", err)
		return
	}

	// name对应值字符串
	fmt.Printf("conn succ\nr1 = %v\nr2 = %v\n", r1, r2)
}
