use std::slice::from_raw_parts;

use filcrypto_helpers::pedersen_hash;

#[repr(C)]
#[derive(Clone, Copy, Default)]
pub struct PedersenHash {
    pub data: [u8; 32],
}

// Caller is responsible for de-allocating the preimage input (if necessary).
#[no_mangle]
pub unsafe extern "C" fn pedersen_hash_ffi(preimage: *const u8, preimage_len: usize) -> PedersenHash {
    let input = from_raw_parts(preimage, preimage_len);

    // Sanity check
    assert_eq!(input.len(), preimage_len);

    PedersenHash {
        data: pedersen_hash(input),
    }
}
    
