#include <libavutil/avutil.h>
#include <stdlib.h>
#include "errorMacros.h"

int macro_AVERROR(int e) {
    return AVERROR(e);
}

const int macro_AVERROR_BSF_NOT_FOUND = AVERROR_BSF_NOT_FOUND;
const int macro_AVERROR_BUG = AVERROR_BUG;
const int macro_AVERROR_BUFFER_TOO_SMALL = AVERROR_BUFFER_TOO_SMALL;
const int macro_AVERROR_DECODER_NOT_FOUND = AVERROR_DECODER_NOT_FOUND;
const int macro_AVERROR_DEMUXER_NOT_FOUND = AVERROR_DEMUXER_NOT_FOUND;
const int macro_AVERROR_ENCODER_NOT_FOUND = AVERROR_ENCODER_NOT_FOUND;
const int macro_AVERROR_EOF = AVERROR_EOF;
const int macro_AVERROR_EXIT = AVERROR_EXIT;
const int macro_AVERROR_EXTERNAL = AVERROR_EXTERNAL;
const int macro_AVERROR_FILTER_NOT_FOUND = AVERROR_FILTER_NOT_FOUND;
const int macro_AVERROR_INVALIDDATA = AVERROR_INVALIDDATA;
const int macro_AVERROR_MUXER_NOT_FOUND = AVERROR_MUXER_NOT_FOUND;
const int macro_AVERROR_OPTION_NOT_FOUND = AVERROR_OPTION_NOT_FOUND;
const int macro_AVERROR_PATCHWELCOME = AVERROR_PATCHWELCOME;
const int macro_AVERROR_PROTOCOL_NOT_FOUND = AVERROR_PROTOCOL_NOT_FOUND;
const int macro_AVERROR_STREAM_NOT_FOUND = AVERROR_STREAM_NOT_FOUND;
const int macro_AVERROR_UNKNOWN = AVERROR_UNKNOWN;
const int macro_AVERROR_EXPERIMENTAL = AVERROR_EXPERIMENTAL;
const int macro_AVERROR_INPUT_CHANGED = AVERROR_INPUT_CHANGED;
const int macro_AVERROR_OUTPUT_CHANGED = AVERROR_OUTPUT_CHANGED;
const int macro_AVERROR_HTTP_BAD_REQUEST = AVERROR_HTTP_BAD_REQUEST;
const int macro_AVERROR_HTTP_UNAUTHORIZED = AVERROR_HTTP_UNAUTHORIZED;
const int macro_AVERROR_HTTP_FORBIDDEN = AVERROR_HTTP_FORBIDDEN;
const int macro_AVERROR_HTTP_NOT_FOUND = AVERROR_HTTP_NOT_FOUND;
const int macro_AVERROR_HTTP_OTHER_4XX = AVERROR_HTTP_OTHER_4XX;
const int macro_AVERROR_HTTP_SERVER_ERROR = AVERROR_HTTP_SERVER_ERROR;
const int macro_AV_ERROR_MAX_STRING_SIZE = AV_ERROR_MAX_STRING_SIZE;
