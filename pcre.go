package pcre

// #cgo LDFLAGS: -lpcre
// #include <pcre.h>
// #include <string.h>
//
// void call_pcre_free(void* ptr) {
//     pcre_free(ptr);
// }
//
import "C"

import (
	"errors"
	"unsafe"
)

const (
	Major  = C.PCRE_MAJOR
	Minor     = C.PCRE_MINOR
	Date      = C.PCRE_DATE
)

const StudyJITCompile = C.PCRE_STUDY_JIT_COMPILE

type Option int

const (
	Caseless      = C.PCRE_CASELESS
	Multiline     = C.PCRE_MULTILINE
	DotAll        = C.PCRE_DOTALL
	Extended      = C.PCRE_EXTENDED
	Anchored      = C.PCRE_ANCHORED
	DollarEndOnly = C.PCRE_DOLLAR_ENDONLY
	Extra         = C.PCRE_EXTRA
	NotBOL        = C.PCRE_NOTBOL
	NotEOL        = C.PCRE_NOTEOL
	UnGreedy         = C.PCRE_UNGREEDY
	NotEmpty         = C.PCRE_NOTEMPTY
	UTF8             = C.PCRE_UTF8
	UTF16            = C.PCRE_UTF16
	NoAutoCapture    = C.PCRE_NO_AUTO_CAPTURE
	NoUTF8Check      = C.PCRE_NO_UTF8_CHECK
	NoUTF16Check     = C.PCRE_NO_UTF16_CHECK
	AutoCallout      = C.PCRE_AUTO_CALLOUT
	PartialSoft      = C.PCRE_PARTIAL_SOFT
	Partial          = C.PCRE_PARTIAL
	DFAShortest      = C.PCRE_DFA_SHORTEST
	DFARestart       = C.PCRE_DFA_RESTART
	FirstLine        = C.PCRE_FIRSTLINE
	DupNames         = C.PCRE_DUPNAMES
	NewlineCR        = C.PCRE_NEWLINE_CR
	NewlineLF        = C.PCRE_NEWLINE_LF
	NewlineCRLF      = C.PCRE_NEWLINE_CRLF
	NewlineAny       = C.PCRE_NEWLINE_ANY
	NewlineAnyCRLF   = C.PCRE_NEWLINE_ANYCRLF
	BSRAnyCRLF       = C.PCRE_BSR_ANYCRLF
	BSRUnicode       = C.PCRE_BSR_UNICODE
	JavascriptCompat = C.PCRE_JAVASCRIPT_COMPAT
	NoStartOptimize  = C.PCRE_NO_START_OPTIMIZE
	NoStartOptimise  = C.PCRE_NO_START_OPTIMISE // <3 Rule britannia
	PartialHard      = C.PCRE_PARTIAL_HARD
	NotEmptyAtStart  = C.PCRE_NOTEMPTY_ATSTART
	UCP              = C.PCRE_UCP
)

type Info int

const (
	InfoOptions       = C.PCRE_INFO_OPTIONS
	InfoSize          = C.PCRE_INFO_SIZE
	InfoCaptureCount  = C.PCRE_INFO_CAPTURECOUNT
	InfoBackrefMax    = C.PCRE_INFO_BACKREFMAX
	InfoFirstByte     = C.PCRE_INFO_FIRSTBYTE
	InfoFirstChar     = C.PCRE_INFO_FIRSTCHAR // For backwards compatibility
	InfoFirstTable    = C.PCRE_INFO_FIRSTTABLE
	InfoLastLiteral   = C.PCRE_INFO_LASTLITERAL
	InfoNameEntrySize = C.PCRE_INFO_NAMEENTRYSIZE
	InfoNameCount     = C.PCRE_INFO_NAMECOUNT
	InfoNameTable     = C.PCRE_INFO_NAMETABLE
	InfoStudySize     = C.PCRE_INFO_STUDYSIZE
	InfoDefaultTables = C.PCRE_INFO_DEFAULT_TABLES
	InfoOkPartial     = C.PCRE_INFO_OKPARTIAL
	InfoJchanged      = C.PCRE_INFO_JCHANGED
	InfoHasCRorLF     = C.PCRE_INFO_HASCRORLF
	InfoMinLength     = C.PCRE_INFO_MINLENGTH
	InfoJIT           = C.PCRE_INFO_JIT
	InfoJITSize       = C.PCRE_INFO_JITSIZE
	InfoMaxLookBehind = C.PCRE_INFO_MAXLOOKBEHIND
)

type PCRE C.struct_real_pcre8_or_16

func Compile(expr string, options Option, _ interface{}) (*PCRE, error) {
	var (
		errPtr    *C.char
		errOffset C.int
	)

	pattern := C.CString(expr)
	defer C.free(unsafe.Pointer(pattern))

	re := C.pcre_compile(pattern, C.int(options), &errPtr, &errOffset, nil)
	if re == nil {
		return nil, errors.New(C.GoString(errPtr))
	}

	return (*PCRE)(re), nil
}


type PCREExtra C.struct_pcre_extra
func Study(code *PCRE, options Option, _ interface{}) (*PCREExtra, error) {
	var errPtr *C.char

	extra := C.pcre_study((*C.struct_real_pcre8_or_16)(code), C.int(options), &errPtr)
	if errPtr != nil {
		return nil, errors.New(C.GoString(errPtr))
	}
	return (*PCREExtra)(extra), nil
}

type Error int

const (
	ErrNoMatch       = C.PCRE_ERROR_NOMATCH
	ErrNull          = C.PCRE_ERROR_NULL
	ErrBadOption     = C.PCRE_ERROR_BADOPTION
	ErrBadMagic      = C.PCRE_ERROR_BADMAGIC
	ErrUnknownOpcode = C.PCRE_ERROR_UNKNOWN_OPCODE
)

/*
PCRE_ERROR_UNKNOWN_NODE
PCRE_ERROR_NOMEMORY
PCRE_ERROR_NOSUBSTRING
PCRE_ERROR_MATCHLIMIT
PCRE_ERROR_CALLOUT
PCRE_ERROR_BADUTF8
PCRE_ERROR_BADUTF16
PCRE_ERROR_BADUTF8_OFFSET
PCRE_ERROR_BADUTF16_OFFSET
PCRE_ERROR_PARTIAL
PCRE_ERROR_BADPARTIAL
PCRE_ERROR_INTERNAL
PCRE_ERROR_BADCOUNT
PCRE_ERROR_DFA_UITEM
PCRE_ERROR_DFA_UCOND
PCRE_ERROR_DFA_UMLIMIT
PCRE_ERROR_DFA_WSSIZE
PCRE_ERROR_DFA_RECURSE
PCRE_ERROR_RECURSIONLIMIT
PCRE_ERROR_NULLWSLIMIT
PCRE_ERROR_BADNEWLINE
PCRE_ERROR_BADOFFSET
PCRE_ERROR_SHORTUTF8
PCRE_ERROR_SHORTUTF16
PCRE_ERROR_RECURSELOOP
PCRE_ERROR_JIT_STACKLIMIT
PCRE_ERROR_BADMODE
PCRE_ERROR_BADENDIANNESS
PCRE_ERROR_DFA_BADRESTART

// Specific error codes for UTF-8 validity checks

#define PCRE_UTF8_ERR0               0
#define PCRE_UTF8_ERR1               1
#define PCRE_UTF8_ERR2               2
#define PCRE_UTF8_ERR3               3
#define PCRE_UTF8_ERR4               4
#define PCRE_UTF8_ERR5               5
#define PCRE_UTF8_ERR6               6
#define PCRE_UTF8_ERR7               7
#define PCRE_UTF8_ERR8               8
#define PCRE_UTF8_ERR9               9
#define PCRE_UTF8_ERR10             10
#define PCRE_UTF8_ERR11             11
#define PCRE_UTF8_ERR12             12
#define PCRE_UTF8_ERR13             13
#define PCRE_UTF8_ERR14             14
#define PCRE_UTF8_ERR15             15
#define PCRE_UTF8_ERR16             16
#define PCRE_UTF8_ERR17             17
#define PCRE_UTF8_ERR18             18
#define PCRE_UTF8_ERR19             19
#define PCRE_UTF8_ERR20             20
#define PCRE_UTF8_ERR21             21

// Specific error codes for UTF-16 validity checks

#define PCRE_UTF16_ERR0              0
#define PCRE_UTF16_ERR1              1
#define PCRE_UTF16_ERR2              2
#define PCRE_UTF16_ERR3              3
#define PCRE_UTF16_ERR4              4

// Request types for pcre_fullinfo()

// Request types for pcre_config(). Do not re-arrange, in order to remain compatible.

#define PCRE_CONFIG_UTF8                    0
#define PCRE_CONFIG_NEWLINE                 1
#define PCRE_CONFIG_LINK_SIZE               2
#define PCRE_CONFIG_POSIX_MALLOC_THRESHOLD  3
#define PCRE_CONFIG_MATCH_LIMIT             4
#define PCRE_CONFIG_STACKRECURSE            5
#define PCRE_CONFIG_UNICODE_PROPERTIES      6
#define PCRE_CONFIG_MATCH_LIMIT_RECURSION   7
#define PCRE_CONFIG_BSR                     8
#define PCRE_CONFIG_JIT                     9
#define PCRE_CONFIG_UTF16                  10
#define PCRE_CONFIG_JITTARGET              11

// Request types for pcre_study(). Do not re-arrange, in order to remain compatible.

#define PCRE_STUDY_JIT_COMPILE                0x0001
#define PCRE_STUDY_JIT_PARTIAL_SOFT_COMPILE   0x0002
#define PCRE_STUDY_JIT_PARTIAL_HARD_COMPILE   0x0004

// Bit flags for the pcre[16]_extra structure. Do not re-arrange or redefine these bits, just add new ones on the end, in order to remain compatible.

#define PCRE_EXTRA_STUDY_DATA             0x0001
#define PCRE_EXTRA_MATCH_LIMIT            0x0002
#define PCRE_EXTRA_CALLOUT_DATA           0x0004
#define PCRE_EXTRA_TABLES                 0x0008
#define PCRE_EXTRA_MATCH_LIMIT_RECURSION  0x0010
#define PCRE_EXTRA_MARK                   0x0020
#define PCRE_EXTRA_EXECUTABLE_JIT         0x0040
*/
