package main

import "C"
import (
	"encoding/base64"
	"unsafe"
)

//export CharPtrToB64
func CharPtrToB64(src *C.char) unsafe.Pointer {
	var encode string = base64.StdEncoding.EncodeToString([]byte(C.GoString(src)))
	return unsafe.Pointer(C.CString(encode))
}

//export B64ToCharPtr
func B64ToCharPtr(enc *C.char) unsafe.Pointer {
	decode, _ := base64.StdEncoding.DecodeString(C.GoString(enc))
	return unsafe.Pointer(C.CString(string(decode)))
}

func main() {}