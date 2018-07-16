package main

import (
	"C"

	"flag"
	"unsafe"

	"github.com/perlin-network/noise/crypto"
	"github.com/perlin-network/noise/crypto/ed25519"
)

//export Noise_Release
func Noise_Release(p unsafe.Pointer) {
	Resources.Delete(UnmarshalPointer(p))
}

//export Noise_SetLogToStderr
func Noise_SetLogToStderr(val uint32) {
	if val != 0 {
		flag.Set("logtostderr", "true")
	} else {
		flag.Set("logtostderr", "false")
	}
}

//export Noise_NewKeyPair
func Noise_NewKeyPair(publicKey []byte, privateKey []byte) unsafe.Pointer {
	kp := &crypto.KeyPair {
		PublicKey: make([]byte, len(publicKey)),
		PrivateKey: make([]byte, len(privateKey)),
	}
	copy(kp.PublicKey, publicKey)
	copy(kp.PrivateKey, privateKey)

	ret := unsafe.Pointer(kp)
	Resources.Add(ret)
	return MarshalPointer(ret)
}

//export Noise_RandomKeyPair
func Noise_RandomKeyPair(provider string) unsafe.Pointer {
	var kp *crypto.KeyPair

	switch provider {
	case "ed25519":
		kp = ed25519.RandomKeyPair()
	}

	if kp != nil {
		ret := unsafe.Pointer(kp)
		Resources.Add(ret)
		return MarshalPointer(ret)
	} else {
		return nil
	}
}

func main() {}
