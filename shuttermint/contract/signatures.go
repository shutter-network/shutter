package contract

// SignatureToContractFormat converts a signature created with goethereum's crypto.Sign to the
// format understood by the contracts. I.e., it updates the v value from 0/1 to 27/28.
func SignatureToContractFormat(sig []byte) []byte {
	c := make([]byte, len(sig))
	copy(c, sig)
	c[64] += 27
	return c
}

// SignaturesToContractFormat converts a list of signatures created with goethereum's crypto.Sign
// to the format understood by the contracts. I.e., it updates the v values from 0/1 to 27/28.
func SignaturesToContractFormat(sigs [][]byte) [][]byte {
	cs := [][]byte{}
	for _, s := range sigs {
		cs = append(cs, SignatureToContractFormat(s))
	}
	return cs
}
