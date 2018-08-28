// +build !darwin linux

package pcre

// #cgo LDFLAGS: -lpcre
// #include <pcre.h>
// #include <string.h>
//
// void call_pcre_free(void *ptr);
//
import "C"

func (pcre *PCRE) Free() { C.call_pcre_free(unsafe.Pointer(pcre)) }

func (pcre *PCRE) Exec(extra interface{}, subject string, startOffset int, options Option, oVector []int) Error {
	subjectCStr := C.CString(subject)
	defer C.free(unsafe.Pointer(subjectCStr))

	oVectorC := make([]C.int, len(oVector))
	for n, i := range oVector {
		oVectorC[n] = C.int(i)
	}

	var oVectorPtr *C.int
	if len(oVector) > 0 {
		oVectorPtr = &oVectorC[0]
	}

	r := C.pcre_exec((*C.struct_real_pcre)(pcre), nil, subjectCStr, C.int(len(subject)), C.int(startOffset), C.int(options), oVectorPtr, C.int(len(oVector)))

	for n, i := range oVectorC {
		oVector[n] = int(i)
	}

	return Error(r)
}

func (pcre *PCRE) CaptureCount() int {
	var i C.int
	if rc := C.pcre_fullinfo((*C.struct_real_pcre)(pcre), nil, InfoCaptureCount, unsafe.Pointer(&i)); rc != 0 {
		log.Panicf("pcre_fullinfo: %v", rc)
	}
	return int(i)
}

func (pcre *PCRE) NameCount() int {
	var i C.int
	if rc := C.pcre_fullinfo((*C.struct_real_pcre)(pcre), nil, InfoNameCount, unsafe.Pointer(&i)); rc != 0 {
		log.Panicf("pcre_fullinfo: %v", rc)
	}
	return int(i)
}

func (pcre *PCRE) NameEntrySize() int {
	var i C.int
	if rc := C.pcre_fullinfo((*C.struct_real_pcre)(pcre), nil, InfoNameEntrySize, unsafe.Pointer(&i)); rc != 0 {
		log.Panicf("pcre_fullinfo: %v", rc)
	}
	return int(i)
}

func (pcre *PCRE) NameTable() []string {
	names := make([]string, pcre.CaptureCount()+1)
	if pcre.NameCount() == 0 {
		return names
	}

	var dataPtr uintptr
	if rc := C.pcre_fullinfo((*C.struct_real_pcre)(pcre), nil, InfoNameTable, unsafe.Pointer(&dataPtr)); rc != 0 {
		log.Panicf("pcre_fullinfo: %v", rc)
	}

	var data = *(*[]byte)(unsafe.Pointer(
		&reflect.SliceHeader{
			Data: dataPtr,
			Len:  pcre.NameCount() * pcre.NameEntrySize(),
			Cap:  pcre.NameCount() * pcre.NameEntrySize(),
		}))

	for i := 0; i < len(data); {
		n := (int(data[i]) << 8) | int(data[i+1])
		s := string(data[i+2 : i+pcre.NameEntrySize()-1])

		names[n] = s

		i += pcre.NameEntrySize()
	}

	return names
}