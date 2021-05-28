package main

//#cgo CFLAGS: -I/usr/local/include/wkhtmltox
//#cgo LDFLAGS: -L/usr/local/lib -lwkhtmltox
//#include <pdf.h>
import "C"
import "log"

// Setting defines wkhtmltopdf setting.
type Setting struct {
	Attr, Value string
}

func main() {
	C.wkhtmltopdf_init(0)

	gs := C.wkhtmltopdf_create_global_settings()
	os := C.wkhtmltopdf_create_object_settings()

	objectSettings := []Setting{
		{
			Attr:  "page",
			Value: "http://wkhtmltopdf.org/libwkhtmltox/",
		},
		Setting{
			Attr:  "footer.center",
			Value: "[page] / [toPage]",
		},
	}
	C.wkhtmltopdf_set_global_setting(gs, C.CString("out"), C.CString("/tmp/test.pdf"))
	for _, setting := range objectSettings {
		C.wkhtmltopdf_set_object_setting(os, C.CString(setting.Attr), C.CString(setting.Value))
	}

	c := C.wkhtmltopdf_create_converter(gs)

	C.wkhtmltopdf_add_object(c, os, nil)

	if C.wkhtmltopdf_convert(c) != 1 {
		log.Fatalln("http error code", C.wkhtmltopdf_http_error_code(c))
	}

	C.wkhtmltopdf_destroy_converter(c)

	C.wkhtmltopdf_deinit()
}
