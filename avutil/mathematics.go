package avutil

/*
	#cgo pkg-config: libavutil
	#include <libavutil/mathematics.h>
*/
import "C"

//
// C-Function: av_rescale
func Rescale(a, b, c int64) int64 {
	return int64(C.av_rescale(C.int64(a), C.int64(b), C.int64(c)))
}

//
// C-Function: av_rescale_q
func RescaleQ(a int64, bq Rational, cq Rational) int64 {
	return int64(C.av_rescale_q(C.int64(a), (C.struct_AVRational)(bq), (C.struct_AVRational)(cq)))
}