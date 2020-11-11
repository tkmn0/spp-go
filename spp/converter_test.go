package spp

import (
	"encoding/binary"
	"testing"
	"unsafe"
)

const sizeOfPayload = unsafe.Sizeof(uint8(0))
const sizeOfBodyLength = unsafe.Sizeof(uint32(0))
const sizeOfMetaData = sizeOfPayload + sizeOfBodyLength

func TestConvertBytesToPtr(t *testing.T) {
	c := NewConverter()
	s := "test message"
	buff := []byte(s)
	ptr := c.ConvertBytesToPtr(PayloadTypeBytes, buff)

	metaData := (*[sizeOfMetaData]byte)(unsafe.Pointer(ptr))[:]
	pt := Payload(metaData[0])
	l := metaData[sizeOfPayload:sizeOfMetaData]

	length := binary.LittleEndian.Uint32(l)
	data := (*[1024]byte)(unsafe.Pointer(ptr))[:]
	data = data[sizeOfMetaData : sizeOfMetaData+uintptr(length)]

	if pt != PayloadTypeBytes {
		t.Errorf("paylaod type miss match")
	}

	if string(data) != s {
		t.Errorf("data miss match")
	}
}

func TestConvertStringToPtr(t *testing.T) {
	c := NewConverter()
	s := "test string message"
	ptr := c.ConvertStringToPtr(s)

	metaData := (*[sizeOfMetaData]byte)(unsafe.Pointer(ptr))[:]
	pt := Payload(metaData[0])
	l := metaData[sizeOfPayload:sizeOfMetaData]

	length := binary.LittleEndian.Uint32(l)
	data := (*[1024]byte)(unsafe.Pointer(ptr))[:]
	data = data[sizeOfMetaData : sizeOfMetaData+uintptr(length)]

	if pt != PayloadTypeString {
		t.Errorf("paylaod type miss match")
	}

	if string(data) != s {
		t.Errorf("data miss match")
	}
}

func TestEmpty(t *testing.T) {
	c := NewConverter()
	ptr := c.ConvertBytesToPtr(PayloadTypeEmpty, nil)

	metaData := (*[sizeOfMetaData]byte)(unsafe.Pointer(ptr))[:]
	pt := Payload(metaData[0])
	l := metaData[sizeOfPayload:sizeOfMetaData]

	length := binary.LittleEndian.Uint32(l)
	data := (*[1024]byte)(unsafe.Pointer(ptr))[:]
	data = data[sizeOfMetaData : sizeOfMetaData+uintptr(length)]

	if pt != PayloadTypeEmpty {
		t.Errorf("paylaod type miss match")
	}
}
