type ElectionProof struct {
    // The VRFProof is generated by running our VRF on a past ticket in the ticket chain
    // signed with the miner's keypair. This field is 97 bytes long (may be compressible
    // to 80).
    VRFProof Bytes
}