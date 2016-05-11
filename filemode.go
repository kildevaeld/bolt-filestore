package files

import (
	"bytes"
	"encoding/binary"
	"errors"
)

type FileMode int32

func (u FileMode) Marshal() ([]byte, error) {
	buffer := make([]byte, 8)
	_, err := u.MarshalTo(buffer)
	return buffer, err
}

func (u FileMode) MarshalTo(data []byte) (n int, err error) {

	binary.LittleEndian.PutUint32(data, uint32(u))
	return 8, nil
}

func (u *FileMode) Unmarshal(data []byte) error {
	if data == nil {
		u = nil
		return nil
	}
	if len(data) == 0 {
		pu := FileMode(0)
		*u = pu
		return nil
	}
	if len(data) != 8 {
		return errors.New("FileMode: invalid length")
	}

	pu := FileMode(binary.LittleEndian.Uint32(data))
	*u = pu
	return nil
}

func (this FileMode) Equal(that FileMode) bool {
	return this == that
}

func (u FileMode) Size() int {
	return 8
}

func (this FileMode) Compare(that FileMode) int {
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
