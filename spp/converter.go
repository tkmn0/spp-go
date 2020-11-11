package spp

import (
	"unsafe"
)

type Payload byte

const (
	PayloadTypeEmpty Payload = iota
	PayloadTypeString
	PayloadTypeBytes
)

type Converter struct{}

func NewConverter() *Converter {
	return &Converter{}
}

func (c *Converter) ConvertBytesToPtr(p Payload, body []byte) uintptr {
	length := 0
	if body != nil {
		length = len(body)
	}
	data := append([]byte{byte(p)}, (append(i32tob((uint32(length))), body...))...)
	return uintptr(unsafe.Pointer(&data[0]))
}

func (c *Converter) ConvertStringToPtr(body string) uintptr {
	data := *(*[]byte)(unsafe.Pointer(&body))
	return c.ConvertBytesToPtr(PayloadTypeString, data)
}

func i32tob(val uint32) []byte {
	r := make([]byte, 4)
	for i := uint32(0); i < 4; i++ {
		r[i] = byte((val >> (8 * i)) & 0xff)
	}
	return r
}
