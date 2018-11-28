package main

import "github.com/llSourcell/kerala"



func main()  {
	// Запустить узел
	node, err := kerala.StartNode()
	if err != nil{
		panic(err)
	}

	// Добавить текст в IPFS (создать MerkleDAG)
	var userInput = r.Form["sometext"]
	Key, err := kerala.AddString(node, userInput[0])

	// Получить текст из IPFS (извлечь MerkleDAG)
	tweetArray, _ := kerala.GetStrings(node)
}
