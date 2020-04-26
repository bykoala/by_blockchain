package main

import "fmt"

func main(){
	fmt.Println("start...")
	bc := NewBlockchain() //创建区块链
	bc.AddBlock("bysu pay bsu 10")
	bc.AddBlock("bysu1 pay bsu 20")
	bc.AddBlock("bysu3 pay bsu 30")

	for _,block := range bc.blocks{
		fmt.Printf("上一块哈希%x\n",block.PrevBlockHash)
		fmt.Printf("数据%s\n",block.Data)
		fmt.Printf("当前哈希%x\n",block.Hash)
	}
}
