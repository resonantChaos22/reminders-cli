package main

import (
	"flag"
	"fmt"
	"strings"
	"time"
)

// This interface is given in the custome value interface required for flag.Var
type idsFlag []string

// This function runs after the parsing is complete
func (ids idsFlag) String() string {
	return strings.Join(ids, ",")
}

// This function runs when parsing the flags
func (ids *idsFlag) Set(id string) error {
	*ids = append(*ids, id)
	return nil
}

type person struct {
	name string
	born time.Time
}

func (p person) String() string {
	return fmt.Sprintf("My name is: %s, and currently the time is %s", p.name, p.born.String())
}

func (p *person) Set(name string) error {
	p.name = name
	p.born = time.Now()
	return nil
}

func main() {
	var ids idsFlag
	var p person

	flag.Var(&ids, "id", "the id to be appended to the array")
	flag.Var(&p, "name", "name of the person")
	flag.Parse()
	fmt.Println(ids)
	fmt.Println(p)
}
