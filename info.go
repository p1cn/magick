package magick

// #include <string.h>
// #include <magick/api.h>
import "C"

import (
	"runtime"
	"unsafe"
)

// Info is used to specify the encoding parameters like
// format and quality when encoding and image.
type Info struct {
	info *C.ImageInfo
}

// Format returns the format used for encoding the image.
func (in *Info) Format() string {
	return C.GoString(&in.info.magick[0])
}

// SetFormat sets the image format for encoding this image.
// See http://www.graphicsmagick.org for a list of supported
// formats.
func (in *Info) SetFormat(format string) {
	if format == "" {
		in.info.magick[0] = 0
	} else {
		s := C.CString(format)
		defer C.free(unsafe.Pointer(s))
		C.strncpy(&in.info.magick[0], s, C.MaxTextExtent)
	}
}

// Quality returns the quality used when compressing the image.
// This parameter does not affect all formats.
func (in *Info) Quality() uint64 {
	return uint64(in.info.quality)
}

// SetQuality sets the quality used when compressing the image.
// This parameter does not affect all formats.
func (in *Info) SetQuality(q uint64) {
	in.info.quality = magicUlong(q)
}

// Colorspace returns the colorspace used when encoding the image.
func (in *Info) Colorspace() Colorspace {
	return Colorspace(in.info.colorspace)
}

// SetColorspace set the colorspace used when encoding the image.
// Note that not all colorspaces are supported for encoding. See
// the documentation on Colorspace.
func (in *Info) SetColorspace(cs Colorspace) {
	in.info.colorspace = C.ColorspaceType(cs)
}

// NewInfo returns a newly allocated *Info structure. Do not
// create Info objects directly, since they need to allocate
// some internal structures while being created.
func NewInfo() *Info {
	cinfo := C.CloneImageInfo(nil)
	info := new(Info)
	info.info = cinfo
	runtime.SetFinalizer(info, func(i *Info) {
		if i.info != nil {
			C.DestroyImageInfo(i.info)
			i.info = nil
		}
	})
	return info
}

//MagickPassFail AddDefinitions( ImageInfo *image_info, const char *options );
//The format of the string is "key1[=[value1]],key2[=[value2]],...".
func (in *Info) AddDefinitions(def string) error {
	d := C.CString(def)
	defer C.free(unsafe.Pointer(d))

	var ex C.ExceptionInfo
	C.GetExceptionInfo(&ex)
	defer C.DestroyExceptionInfo(&ex)

	result := int(C.AddDefinitions(in.info, d, &ex))
	if result <= 0 {
		return exError(&ex, "AddDefinitions")
	}

	return nil
}
