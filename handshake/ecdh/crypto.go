package ecdh

import (
	"crypto/sha512"

	"github.com/perlin-network/noise/internal/edwards25519"
)

func computeSharedKey(nodePrivateKey edwards25519.PrivateKey, remotePublicKey edwards25519.PublicKey) []byte {
	var nodeSecretKeyBuf, remotePublicKeyBuf, sharedKeyBuf [32]byte
	copy(nodeSecretKeyBuf[:], deriveSecretKey(nodePrivateKey))
	copy(remotePublicKeyBuf[:], remotePublicKey[:])

	var sharedKeyElement, publicKeyElement edwards25519.ExtendedGroupElement
	publicKeyElement.FromBytes(&remotePublicKeyBuf)

	edwards25519.GeScalarMult(&sharedKeyElement, &nodeSecretKeyBuf, &publicKeyElement)

	sharedKeyElement.ToBytes(&sharedKeyBuf)

	return sharedKeyBuf[:]
}

func deriveSecretKey(privateKey edwards25519.PrivateKey) []byte {
	digest := sha512.Sum512(privateKey[:32])
	digest[0] &= 248
	digest[31] &= 127
	digest[31] |= 64

	return digest[:32]
}

func isEd25519GroupElement(buf []byte) bool {
	if len(buf) != edwards25519.PublicKeySize {
		return false
	}

	var buff [32]byte
	copy(buff[:], buf)

	var A edwards25519.ExtendedGroupElement
	return A.FromBytes(&buff)
}
