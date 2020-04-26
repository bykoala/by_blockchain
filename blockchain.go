package main

type BlockChain struct {
	blocks[] *Block   //一个数组，每个元素都是指针，存储block区块
}

//新增一个区块
func (b *BlockChain) AddBlock(data string)  {
	prevBlock := b.blocks[len(b.blocks)-1]  //取出最后一个区块
	newBlcok := NewBlcok(data,prevBlock.Hash) //创建一个区块
	b.blocks = append(b.blocks,newBlcok)//区块链插入新的区块
}

//创建一个区块链
func NewBlockchain() *BlockChain{
	return &BlockChain{[]*Block{NewGenesisBlock()}}
}
