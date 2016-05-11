package files

import (
	"bytes"
	"encoding/binary"
	"errors"

	"github.com/kildevaeld/percy/utils"
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

type Fid utils.Sid

func (u Fid) Marshal() ([]byte, error) {
	buffer := make([]byte, 12)
	_, err := u.MarshalTo(buffer)
	return buffer, err
}

func (u Fid) MarshalTo(data []byte) (n int, err error) {
	for i, b := range u {
		data[i] = byte(b)
	}
	return 12, nil
}

func (u *Fid) Unmarshal(data []byte) error {
	if data == nil {
		u = nil
		return nil
	}
	if len(data) == 0 {
		pu := Fid(0)
		*u = pu
		return nil
	}
	if len(data) != 8 {
		return errors.New("Fid: invalid length")
	}

	pu := Fid(data)
	*u = pu
	return nil
}

func (this Fid) Equal(that Fid) bool {
	return this == that
}

func (u Fid) Size() int {
	return 12
}

func (this Fid) Compare(that Fid) int {
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
