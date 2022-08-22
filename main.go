package main

import (
	"fmt"
	"os"
	"strconv"
	"encoding/json"
	"flag"
)

type NftData struct {
	Name string
	Description string
	Image string
}

func fileprint(filename string, data *NftData){
	fmt.Println(filename)
	stream, err := json.Marshal(data)
	if err != nil {
        panic(err)
    }
	os.WriteFile(filename, stream, 0755)
	fmt.Println(string(stream))
}

func main() {
	fmt.Println("Hello, world!")
	var count = flag.Int("count", 5, "the number of NFT JSON file to generate")
	var name = flag.String("name", "nft number: ", "the name of the NFT")
	var description = flag.String("desc", "", "a description of the NFT")
	var image = flag.String("image", "https://image.lol/", "The cover image of the NFT")
	var outDir = "./nftmeta/"
	_, e := os.Stat(outDir)
	if !os.IsNotExist(e) {
		fmt.Printf("File or directory %s already exist \n", outDir)
		return 
	}
	os.Mkdir(outDir, 0755)
	for i := 0; i < *count; i++{
		data := &NftData{
			*name + strconv.Itoa(i),
			*description,
			*image + strconv.Itoa(i),
		}
		fileprint(outDir + "/" + strconv.Itoa(i) + ".json", data)
	}
}