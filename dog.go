package dog

import "fmt"

type Dog struct {
	Name string
	Age  int16
}

func (d Dog) String() string {
	return fmt.Sprintf("Hey %s, you are %d years old", d.Name, d.Age)
}
