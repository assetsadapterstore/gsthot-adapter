package addrdec

import (
	"fmt"
	"strings"

	"github.com/blocktree/go-owcdrivers/addressEncoder"
)

var (
	PublicKeyPrefix       = "PUB_"
	PublicKeyK1Prefix     = "PUB_K1_"
	PublicKeyR1Prefix     = "PUB_R1_"
	PublicKeyPrefixCompat = "GST"

	MainnetPublic = addressEncoder.AddressType{"gst", addressEncoder.BTCAlphabet, "ripemd160", "", 33, []byte(PublicKeyPrefixCompat), nil}

	Default = AddressDecoderV2{}
)

//AddressDecoderV2
type AddressDecoderV2 struct {
	IsTestNet bool
}

// AddressDecode decode address
func (dec *AddressDecoderV2) AddressDecode(pubKey string) ([]byte, error) {

	var pubKeyMaterial string
	if strings.HasPrefix(pubKey, PublicKeyR1Prefix) {
		pubKeyMaterial = pubKey[len(PublicKeyR1Prefix):] // strip "PUB_R1_"
	} else if strings.HasPrefix(pubKey, PublicKeyK1Prefix) {
		pubKeyMaterial = pubKey[len(PublicKeyK1Prefix):] // strip "PUB_K1_"
	} else if strings.HasPrefix(pubKey, PublicKeyPrefixCompat) { // "GST"
		pubKeyMaterial = pubKey[len(PublicKeyPrefixCompat):] // strip "GST"
	} else {
		return nil, fmt.Errorf("public key should start with [%q | %q] (or the old %q)", PublicKeyK1Prefix, PublicKeyR1Prefix, PublicKeyPrefixCompat)
	}

	ret, err := addressEncoder.Base58Decode(pubKeyMaterial, addressEncoder.NewBase58Alphabet(MainnetPublic.Alphabet))
	if err != nil {
		return nil, addressEncoder.ErrorInvalidAddress
	}
	if addressEncoder.VerifyChecksum(ret, MainnetPublic.ChecksumType) == false {
		return nil, addressEncoder.ErrorInvalidAddress
	}

	return ret[:len(ret)-4], nil
}

// AddressEncode encode address
func (dec *AddressDecoderV2) AddressEncode(hash []byte) string {
	data := addressEncoder.CatData(hash, addressEncoder.CalcChecksum(hash, MainnetPublic.ChecksumType))
	return string(MainnetPublic.Prefix) + addressEncoder.EncodeData(data, "base58", MainnetPublic.Alphabet)
}
