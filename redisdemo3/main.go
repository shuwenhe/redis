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

	_, err = conn.Do("HMSet", "user02", "name", "john", "age", 19)
	if err != nil {
		fmt.Println("HMset err", err)
		return
	}

	r, err := redis.Strings(conn.Do("HMGet","user02", "name", "age"))
	if err != nil{
		fmt.Println("hget err = ", err)
		return
	}

	for i, v := range r{
		fmt.Printf("r[%d]=%s\n", i, v)
	}
}
