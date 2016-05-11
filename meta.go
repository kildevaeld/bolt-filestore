package filestore

import (
	"bytes"
	"encoding/json"
)

type Meta map[string]interface{}

func (u Meta) Marshal() ([]byte, error) {
	buffer := make([]byte, u.Size())
	_, err := u.MarshalTo(buffer)
	return buffer, err
}

func (u Meta) MarshalTo(data []byte) (n int, err error) {

	b, e := json.Marshal(u)

	if e != nil {
		return 0, e
	}
	copy(data, b)
	return len(b), nil
}

func (u *Meta) Unmarshal(data []byte) error {
	if data == nil {
		u = nil
		return nil
	}
	if len(data) == 0 {
		pu := Meta{}
		*u = pu
		return nil
	}

	var pu Meta
	if e := json.Unmarshal(data, &pu); e != nil {
		return e
	}
	*u = pu
	return nil
}

func (this Meta) Equal(that Meta) bool {
	thisdata, err := this.Marshal()
	if err != nil {
		panic(err)
	}
	thatdata, err := that.Marshal()
	if err != nil {
		panic(err)
	}
	return bytes.Equal(thisdata, thatdata)
}

func (u Meta) Size() int {
	b, e := json.Marshal(u)

	if e != nil {
		return 0
	}
	return len(b)

}

func (this Meta) Compare(that Meta) int {
	thisdata, err := this.Marshal()
	if err != nil {
		panic(err)
	}
	thatdata, err := that.Marshal()
	if err != nil {
		panic(err)
	}
	return bytes.Compare(thisdata, thatdata)
}
