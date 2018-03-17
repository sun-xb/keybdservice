package btsdp

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"strings"
)

type SdpAttribute struct {
	rw io.ReadWriter
}

func NewSdpAttribute(id uint16, attrType string, attr interface{}) *SdpAttribute {
	sa := &SdpAttribute{
		rw: bytes.NewBufferString(""),
	}
	sa.writeDataElement("uint16", id)
	sa.writeDataElement(attrType, attr)
	return sa
}

func NewSdpElement(attrType string, attr interface{}) *SdpAttribute {
	sa := &SdpAttribute{
		rw: bytes.NewBufferString(""),
	}
	sa.writeDataElement(attrType, attr)
	return sa
}

func (r *SdpAttribute) GetData() ([]byte, error) {
	return ioutil.ReadAll(r.rw)
}

func (r *SdpAttribute) writeTypeSizeDesc(typeDesc, sizeDesc int) {
	binary.Write(r.rw, binary.BigEndian, byte((typeDesc<<3)|sizeDesc))
}

func (r *SdpAttribute) writeTypeSizeDescLong(typeDesc, sizeDesc int) {
	switch {
	case sizeDesc < 0x100:
		binary.Write(r.rw, binary.BigEndian, uint8((typeDesc<<3)|5))
		binary.Write(r.rw, binary.BigEndian, uint8(sizeDesc))
	case sizeDesc < 0x10000:
		binary.Write(r.rw, binary.BigEndian, uint8((typeDesc<<3)|6))
		binary.Write(r.rw, binary.BigEndian, uint16(sizeDesc))
	default:
		binary.Write(r.rw, binary.BigEndian, uint8((typeDesc<<3)|7))
		binary.Write(r.rw, binary.BigEndian, uint32(sizeDesc))
	}
}

func (r *SdpAttribute) writeDataElement(typeDesc string, value interface{}) {
	getNum := func(v interface{}) (low, high uint64) {
		switch value := v.(type) {
		case int:
			low = uint64(value)
		case int8:
			low = uint64(value)
		case int16:
			low = uint64(value)
		case int32:
			low = uint64(value)
		case int64:
			low = uint64(value)
		case uint:
			low = uint64(value)
		case uint8:
			low = uint64(value)
		case uint16:
			low = uint64(value)
		case uint32:
			low = uint64(value)
		case uint64:
			low = uint64(value)
		default:
			panic(fmt.Sprint("not implemented: ", value))
		}
		return
	}
	switch typeDesc {
	case "nil":
		r.writeTypeSizeDesc(0, 0)
	case "uint8":
		r.writeTypeSizeDesc(1, 0)
		low, _ := getNum(value)
		binary.Write(r.rw, binary.BigEndian, uint8(low))
	case "uint16":
		r.writeTypeSizeDesc(1, 1)
		low, _ := getNum(value)
		binary.Write(r.rw, binary.BigEndian, uint16(low))
	case "uint32":
		r.writeTypeSizeDesc(1, 2)
		low, _ := getNum(value)
		binary.Write(r.rw, binary.BigEndian, uint32(low))
	case "uint64":
		r.writeTypeSizeDesc(1, 3)
		low, _ := getNum(value)
		binary.Write(r.rw, binary.BigEndian, uint64(low))
	case "uint128":
		r.writeTypeSizeDesc(1, 4)
		low, high := getNum(value)
		binary.Write(r.rw, binary.BigEndian, uint64(high))
		binary.Write(r.rw, binary.BigEndian, uint64(low))
	case "int8":
		r.writeTypeSizeDesc(2, 0)
		low, _ := getNum(value)
		binary.Write(r.rw, binary.BigEndian, uint8(low))
	case "int16":
		r.writeTypeSizeDesc(2, 1)
		low, _ := getNum(value)
		binary.Write(r.rw, binary.BigEndian, uint16(low))
	case "int32":
		r.writeTypeSizeDesc(2, 2)
		low, _ := getNum(value)
		binary.Write(r.rw, binary.BigEndian, uint32(low))
	case "int64":
		r.writeTypeSizeDesc(2, 3)
		low, _ := getNum(value)
		binary.Write(r.rw, binary.BigEndian, uint64(low))
	case "int128":
		r.writeTypeSizeDesc(2, 4)
		low, high := getNum(value)
		binary.Write(r.rw, binary.BigEndian, uint64(high))
		binary.Write(r.rw, binary.BigEndian, uint64(low))
	case "uuid":
		uuid, ok := value.(string)
		if !ok {
			panic("error uuid type.")
		}
		if len(uuid) == 36 {
			uuid = strings.Replace(uuid, "-", "", -1)
		}
		switch len(uuid) {
		case 4:
			r.writeTypeSizeDesc(3, 1)
		case 8:
			r.writeTypeSizeDesc(3, 2)
		case 32:
			r.writeTypeSizeDesc(3, 4)
		default:
			panic("invalid uuid length.")
		}
		if dst, err := hex.DecodeString(uuid); err == nil {
			binary.Write(r.rw, binary.BigEndian, dst)
		}
	case "string":
		str, ok := value.(string)
		if !ok {
			panic("error string type.")
		}
		r.writeTypeSizeDescLong(4, len(str))
		binary.Write(r.rw, binary.BigEndian, []byte(str))
	case "bool":
		b, ok := value.(bool)
		if !ok {
			panic("error bool type.")
		}
		r.writeTypeSizeDesc(5, 0)
		if b {
			binary.Write(r.rw, binary.BigEndian, byte(1))
		} else {
			binary.Write(r.rw, binary.BigEndian, byte(0))
		}
	case "sequence":
		seq, ok := value.([]*SdpAttribute)
		if !ok {
			panic("error sequence type.")
		}
		packedseq := make([]byte, 0)
		for _, v := range seq {
			if e, err := v.GetData(); err == nil {
				packedseq = append(packedseq, e...)
			}
		}
		r.writeTypeSizeDescLong(6, len(packedseq))
		binary.Write(r.rw, binary.BigEndian, packedseq)
	case "alternative":
		seq, ok := value.([]*SdpAttribute)
		if !ok {
			panic("error sequence type.")
		}
		packedseq := make([]byte, 0)
		for _, v := range seq {
			if e, err := v.GetData(); err == nil {
				packedseq = append(packedseq, e...)
			}
		}
		r.writeTypeSizeDescLong(7, len(packedseq))
		binary.Write(r.rw, binary.BigEndian, packedseq)
	case "url":
		url, ok := value.(string)
		if !ok {
			panic("error url type.")
		}
		r.writeTypeSizeDescLong(8, len(url))
		binary.Write(r.rw, binary.BigEndian, []byte(url))
	default:
		panic("not implemented type.")
	}
}
