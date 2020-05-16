package main

import (
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func NewRedisPool() *redis.Pool{
	return &redis.Pool{
		MaxIdle: 80,
		MaxActive: 12000,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", ":6379")
			if err != nil{
				panic(err.Error())
			}
			return c, err
		},
	}
}

func Ping(c redis.Conn) error{
	s, err := redis.String(c.Do("PING"))
	if err != nil{
		return err
	}
	
	fmt.Printf("PING Response = %s\n", s)

	return nil
}

func Set(c redis.Conn, key, value string) error{
	_, err := c.Do("SET" , key, value)
	if err != nil{
		return err
	}

	return nil
}

func Get(c redis.Conn, key string) (string, error){
	return redis.String(c.Do("GET" , key))
}

func SetStruct(c redis.Conn, key string, val interface{}) error{
	js, err := json.Marshal(val)
	if err != nil{
		return err
	}

	if err := Set(c, key, string(js)); err != nil{
		return err
	}

	return nil
}

func GetStruct(c redis.Conn, key string, t interface{}) (interface{}, error){
	res, err := Get(c, key)
	if err != nil{
		return nil, err
	}

	switch t.(type) {
	case Student:
		var resInterface Student
		err = json.Unmarshal([]byte(res), &resInterface)
		return resInterface, err
	}

	return nil, nil
}