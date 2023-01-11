package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

type Size int

const (
	Unrecognized Size = iota
	Small
	Large
	May
)

func (s *Size) UnmarshalText(text []byte) error {
	switch strings.ToLower(string(text)) {
	default:
		*s = Unrecognized
	case "small":
		*s = Small
	case "large":
		*s = Large
	}
	return nil
}

func (s Size) MarshalText() ([]byte, error) {
	var name string
	switch s {
	default:
		name = "unrecognized"
	case Small:
		name = "small"
	case Large:
		name = "large"
	}
	return []byte(name), nil
}

func main() {
	blob := `["small","regular","large","unrecognized","small","normal","small","large"]`
	var inventory []Size
	if err := json.Unmarshal([]byte(blob), &inventory); err != nil {
		log.Fatal(err)
	}

	if bytes, err := json.Marshal(inventory); err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("%s \n", string(bytes))
	}

	counts := make(map[Size]int)
	for _, size := range inventory {
		counts[size] += 1
	}

	fmt.Printf("Inventory Counts:\n* Small:        %d\n* Large:        %d\n* Unrecognized: %d\n",
		counts[Small], counts[Large], counts[Unrecognized])

	// fs, e := os.Create("./test.txt")
	// if e != nil {
	// 	fmt.Println(e.Error())
	// }
	// fs.Write([]byte("hhhhh"))
	// fs.Close()
	// fmt.Println(fs.Name())

	os.WriteFile("./test.txt", []byte("aaa hello world"), os.ModeAppend)
	f, _ := os.Stat("./test.txt")
	f.Size()
	fmt.Println(f.Size())
}
