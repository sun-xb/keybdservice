package btsdp

/*
#cgo LDFLAGS: -lWs2_32
#include <Winsock2.h>
#include <ws2bth.h>

#include <stdio.h>

void printhex(unsigned char* p, int len, FILE* f)
{
	int i,j;
	for(i = 0;i < len;i += 16)
	{
		fprintf(f, "%08x: ", i);
		for(j = 0; j < (16 < (len - i) ? 16 : (len - i)); j++)
		{
			fprintf(f, "%02x ",p[i+j]&255);
			if(j == 7) fprintf(f, " ");
		}
		if(j < 7) fprintf(f, " ");
		for(; j < 17; j++)
			fprintf(f ,"   ");
		for(j = 0; j < (16 < (len - i) ? 16 : (len - i)); j++)
		{
			fprintf(f, "%c",(p[i+j]&255) <= 32 || (p[i+j]&255) > 127 ? '.' : (p[i+j]&255));
			if(j == 7) fprintf(f, " ");
		}
		fprintf(f, "\r\n");
	}
}

int SetService(void *buf, int len) {
	int ret = 0;
	ULONG SdpVersion = BTH_SDP_VERSION;
	HANDLE RecordHandle = 0;
	BYTE serviceBuf[sizeof(BTH_SET_SERVICE) + len - 1];
	PBTH_SET_SERVICE pService = (PBTH_SET_SERVICE)serviceBuf;
	BLOB blob;
	WSAQUERYSET wqs;

	memset(serviceBuf, 0, sizeof(serviceBuf));

	pService->pSdpVersion = &SdpVersion;
	pService->pRecordHandle = &RecordHandle;
	SET_COD_MAJOR((pService->fCodService), COD_MAJOR_COMPUTER);
	SET_COD_MINOR((pService->fCodService), COD_COMPUTER_MINOR_DESKTOP);
	SET_COD_SERVICE((pService->fCodService), COD_SERVICE_LIMITED);
	pService->ulRecordLength = len;
	memcpy(pService->pRecord, buf, len);


	blob.cbSize = sizeof(serviceBuf);
	blob.pBlobData = serviceBuf;

	memset(&wqs, 0, sizeof(wqs));
	wqs.dwSize = sizeof(wqs);
	wqs.dwNameSpace = NS_BTH;
	wqs.lpBlob = &blob;

	ret = WSASetService(&wqs, RNRSERVICE_REGISTER, 0);
	if (SOCKET_ERROR == ret) {
		printf("WSASetService failed. WSAGetLastError ret: %d\n", WSAGetLastError());
		printf("raw record data:\n");
		printhex((unsigned char*)buf, len, stdout);
	}
	return 0 == ret;
}
*/
import "C"

func SetService(record []byte) bool {
	ptr := C.CBytes(record)
	if ptr == nil {
		return false
	}
	ret := C.SetService(ptr, C.int(len(record)))
	C.free(ptr)
	return 1 == ret
}
