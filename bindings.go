package pedersen_hash

import (
        "unsafe"
)

// #cgo LDFLAGS: ${SRCDIR}/target/release/libpedersen_hash_ffi.a -lutil -lutil -ldl -lrt -lpthread -lgcc_s -lc -lm -lrt -lpthread -lutil -lutil
// #include "./pedersen_hash_ffi.h"
import "C"

const outputHashLength = 32

func goBytes(src *C.uint8_t, size C.size_t) []byte {
        return C.GoBytes(unsafe.Pointer(src), C.int(size))
}

func PedersenHashFFI(preimage []byte) [32]byte {

        preimageCBytes := C.CBytes(preimage[:])
	defer C.free(preimageCBytes)

	hash := C.pedersen_hash_ffi(
		(*C.uint8_t)(preimageCBytes),
		C.size_t(len(preimage)),
	)

        hashBytes := goBytes(&hash.data[0], outputHashLength)

        var out [outputHashLength]byte
        copy(out[:], hashBytes)
        return out
}
