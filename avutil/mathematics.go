package avutil

/*
	#cgo pkg-config: libavutil
	#include <libavutil/mathematics.h>
*/
import "C"

//
// C-Function: av_make_q
func MakeQ(num, den int) Rational {
	return (Rational)(C.av_make_q(C.int(num), C.int(den)))
}

//
// C-Function: av_rescale
func Rescale(a, b, c int64) int64 {
	return int64(C.av_rescale(C.int64_t(a), C.int64_t(b), C.int64_t(c)))
}

//
// C-Function: av_rescale_q
func RescaleQ(a int64, bq Rational, cq Rational) int64 {
	return int64(C.av_rescale_q(C.int64_t(a), (C.struct_AVRational)(bq), (C.struct_AVRational)(cq)))
}
