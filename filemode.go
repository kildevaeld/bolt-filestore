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

func (m FileMode) String() string {
	const str = "dalTLDpSugct"
	var buf [32]byte // Mode is uint32.
	w := 0
	for i, c := range str {
		if m&(1<<uint(32-1-i)) != 0 {
			buf[w] = byte(c)
			w++
		}
	}
	if w == 0 {
		buf[w] = '-'
		w++
	}
	const rwx = "rwxrwxrwx"
	for i, c := range rwx {
		if m&(1<<uint(9-1-i)) != 0 {
			buf[w] = byte(c)
		} else {
			buf[w] = '-'
		}
		w++
	}
	return string(buf[:w])
}
