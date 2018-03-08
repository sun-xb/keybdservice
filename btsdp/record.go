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

const (
	SERVICE_RECORD_HANDLE_ATTRID             = 0x0000
	SERVICE_CLASS_ID_LIST_ATTRID             = 0x0001
	SERVICE_RECORD_STATE_ATTRID              = 0x0002
	SERVICE_ID_ATTRID                        = 0x0003
	PROTOCOL_DESCRIPTOR_LIST_ATTRID          = 0x0004
	BROWSE_GROUP_LIST_ATTRID                 = 0x0005
	LANGUAGE_BASE_ATTRID_LIST_ATTRID         = 0x0006
	SERVICE_INFO_TIME_TO_LIVE_ATTRID         = 0x0007
	SERVICE_AVAILABILITY_ATTRID              = 0x0008
	BLUETOOTH_PROFILE_DESCRIPTOR_LIST_ATTRID = 0x0009
	DOCUMENTATION_URL_ATTRID                 = 0x000a
	CLIENT_EXECUTABLE_URL_ATTRID             = 0x000b
	ICON_URL_ATTRID                          = 0x000c
	ADD_PROTO_DESC_LIST_ATTRID               = 0x000d
	SERVICE_NAME_ATTRID                      = 0x0100
	SERVICE_DESCRIPTION_ATTRID               = 0x0101
	PROVIDER_NAME_ATTRID                     = 0x0102

	SDP_ATTR_HID_DEVICE_RELEASE_NUMBER = 0x0200
	SDP_ATTR_HID_PARSER_VERSION        = 0x0201
	SDP_ATTR_HID_DEVICE_SUBCLASS       = 0x0202
	SDP_ATTR_HID_COUNTRY_CODE          = 0x0203
	SDP_ATTR_HID_VIRTUAL_CABLE         = 0x0204
	SDP_ATTR_HID_RECONNECT_INITIATE    = 0x0205
	SDP_ATTR_HID_DESCRIPTOR_LIST       = 0x0206
	SDP_ATTR_HID_LANG_ID_BASE_LIST     = 0x0207
	SDP_ATTR_HID_SDP_DISABLE           = 0x0208
	SDP_ATTR_HID_BATTERY_POWER         = 0x0209
	SDP_ATTR_HID_REMOTE_WAKEUP         = 0x020a
	SDP_ATTR_HID_PROFILE_VERSION       = 0x020b
	SDP_ATTR_HID_SUPERVISION_TIMEOUT   = 0x020c
	SDP_ATTR_HID_NORMALLY_CONNECTABLE  = 0x020d
	SDP_ATTR_HID_BOOT_DEVICE           = 0x020e

	HID_SVCLASS_ID = "1124"
	HID_PROFILE_ID = HID_SVCLASS_ID
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
		binary.Write(r.rw, binary.BigEndian, byte((typeDesc<<3)|5))
		binary.Write(r.rw, binary.BigEndian, uint8(sizeDesc))
	case sizeDesc < 0x10000:
		binary.Write(r.rw, binary.BigEndian, byte((typeDesc<<3)|6))
		binary.Write(r.rw, binary.BigEndian, uint16(sizeDesc))
	default:
		binary.Write(r.rw, binary.BigEndian, byte((typeDesc<<3)|7))
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
		binary.Write(r.rw, binary.BigEndian, str)
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
		binary.Write(r.rw, binary.BigEndian, url)
	default:
		panic("not implemented type.")
	}
}
