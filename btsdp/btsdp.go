package btsdp

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
			NewSdpElement("uuid", "0100"),
			NewSdpElement("uint16", 0x11),
		}),
		NewSdpElement("sequence", []*SdpAttribute{
			NewSdpElement("uuid", "0011"),
		}),
	})
	browseGroupList := NewSdpAttribute(BROWSE_GROUP_LIST_ATTRID, "sequence", []*SdpAttribute{
		NewSdpElement("uuid", "1002"),
	})
	languageBase := NewSdpAttribute(LANGUAGE_BASE_ATTRID_LIST_ATTRID, "sequence", []*SdpAttribute{
		NewSdpElement("uint16", (0x65<<8)|0x6e),
		NewSdpElement("uint16", 106),
		NewSdpElement("uint16", 0x0100),
	})
	profileDescList := NewSdpAttribute(BLUETOOTH_PROFILE_DESCRIPTOR_LIST_ATTRID, "sequence", []*SdpAttribute{
		NewSdpElement("uuid", HID_PROFILE_ID),
		NewSdpElement("uint16", 0x0100),
	})
	addProtoDescList := NewSdpAttribute(ADD_PROTO_DESC_LIST_ATTRID, "sequence", []*SdpAttribute{
		NewSdpElement("sequence", []*SdpAttribute{
			NewSdpElement("uuid", "0100"),
			NewSdpElement("uint16", 0x13),
		}),
		NewSdpElement("sequence", []*SdpAttribute{
			NewSdpElement("uuid", "0011"),
		}),
	})
	serviceName := NewSdpAttribute(SERVICE_NAME_ATTRID, "string", "keyboard emulation service")
	serviceDesc := NewSdpAttribute(SERVICE_DESCRIPTION_ATTRID, "string", "keyboard emulation service")
	providerName := NewSdpAttribute(PROVIDER_NAME_ATTRID, "string", "sun")

	hidDeviceReleaseNumber := NewSdpAttribute(SDP_ATTR_HID_DEVICE_RELEASE_NUMBER, "uint16", 1234)
	hidParserVersion := NewSdpAttribute(SDP_ATTR_HID_PARSER_VERSION, "uint16", 1234)
	hidDeviceSubClass := NewSdpAttribute(SDP_ATTR_HID_DEVICE_SUBCLASS, "uint8", 12)
	hidCountryCode := NewSdpAttribute(SDP_ATTR_HID_COUNTRY_CODE, "uint8", 12)
	hidVirtualCable := NewSdpAttribute(SDP_ATTR_HID_VIRTUAL_CABLE, "bool", true)
	hidReconnectInitiate := NewSdpAttribute(SDP_ATTR_HID_RECONNECT_INITIATE, "bool", true)
	//hidDescList := NewSdpAttribute(SDP_ATTR_HID_DESCRIPTOR_LIST, "")
	//hidLangIdBaseList := NewSdpAttribute(SDP_ATTR_HID_LANG_ID_BASE_LIST, "")
	hidSdpDisable := NewSdpAttribute(SDP_ATTR_HID_SDP_DISABLE, "bool", false)
	hidBatteryPower := NewSdpAttribute(SDP_ATTR_HID_BATTERY_POWER, "bool", true)
	hidRemoteWakeUp := NewSdpAttribute(SDP_ATTR_HID_REMOTE_WAKEUP, "bool", true)
	hidProfileVersion := NewSdpAttribute(SDP_ATTR_HID_PROFILE_VERSION, "uint16", 1234)
	hidSuperVersionTimeout := NewSdpAttribute(SDP_ATTR_HID_SUPERVISION_TIMEOUT, "uint16", 1234)
	hidNormallyConnectable := NewSdpAttribute(SDP_ATTR_HID_NORMALLY_CONNECTABLE, "bool", true)
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
		//hidDescList,
		//hidLangIdBaseList,
		hidSdpDisable,
		hidBatteryPower,
		hidRemoteWakeUp,
		hidProfileVersion,
		hidSuperVersionTimeout,
		hidNormallyConnectable,
		hidBootVersion,
	})
	record.GetData()
}
