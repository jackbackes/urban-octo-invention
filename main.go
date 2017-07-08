
package main

import (
	"github.com/BurntSushi/toml"
	"fmt"
	"github.com/davecgh/go-spew/spew"
)

type Entity struct {
	Type string
	Cls string
}

type GameData struct {
	Entities []Entity `toml:"entities"`
	Inventory []Inventory
}

type Inventory struct {
	Type string
	Qty int
}

func main() {
	var data GameData
	if _, err := toml.DecodeFile("data.toml", &data); err != nil {
		fmt.Println(err)
		return
	}

	spew.Dump(data.Inventory)
}

