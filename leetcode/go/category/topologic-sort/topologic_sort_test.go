package topologic_sort

import (
	"bytes"
	"fmt"
	"testing"
	"time"
)

/*
@Author: by LH
@date:  2020/8/31
@function:
*/

const (
	DateTimeFormat = "2006-01-02 15:04:05"
	DateFormat     = "2006-01-02"
)

type XTime struct {
	time.Time
}

func Now() XTime {
	return XTime{
		time.Now(),
	}
}

func (t *XTime) UnmarshalJSON(b []byte) error {
	b = bytes.Trim(b, "\"")
	ext, err := time.Parse(DateTimeFormat, string(b))
	if err != nil {
		return err
	}
	*t = XTime{ext}
	return nil
}
func (t *XTime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", t.Time.Format(DateTimeFormat))
	return []byte(stamp), nil
}
func TestCanFinish(t *testing.T) {
	//x:= Now()
	//b,e:=json.Marshal(&x)
	//if e != nil {
	//	log.Fatal(e)
	//}
	//fmt.Println(string(b))

	//x := time.Now()
	//b, _ := json.Marshal(&x)
	//fmt.Println(string(b))
	//
	//var y time.Time
	//e := json.Unmarshal(b, &y)
	//if e != nil {
	//	log.Fatal(e)
	//}
	//fmt.Println(y.MarshalJSON())
	//z := x
	//fmt.Println(unsafe.Pointer(&x), unsafe.Pointer(&z))
	var x TT
	fmt.Println(x,x.AAA)
}

type TT struct {
	A int `json:"a"`
	AAA AA `json:"aaa"`
}
type AA struct {
	B int `json:"b"`
}
