package files

import (
	"bytes"
	"crypto/md5"
	"crypto/rand"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"os"
	"sync/atomic"
	"time"
)

func FidHex(s string) Fid {
	d, err := hex.DecodeString(s)
	if err != nil || len(d) != 12 {
		panic(fmt.Sprintf("Invalid input to FidHex: %q", s))
	}
	return Fid(d)
}

// IsFidHex returns whether s is a valid hex representation of
// an ObjectId. See the FidHex function.
func IsFidHex(s string) bool {
	if len(s) != 24 {
		return false
	}
	_, err := hex.DecodeString(s)
	return err == nil
}

// FidCounter is atomically incremented when generating a new Fid
// using NewFid() function. It's used as a counter part of an id.
var FidCounter uint32 = readRandomUint32()

// readRandomUint32 returns a random FidCounter.
func readRandomUint32() uint32 {
	var b [4]byte
	_, err := io.ReadFull(rand.Reader, b[:])
	if err != nil {
		panic(fmt.Errorf("cannot read random object id: %v", err))
	}
	return uint32((uint32(b[0]) << 0) | (uint32(b[1]) << 8) | (uint32(b[2]) << 16) | (uint32(b[3]) << 24))
}

var machineId = readMachineId()

// readMachineId generates and returns a machine id.
// If this function fails to get the hostname it will cause a runtime error.
func readMachineId() []byte {
	var sum [3]byte
	id := sum[:]
	hostname, err1 := os.Hostname()
	if err1 != nil {
		_, err2 := io.ReadFull(rand.Reader, id)
		if err2 != nil {
			panic(fmt.Errorf("cannot get hostname: %v; %v", err1, err2))
		}
		return id
	}
	hw := md5.New()
	hw.Write([]byte(hostname))
	copy(id, hw.Sum(nil))
	return id
}

type Fid string

func NewFid() Fid {
	var b [12]byte
	// Timestamp, 4 bytes, big endian
	binary.BigEndian.PutUint32(b[:], uint32(time.Now().Unix()))
	// Machine, first 3 bytes of md5(hostname)
	b[4] = machineId[0]
	b[5] = machineId[1]
	b[6] = machineId[2]
	// Pid, 2 bytes, specs don't specify endianness, but we use big endian.
	pid := os.Getpid()
	b[7] = byte(pid >> 8)
	b[8] = byte(pid)
	// Increment, 3 bytes, big endian
	i := atomic.AddUint32(&FidCounter, 1)
	b[9] = byte(i >> 16)
	b[10] = byte(i >> 8)
	b[11] = byte(i)
	return Fid(b[:])
}

// String returns a hex string representation of the id.
// Example: FidHex("4d88e15b60f486e428412dc9").
func (id Fid) String() string {
	return fmt.Sprintf(`FidHex("%x")`, string(id))
}

// Hex returns a hex representation of the Fid.
func (id Fid) Hex() string {
	return hex.EncodeToString([]byte(id))
}

// MarshalJSON turns a bson.Fid into a json.Marshaller.
func (id Fid) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%x"`, string(id))), nil
}

var nullBytes = []byte("null")

// UnmarshalJSON turns *bson.Fid into a json.Unmarshaller.
func (id *Fid) UnmarshalJSON(data []byte) error {
	if len(data) == 2 && data[0] == '"' && data[1] == '"' || bytes.Equal(data, nullBytes) {
		*id = ""
		return nil
	}
	if len(data) != 26 || data[0] != '"' || data[25] != '"' {
		return errors.New(fmt.Sprintf("Invalid Fid in JSON: %s", string(data)))
	}
	var buf [12]byte
	_, err := hex.Decode(buf[:], data[1:25])
	if err != nil {
		return errors.New(fmt.Sprintf("Invalid Fid in JSON: %s (%s)", string(data), err))
	}
	*id = Fid(string(buf[:]))
	return nil
}

// Valid returns true if id is valid. A valid id must contain exactly 12 bytes.
func (id Fid) Valid() bool {
	return len(id) == 12
}

// byteSlice returns byte slice of id from start to end.
// Calling this function with an invalid id will cause a runtime panic.
func (id Fid) byteSlice(start, end int) []byte {
	if len(id) != 12 {
		panic(fmt.Sprintf("Invalid Fid: %q", string(id)))
	}
	return []byte(string(id)[start:end])
}

// Time returns the timestamp part of the id.
// It's a runtime error to call this method with an invalid id.
func (id Fid) Time() time.Time {
	// First 4 bytes of Fid is 32-bit big-endian seconds from epoch.
	secs := int64(binary.BigEndian.Uint32(id.byteSlice(0, 4)))
	return time.Unix(secs, 0)
}

// Machine returns the 3-byte machine id part of the id.
// It's a runtime error to call this method with an invalid id.
func (id Fid) Machine() []byte {
	return id.byteSlice(4, 7)
}

// Pid returns the process id part of the id.
// It's a runtime error to call this method with an invalid id.
func (id Fid) Pid() uint16 {
	return binary.BigEndian.Uint16(id.byteSlice(7, 9))
}

// Counter returns the incrementing value part of the id.
// It's a runtime error to call this method with an invalid id.
func (id Fid) Counter() int32 {
	b := id.byteSlice(9, 12)
	// Counter is stored as big-endian 3-byte value
	return int32(uint32(b[0])<<16 | uint32(b[1])<<8 | uint32(b[2]))
}

// Protobuf serializing
func (u Fid) Marshal() ([]byte, error) {
	buffer := make([]byte, 12)
	_, err := u.MarshalTo(buffer)
	return buffer, err
}

func (u Fid) MarshalTo(data []byte) (n int, err error) {
	copy(data, []byte(u))
	return 12, nil
}

func (u *Fid) Unmarshal(data []byte) error {
	if data == nil {
		u = nil
		return nil
	}
	if len(data) == 0 {
		pu := Fid("")
		*u = pu
		return nil
	}
	if len(data) != 12 {
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
