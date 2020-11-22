package easyjson

/*
@Author: by LH
@Date:  2020/11/19 9:58 上午
@Function:
*/

type Foo struct {
	UUID  string `json:"uuid"`         // will not be interned during unmarshaling
	State string `json:"state,intern"` // will be interned during unmarshaling
}
