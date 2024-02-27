package bip32

import "github.com/pkg/errors"

// BitcoinMainnetPrivate is the version that is used for
// bitcoin mainnet bip32 private extended keys.
// Ecnodes to xprv in base58.
var BitcoinMainnetPrivate = [4]byte{
	0x04,
	0x88,
	0xad,
	0xe4,
}

// BitcoinMainnetPublic is the version that is used for
// bitcoin mainnet bip32 public extended keys.
// Ecnodes to xpub in base58.
var BitcoinMainnetPublic = [4]byte{
	0x04,
	0x88,
	0xb2,
	0x1e,
}

// SeraMainnetPrivate is the version that is used for
// sera mainnet bip32 private extended keys.
// Ecnodes to xprv in base58.
var SeraMainnetPrivate = [4]byte{
	0x03,
	0x8f,
	0x2e,
	0xf4,
}

// SeraMainnetPublic is the version that is used for
// sera mainnet bip32 public extended keys.
// Ecnodes to kpub in base58.
var SeraMainnetPublic = [4]byte{
	0x03,
	0x8f,
	0x33,
	0x2e,
}

// SeraTestnetPrivate is the version that is used for
// sera testnet bip32 public extended keys.
// Ecnodes to ktrv in base58.
var SeraTestnetPrivate = [4]byte{
	0x03,
	0x90,
	0x9e,
	0x07,
}

// SeraTestnetPublic is the version that is used for
// sera testnet bip32 public extended keys.
// Ecnodes to ktub in base58.
var SeraTestnetPublic = [4]byte{
	0x03,
	0x90,
	0xa2,
	0x41,
}

// SeraDevnetPrivate is the version that is used for
// sera devnet bip32 public extended keys.
// Ecnodes to kdrv in base58.
var SeraDevnetPrivate = [4]byte{
	0x03,
	0x8b,
	0x3d,
	0x80,
}

// SeraDevnetPublic is the version that is used for
// sera devnet bip32 public extended keys.
// Ecnodes to xdub in base58.
var SeraDevnetPublic = [4]byte{
	0x03,
	0x8b,
	0x41,
	0xba,
}

// SeraSimnetPrivate is the version that is used for
// sera simnet bip32 public extended keys.
// Ecnodes to ksrv in base58.
var SeraSimnetPrivate = [4]byte{
	0x03,
	0x90,
	0x42,
	0x42,
}

// SeraSimnetPublic is the version that is used for
// sera simnet bip32 public extended keys.
// Ecnodes to xsub in base58.
var SeraSimnetPublic = [4]byte{
	0x03,
	0x90,
	0x46,
	0x7d,
}

func toPublicVersion(version [4]byte) ([4]byte, error) {
	switch version {
	case BitcoinMainnetPrivate:
		return BitcoinMainnetPublic, nil
	case SeraMainnetPrivate:
		return SeraMainnetPublic, nil
	case SeraTestnetPrivate:
		return SeraTestnetPublic, nil
	case SeraDevnetPrivate:
		return SeraDevnetPublic, nil
	case SeraSimnetPrivate:
		return SeraSimnetPublic, nil
	}

	return [4]byte{}, errors.Errorf("unknown version %x", version)
}

func isPrivateVersion(version [4]byte) bool {
	switch version {
	case BitcoinMainnetPrivate:
		return true
	case SeraMainnetPrivate:
		return true
	case SeraTestnetPrivate:
		return true
	case SeraDevnetPrivate:
		return true
	case SeraSimnetPrivate:
		return true
	}

	return false
}
