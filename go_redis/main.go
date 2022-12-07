package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"

)
var(
	ctx=context.Background()
	r *redis.Client
)
type user struct{
	name string
	age int
}
func main(){
	rdb:=redis.NewClient(&redis.Options{Addr:"localhost:6379",Password:"",DB:0})
	//set 普通set
	err:=rdb.Set(ctx,"name","Go",0).Err()
	if err!=nil{
		fmt.Println("err=",err)
	}
	val,err:=rdb.Get(ctx,"name").Result()
	if err!=nil{
		fmt.Println("err=",err)
	}
	fmt.Println(val)
	//setnx 不存在才赋值
	err=rdb.SetNX(ctx,"city","mianyang",0).Err()
	if err!=nil{
		fmt.Println("err=",err)
	}
	val1:=rdb.Get(ctx,"city")
	fmt.Println(val1.Val())
	//批量赋值
	sli:=[]string{"zhangsan","lisi"}
	err=rdb.MSet(ctx,"val1",123,"val2",sli[0]).Err()
	if err!=nil{
		fmt.Println("err=",err)
	}
	val2:=rdb.Get(ctx,"val2")
	fmt.Println(val2.Val())
	//原子性批量赋值，要么同时成功，要么同时失败
	err=rdb.MSetNX(ctx,"val2","new233","val3","第三个值").Err()
	if err!=nil{
		fmt.Println("err=",err)
	}
	fmt.Println(rdb.Get(ctx,"val2").Val(),rdb.Get(ctx,"val3").Val())
	//列表从左插入
	i:=rdb.LLen(ctx,"list").Val()//获取列表长度
	if i<5{
		//列表从左插入
		err=rdb.LPush(ctx,"list","21",123,23,2323232).Err()
		if err!=nil{
			fmt.Println("err=",err)
		}
	}
	fmt.Println("列表数据：",rdb.LRange(ctx,"list",0,-1).Val())
	//下标获取元素

}