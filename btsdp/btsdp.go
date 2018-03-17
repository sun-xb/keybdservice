package btsdp

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

	PSM_HID_CONTROL   = 0x0011
	PSM_HID_INTERRUPT = 0x0013

	L2CAP_UUID = "0100"
	HIDP_UUID  = "0011"

	PUBLIC_BROWSE_GROUP = "1002"
)

var hidReportDesc = []byte{
	0x05, 0x01, // UsagePage GenericDesktop
	0x09, 0x02, // Usage Mouse
	0xA1, 0x01, // Collection Application
	0x85, 0x01, // REPORT ID: 1
	0x09, 0x01, // Usage Pointer
	0xA1, 0x00, // Collection Physical
	0x05, 0x09, // UsagePage Buttons
	0x19, 0x01, // UsageMinimum 1
	0x29, 0x03, // UsageMaximum 3
	0x15, 0x00, // LogicalMinimum 0
	0x25, 0x01, // LogicalMaximum 1
	0x75, 0x01, // ReportSize 1
	0x95, 0x03, // ReportCount 3
	0x81, 0x02, // Input data variable absolute
	0x75, 0x05, // ReportSize 5
	0x95, 0x01, // ReportCount 1
	0x81, 0x01, // InputConstant (padding)
	0x05, 0x01, // UsagePage GenericDesktop
	0x09, 0x30, // Usage X
	0x09, 0x31, // Usage Y
	0x09, 0x38, // Usage ScrollWheel
	0x15, 0x81, // LogicalMinimum -127
	0x25, 0x7F, // LogicalMaximum +127
	0x75, 0x08, // ReportSize 8
	0x95, 0x02, // ReportCount 3
	0x81, 0x06, // Input data variable relative
	0xC0, 0xC0, // EndCollection EndCollection
	0x05, 0x01, // UsagePage GenericDesktop
	0x09, 0x06, // Usage Keyboard
	0xA1, 0x01, // Collection Application
	0x85, 0x02, // REPORT ID: 2
	0xA1, 0x00, // Collection Physical
	0x05, 0x07, // UsagePage Keyboard
	0x19, 0xE0, // UsageMinimum 224
	0x29, 0xE7, // UsageMaximum 231
	0x15, 0x00, // LogicalMinimum 0
	0x25, 0x01, // LogicalMaximum 1
	0x75, 0x01, // ReportSize 1
	0x95, 0x08, // ReportCount 8
	0x81, 0x02, // **Input data variable absolute
	0x95, 0x08, // ReportCount 8
	0x75, 0x08, // ReportSize 8
	0x15, 0x00, // LogicalMinimum 0
	0x25, 0x65, // LogicalMaximum 101
	0x05, 0x07, // UsagePage Keycodes
	0x19, 0x00, // UsageMinimum 0
	0x29, 0x65, // UsageMaximum 101
	0x81, 0x00, // **Input DataArray
	0xC0, 0xC0, // EndCollection
}

type Sdp struct {
}

func New() *Sdp {
	return &Sdp{}
}

func (s *Sdp) RegisterService() {
	clsIdList := NewSdpAttribute(SERVICE_CLASS_ID_LIST_ATTRID, "sequence", []*SdpAttribute{
		NewSdpElement("uuid", HID_SVCLASS_ID),
	})
	protoDescList := NewSdpAttribute(PROTOCOL_DESCRIPTOR_LIST_ATTRID, "sequence", []*SdpAttribute{
		NewSdpElement("sequence", []*SdpAttribute{
			NewSdpElement("uuid", L2CAP_UUID),
			NewSdpElement("uint16", PSM_HID_CONTROL),
		}),
		NewSdpElement("sequence", []*SdpAttribute{
			NewSdpElement("uuid", HIDP_UUID),
		}),
	})
	browseGroupList := NewSdpAttribute(BROWSE_GROUP_LIST_ATTRID, "sequence", []*SdpAttribute{
		NewSdpElement("uuid", PUBLIC_BROWSE_GROUP),
	})
	languageBase := NewSdpAttribute(LANGUAGE_BASE_ATTRID_LIST_ATTRID, "sequence", []*SdpAttribute{
		NewSdpElement("uint16", (0x65<<8)|0x6e),
		NewSdpElement("uint16", 106),
		NewSdpElement("uint16", 0x0100),
	})
	profileDescList := NewSdpAttribute(BLUETOOTH_PROFILE_DESCRIPTOR_LIST_ATTRID, "sequence", []*SdpAttribute{
		NewSdpElement("sequence", []*SdpAttribute{
			NewSdpElement("uuid", HID_PROFILE_ID),
			NewSdpElement("uint16", 0x0100),
		}),
	})
	addProtoDescList := NewSdpAttribute(ADD_PROTO_DESC_LIST_ATTRID, "sequence", []*SdpAttribute{
		NewSdpElement("sequence", []*SdpAttribute{
			NewSdpElement("sequence", []*SdpAttribute{
				NewSdpElement("uuid", L2CAP_UUID),
				NewSdpElement("uint16", PSM_HID_INTERRUPT),
			}),
			NewSdpElement("sequence", []*SdpAttribute{
				NewSdpElement("uuid", HIDP_UUID),
			}),
		}),
	})
	serviceName := NewSdpAttribute(SERVICE_NAME_ATTRID, "string", "keyboard emulation service")
	serviceDesc := NewSdpAttribute(SERVICE_DESCRIPTION_ATTRID, "string", "keyboard emulation service description")
	providerName := NewSdpAttribute(PROVIDER_NAME_ATTRID, "string", "sun")

	hidDeviceReleaseNumber := NewSdpAttribute(SDP_ATTR_HID_DEVICE_RELEASE_NUMBER, "uint16", 0x0100)
	hidParserVersion := NewSdpAttribute(SDP_ATTR_HID_PARSER_VERSION, "uint16", 0x0111)
	hidDeviceSubClass := NewSdpAttribute(SDP_ATTR_HID_DEVICE_SUBCLASS, "uint8", 0x80)
	hidCountryCode := NewSdpAttribute(SDP_ATTR_HID_COUNTRY_CODE, "uint8", 0x21)
	hidVirtualCable := NewSdpAttribute(SDP_ATTR_HID_VIRTUAL_CABLE, "bool", true)
	hidReconnectInitiate := NewSdpAttribute(SDP_ATTR_HID_RECONNECT_INITIATE, "bool", true)
	hidDescList := NewSdpAttribute(SDP_ATTR_HID_DESCRIPTOR_LIST, "sequence", []*SdpAttribute{
		NewSdpElement("sequence", []*SdpAttribute{
			NewSdpElement("uint8", 0x22),
			NewSdpElement("string", string(hidReportDesc)),
		}),
	})
	hidLangIdBaseList := NewSdpAttribute(SDP_ATTR_HID_LANG_ID_BASE_LIST, "sequence", []*SdpAttribute{
		NewSdpElement("sequence", []*SdpAttribute{
			NewSdpElement("uint16", 0x409),
			NewSdpElement("uint16", 0x100),
		}),
	})
	hidSdpDisable := NewSdpAttribute(SDP_ATTR_HID_SDP_DISABLE, "bool", false)
	hidBatteryPower := NewSdpAttribute(SDP_ATTR_HID_BATTERY_POWER, "bool", true)
	hidRemoteWakeUp := NewSdpAttribute(SDP_ATTR_HID_REMOTE_WAKEUP, "bool", true)
	hidProfileVersion := NewSdpAttribute(SDP_ATTR_HID_PROFILE_VERSION, "uint16", 0x0100)
	hidSuperVersionTimeout := NewSdpAttribute(SDP_ATTR_HID_SUPERVISION_TIMEOUT, "uint16", 0x0c80)
	hidNormallyConnectable := NewSdpAttribute(SDP_ATTR_HID_NORMALLY_CONNECTABLE, "bool", false)
	hidBootVersion := NewSdpAttribute(SDP_ATTR_HID_BOOT_DEVICE, "bool", true)
	record := NewSdpElement("sequence", []*SdpAttribute{
		clsIdList,
		protoDescList,
		browseGroupList,
		languageBase,
		profileDescList,
		addProtoDescList,
		serviceName,
		serviceDesc,
		providerName,
		hidDeviceReleaseNumber,
		hidParserVersion,
		hidDeviceSubClass,
		hidCountryCode,
		hidVirtualCable,
		hidReconnectInitiate,
		hidDescList,
		hidLangIdBaseList,
		hidSdpDisable,
		hidBatteryPower,
		hidRemoteWakeUp,
		hidProfileVersion,
		hidSuperVersionTimeout,
		hidNormallyConnectable,
		hidBootVersion,
	})
	data, err := record.GetData()
	if err != nil {
		return
	}

	SetService(data)
}
