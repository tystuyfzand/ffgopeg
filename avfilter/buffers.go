package avfilter

/*
	#cgo pkg-config: libavfilter
	#include <libavfilter/buffersrc.h>
*/
import "C"

import (
	"github.com/tystuyfzand/ffgopeg/avutil"
	"unsafe"
)

// Add a frame to the buffer source.
//
// C-Function: av_buffersrc_add_frame
func (ctx *FilterContext) BufferSrcAddFrame(frame *avutil.Frame) avutil.ReturnCode {
	return avutil.NewReturnCode(int(C.av_buffersrc_add_frame((*C.struct_AVFilterContext)(ctx), (*C.struct_AVFrame)(unsafe.Pointer(frame)))))
}

// Add a frame to the buffer source.
//
// C-Function: av_buffersrc_add_frame_flags
func (ctx *FilterContext) BufferSrcAddFrameFlags(frame *avutil.Frame, flags int) avutil.ReturnCode {
	return avutil.NewReturnCode(int(C.av_buffersrc_add_frame_flags((*C.struct_AVFilterContext)(ctx), (*C.struct_AVFrame)(unsafe.Pointer(frame)), C.int(flags))))
}

// av_buffersink_get_frame

// Add a frame to the buffer source.
//
// C-Function: av_buffersink_get_frame
func (ctx *FilterContext) BufferSinkGetFrame(frame *avutil.Frame) avutil.ReturnCode {
	return avutil.NewReturnCode(int(C.av_buffersink_get_frame((*C.struct_AVFilterContext)(ctx), (*C.struct_AVFrame)(unsafe.Pointer(frame)))))
}

// Add a frame to the buffer source.
//
// C-Function: av_buffersink_get_frame_flags
func (ctx *FilterContext) BufferSinkGetFrameFlags(frame *avutil.Frame, flags int) avutil.ReturnCode {
	return avutil.NewReturnCode(int(C.av_buffersink_get_frame_flags((*C.struct_AVFilterContext)(ctx), (*C.struct_AVFrame)(unsafe.Pointer(frame)), C.int(flags))))
}