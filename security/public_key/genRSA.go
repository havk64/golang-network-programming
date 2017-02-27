package main

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/gob"
	"encoding/pem"
	"fmt"
	"os"
)

// rsaStringer is used to allow the implementation of a Stringer to
// rsa.PrivateKey that is the generated key.
type rsaStringer struct {
	*rsa.PrivateKey
}

func main() {
	// Random number generator posteriorly used to generate the key
	reader := rand.Reader
	bitSize := 512 // key size
	key, err := rsa.GenerateKey(reader, bitSize)
	checkError(err)

	r := rsaStringer{key}
	fmt.Printf("%s\n", r) // Prints the rsa key struct

	publicKey := key.PublicKey

	saveGobKey("private.key", key)
	saveGobKey("public.key", publicKey)

	savePEMKey("private.pem", key)
}

func saveGobKey(fileName string, key interface{}) {
	outFile, err := os.Create(fileName)
	checkError(err)

	encoder := gob.NewEncoder(outFile)
	err = encoder.Encode(key)
	checkError(err)
	outFile.Close()
}

func savePEMKey(fileName string, key *rsa.PrivateKey) {

	outFile, err := os.Create(fileName)
	checkError(err)

	var privateKey = &pem.Block{Type: "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key)}

	pem.Encode(outFile, privateKey)

	outFile.Close()
}

// String method implements stringer interface on our custom type in order to
// represent the struct as a formatted string
func (r rsaStringer) String() string {
	var stream bytes.Buffer
	fmt.Fprintf(&stream, "rsa.PrivateKey{\n")
	fmt.Fprintf(&stream, "    PublicKey:\n")
	fmt.Fprintf(&stream, "        rsa.PublicKey:{\n")
	fmt.Fprintf(&stream, "            N: %v,\n", r.PublicKey.N)
	fmt.Fprintf(&stream, "            E: %v\n", r.PublicKey.E)
	fmt.Fprintf(&stream, "        },\n")
	fmt.Fprintf(&stream, "    D: %v,\n", r.D)
	fmt.Fprintf(&stream, "    Primes: []*big.Int[{\n")
	fmt.Fprintf(&stream, "       %v,\n", r.Primes[0])
	fmt.Fprintf(&stream, "       %v\n", r.Primes[1])
	fmt.Fprintf(&stream, "    }],\n")
	fmt.Fprintf(&stream, "    Precomputed:\n")
	fmt.Fprintf(&stream, "        rsa.PrecomputedValues{\n")
	fmt.Fprintf(&stream, "            Dp: %v,\n", r.Precomputed.Dp)
	fmt.Fprintf(&stream, "            Dq: %v,\n", r.Precomputed.Dq)
	fmt.Fprintf(&stream, "            Qinv: %v,\n", r.Precomputed.Qinv)
	fmt.Fprintf(&stream, "            CRTValues: %#v\n", r.Precomputed.CRTValues)
	fmt.Fprintf(&stream, "        }\n")
	fmt.Fprintf(&stream, "}\n")

	return stream.String()
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
