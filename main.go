package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

type Student struct {
	Name string
	Family string
}

func main(){
	pool := NewRedisPool()
	conn := pool.Get()
	defer conn.Close()

	if err := Ping(conn); err != nil{
		panic(err)
	}

	//if err := Set(conn, "customer:name" , "morteza amzajerdi"); err != nil{
	//	panic(err)
	//}

	key := "customer:name"
	res, err := Get(conn, key)
	if err == redis.ErrNil{
		fmt.Println("the acquire key does not exist.")
		return
	}else if err != nil{
		panic(err)
	}

	fmt.Println(key + " = " + res)

	student := Student{
		Name: "morteza",
		Family: "amzajerdi",
	}

	if err := SetStruct(conn, "student:" + student.Name ,student); err != nil{
		panic(err)
	}

	studentGet, err := GetStruct(conn, "student:" + student.Name, Student{})
	if err == redis.ErrNil{
		fmt.Println("the acquire key does not exist")
		return
	}else if err != nil{
		panic(err)
	}

	fmt.Printf("name:%s , family:%s \n", studentGet.(Student).Name, studentGet.(Student).Family)
}