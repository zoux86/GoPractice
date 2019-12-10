package main

import (
	"fmt"
	//lru "LRU/lru"
	"Practice/LRU/lru"
)



func main(){
	l,err := lru.NewLruCache(5)
	if(err!=nil){
		fmt.Println("初始化LruCache失败")
	}
	l.Add("a","10")
	l.Add("b","5")
	fmt.Println(l.Get("b"))
	l.Remove("a")
	fmt.Println("a")

}
