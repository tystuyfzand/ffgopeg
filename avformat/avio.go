package avformat

/*
#cgo pkg-config: libavformat
#include <libavformat/avio.h>

extern int avioReadPacket(void *data, uint8_t *buf, int buf_size);
extern int64_t avioSeek(void *data, int64_t pos, int whence);
*/
import "C"
import (
	"github.com/mattn/go-pointer"
	"github.com/tystuyfzand/ffgopeg/avutil"
	"io"
	"log"
	"unsafe"
)

type GoAvioReadSeeker struct {
	io.Closer

	reader   io.ReadSeeker
	ctx      *IOContext
	pointer  unsafe.Pointer
	goBuffer []byte
}

func (r *GoAvioReadSeeker) Close() error {
	log.Println("Close")
	pointer.Unref(r.pointer)
	avutil.AvFree(unsafe.Pointer(r.ctx.buffer))
	return nil
}

func AvioAllocContext(reader io.ReadSeeker, bufferSize int) (*IOContext, *GoAvioReadSeeker) {
	buffer := (*C.uchar)(avutil.AvMalloc(bufferSize))

	wrapper := &GoAvioReadSeeker{reader: reader, goBuffer: make([]byte, bufferSize)}

	wrapper.pointer = pointer.Save(wrapper)

	var ptrRead, ptrWrite, ptrSeek *[0]byte = nil, nil, nil

	ptrRead = (*[0]byte)(C.avioReadPacket)
	ptrSeek = (*[0]byte)(C.avioSeek)

	ctx := (*IOContext)(C.avio_alloc_context(buffer, C.int(bufferSize), 0, wrapper.pointer, ptrRead, ptrWrite, ptrSeek))

	wrapper.ctx = ctx

	return ctx, wrapper
}

//export avioReadPacket
func avioReadPacket(data unsafe.Pointer, buf *C.uint8_t, size C.int) C.int {
	log.Println("Read packet")
	v := pointer.Restore(data).(*GoAvioReadSeeker)

	n, err := v.reader.Read(v.goBuffer)

	if err != nil {
		return 1
	}

	if n >= 0 {
		C.memcpy(unsafe.Pointer(buf), unsafe.Pointer(&v.goBuffer[0]), C.size_t(n))
	}

	return 0
}

//export avioSeek
func avioSeek(data unsafe.Pointer, pos C.int64_t, whence C.int) C.int64_t {
	log.Println("Seek", int64(pos))
	v := pointer.Restore(data).(*GoAvioReadSeeker)

	endPos, err := v.reader.Seek(int64(pos), int(whence))

	if err != nil {
		panic("Unable to seek: " + err.Error())
	}

	return C.int64_t(endPos)
}
