package avutil

/*
#cgo pkg-config: libavutil
#include <libavutil/avutil.h>
*/
import "C"
import "unsafe"

func AvMalloc(s int) unsafe.Pointer {
	return unsafe.Pointer(C.av_malloc(C.size_t(s)))
}

//Free a memory block which has been allocated with av_malloc(z)() or av_realloc().
func AvFree(p unsafe.Pointer) {
	C.av_free(p)
}
