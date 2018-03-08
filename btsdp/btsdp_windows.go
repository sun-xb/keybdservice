package btsdp

/*
#include <Winsock2.h>

int SetService(void *buf, int len) {
	return 1;
}
*/
import "C"

func SetService(record []byte) bool {
	ptr := C.CBytes(record)
	if ptr == nil {
		return false
	}
	C.SetService(ptr, C.int(len(record)))
	C.free(ptr)
	return true
}
