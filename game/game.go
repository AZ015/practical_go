package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

const (
	MaxX = 1000
	MaxY = 600
)

type Key byte

const (
	Jade Key = iota + 1
	Copper
	Crystal
	invalidIdKey
)

func (k Key) String() string {
	switch k {
	case Jade:
		return "jade"
	case Copper:
		return "copper"
	case Crystal:
		return "crystal"
	default:
		return fmt.Sprintf("<Key %d>", k)
	}
}

type Item struct {
	X int
	Y int
}

func NewItem(x, y int) (*Item, error) {
	if y < 0 || y > MaxY || x < 0 || x > MaxX {
		return nil, fmt.Errorf("%d/%d out of bounds %d/%d", x, y, MaxX, MaxY)
	}

	return &Item{x, y}, nil
}

// by value without mutation
//func (i Item) Move(x, y int) {
//	i.X = x
//	i.Y = y
//}

// by pointer with mutation
func (i *Item) Move(x, y int) {
	i.X = x
	i.Y = y
}

type T struct {
	X int
}

type Player struct {
	Name string
	//X    int
	Item
	Keys []Key
	T
}

func (p *Player) FoundKey(key Key) error {
	if key < Jade || key > invalidIdKey {
		return fmt.Errorf("invalid key: %#v", key)
	}

	if !containsKey(p.Keys, key) {
		p.Keys = append(p.Keys, key)
	}

	return nil
}

func containsKey(keys []Key, key Key) bool {
	for _, k := range keys {
		if k == key {
			return true
		}
	}

	return false
}

func main() {
	var item1 Item
	fmt.Println(item1)
	fmt.Printf("item1: %#v\n", item1)

	item2 := Item{1, 2}
	fmt.Printf("item2: %#v\n", item2)

	item3 := Item{
		X: 3,
		Y: 4,
	}
	fmt.Printf("item3: %#v\n", item3)

	item4 := Item{
		X: 5,
	}
	fmt.Printf("item4: %#v\n", item4)

	// ------------------------------------------------------------

	fmt.Println("-------------------- CONSTRUCT ---------------------------")

	item5, err := NewItem(7, 8)
	fmt.Printf("item5: %#v, err: %v\n", item5, err)
	item6, err := NewItem(1001, 8)
	fmt.Printf("item6: %#v, err: %v\n", item6, err)
	item7, err := NewItem(100, 604)
	fmt.Printf("item7: %#v, err: %v\n", item7, err)
	fmt.Println(NewItem(-10, 9))

	// ------------------------------------------------------------

	fmt.Println("-------------------- MOVE ---------------------------")
	i, _ := NewItem(1, 1)
	fmt.Println(i)
	i.Move(2, 2)
	fmt.Println(i)

	// ------------------------------------------------------------

	fmt.Println("-------------------- PLAYER ---------------------------")
	p1 := Player{
		Name: "Andrew",
	}
	fmt.Printf("player1: %#v\n", p1)

	p2 := Player{
		Name: "Peter",
		Item: Item{300, 200},
	}
	fmt.Printf("player2: %#v, X(inner):%d, X(embeded):%d, Y:%d\n", p2, p2.T.X, p2.Item.X, p2.Y)

	// ------------------------------------------------------------

	fmt.Println("-------------------- IOTA ---------------------------")
	k := Jade
	fmt.Println("key: ", k)
	fmt.Println("key: ", Key(17))

	// ------------------------------------------------------------

	fmt.Println("-------------------- JSON ---------------------------")
	json.NewEncoder(os.Stdout).Encode(time.Now())
}
