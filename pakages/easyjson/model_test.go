package easyjson

import (
	"fmt"
	"testing"
)

/*
@Author: by LH
@Date:  2020/11/19 10:30 上午
@Function:
*/

func TestFoo(t *testing.T) {
	s := `{"uuid":"uuid","state":"state1"}`
	var q Foo

	_ = q.UnmarshalJSON([]byte(s))
	fmt.Printf("%s", q)
}
