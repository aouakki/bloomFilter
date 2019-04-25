package main

import (
	"fmt"
	"github.com/ouakki/bloomFilter/bf"
)

func main(){
		bf := bf.NewBloomFilter(10000)
		str := "hello world"
		bf.Insert([]byte(str))
		fmt.Printf("%t",bf.Contains([]byte(str)))
}