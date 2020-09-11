package sample

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

/*
@Author: by LH
@date:  2020/8/10
@function:
*/

func TestAB(t *testing.T) {
	type Auth struct {
		Age  int
		Name string
	}
	var a = []Auth{}

	sliceValue := reflect.Indirect(reflect.ValueOf(&a))
	//value.CanSet()
	//fmt.Println(sliceValue.Index(0))

	s := make([]interface{}, 2)
	//for i := 0; i < 2; i++ {
	//	s[i] = new(interface{})
	//}
	//s[0]=999

	//var v reflect.Value
	vp := reflect.New(sliceValue.Type().Elem())
	v := reflect.Indirect(vp)
	//fmt.Println("ssss",v.Elem().NumField())
	age := v.FieldByName("Age")
	//if age.Kind() == reflect.Ptr && age.IsNil() {
	//	fmt.Println("sssss0")
	//	age.Set(reflect.New(age.Type()))
	//}
	s[0] = age.Addr().Interface()

	name := v.FieldByName("Name")
	//if name.Kind() == reflect.Ptr && name.IsNil() {
	//	name.Set(reflect.New(name.Type()))
	//}
	s[1] = name.Addr().Interface()
	s0 := s[0].(*int)
	*s0 = 999
	s1 := s[1].(*string)
	*s1 = "HAHS"

	s[1] = "ssssssss"

	sliceValue.Set(reflect.Append(sliceValue, v))

	fmt.Println(a)
}


func TestSolve(t *testing.T){
	board := [][]byte{
		{'X','O','X','O','X','O'},
		{'O','X','O','X','O','X'},
		{'X','O','X','O','X','O'},
		{'O','X','O','X','O','X'},
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			fmt.Printf("%s\t",string(board[i][j]))
		}
		fmt.Println()
	}
	solve(board)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			fmt.Printf("%s\t",string(board[i][j]))
		}
		fmt.Println()
	}
}

func TestStructSize(t *testing.T){
	type A struct {//占24字节

		//j bool
		//i bool
		h bool
		x int16
		//g bool

		f bool
		y int16
		//e bool
		//
		//d bool

		a bool//以上八个占用一个64位系统的机器字，即8字节，去掉七个也是占用8字节，合理排序struct成员顺序可以节省空间

		b float64
		c int16
	}
	type B struct {//占16字节
		bool
		int16
		float64
	}
	type C struct {//占16字节
		float64
		bool
		int16
	}
	type D struct {//2个机器字，64位系统上为16字节，分别为数组的data头指针 和 len存储
		string
	}
	var xx *uint8
	fmt.Println(unsafe.Sizeof(A{}))
	fmt.Println(unsafe.Sizeof(B{}))
	fmt.Println(unsafe.Sizeof(C{}))
	fmt.Println(unsafe.Sizeof(D{}))
	fmt.Println(unsafe.Sizeof(xx))
}

func TestUnsafePoint(t *testing.T){
	var a *int
	fmt.Println(reflect.TypeOf(a))
}