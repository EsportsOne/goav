package avutil

//#cgo pkg-config: libavutil
//#include <libavutil/avutil.h>
import "C"

const (
	AV_PIX_FMT_GRAY9    = C.AV_PIX_FMT_GRAY9
	AV_PIX_FMT_BGR24    = C.AV_PIX_FMT_BGR24
	AV_PIX_FMT_RGB24    = C.AV_PIX_FMT_RGB24
	AV_PIX_FMT_RGBA     = C.AV_PIX_FMT_RGBA
	AV_PIX_FMT_YUV420P  = C.AV_PIX_FMT_YUV420P
	AV_PIX_FMT_YUVJ420P = C.AV_PIX_FMT_YUVJ420P
)

// PixelFormatFromString returns a pixel format from a string
func PixelFormatFromString(i string) PixelFormat {
	switch i {
	case "bgr24":
		return AV_PIX_FMT_BGR24
	case "rgb24":
		return AV_PIX_FMT_RGB24
	case "rgba":
		return AV_PIX_FMT_RGBA
	case "yuv420p":
		return AV_PIX_FMT_YUV420P
	case "yuvj420p":
		return AV_PIX_FMT_YUVJ420P
	case "gray":
		return AV_PIX_FMT_GRAY9
	default:
		return -1
	}
}
