package backtracking

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

/*
@Author: by LH
@date:  2020/8/18
@function:
*/

func TestReadBinaryWatch(t *testing.T) {
	//x:=[]int{1}
	//var c []int
	//fmt.Println(append(x, c...))
	//fmt.Printf("%b   %b\n",256,)
	//fmt.Println(unsafe.Sizeof(255))
	x := readBinaryWatch(2)
	sort.Strings(x)
	fmt.Println(x)

	//my:=[]string{"0:03","0:05","0:09","0:17","0:33","1:01","2:01","4:01","8:01","0:03","0:06","0:10","0:18","0:34","1:02","2:02","4:02","8:02","0:06","0:05","0:12","0:20","0:36","1:04","2:04","4:04","8:04","0:10","0:12","0:09","0:24","0:40","1:08","2:08","4:08","8:08","0:18","0:20","0:24","0:17","0:48","1:16","2:16","4:16","8:16","0:34","0:36","0:40","0:48","0:33","1:32","2:32","4:32","8:32","1:02","1:04","1:08","1:16","1:32","1:01","3:00","5:00","9:00","2:02","2:04","2:08","2:16","2:32","3:00","2:01","6:00","10:00","4:02","4:04","4:08","4:16","4:32","5:00","6:00","4:01","8:02","8:04","8:08","8:16","8:32","9:00","10:00","8:01"}
	//sort.Strings(my)
	current := []string{"0:03", "0:05", "0:06", "0:09", "0:10", "0:12", "0:17", "0:18", "0:20", "0:24", "0:33", "0:34", "0:36", "0:40", "0:48", "1:01", "1:02", "1:04", "1:08", "1:16", "1:32", "2:01", "2:02", "2:04", "2:08", "2:16", "2:32", "3:00", "4:01", "4:02", "4:04", "4:08", "4:16", "4:32", "5:00", "6:00", "8:01", "8:02", "8:04", "8:08", "8:16", "8:32", "9:00", "10:00"}
	sort.Strings(current)
	fmt.Println(current)
	fmt.Println(reflect.DeepEqual(x, current))
	//fmt.Println(my)
	//fmt.Println(current)


}
