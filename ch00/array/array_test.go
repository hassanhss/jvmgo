package array

import (
	"fmt"
	"testing"
)

func TestPeoples_SetPeople(t *testing.T) {
	peoples := newPeoples(10)
	peoples.SetPeople(0, People{Name: "hassan"})

	p := make([]People, 10)
	fmt.Println("添加后people数量：", len(peoples))
	for i, val := range peoples {
		fmt.Println("调用后index", i)
		fmt.Println("调用后value", val)
	}
}
