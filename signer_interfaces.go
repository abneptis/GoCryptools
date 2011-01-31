package cryptools

import "hash"
import "os"

// A Signer will produce a Signature.
// Generally, developers should not implement signer,
// but instead use a NewSigner(HashSigner(alg, key)).
type Signer interface {
  Sign(Signable)(Signature, os.Error)
}

// A HashingSigner produces a signature based on the
// provided hashing algorithm and Signable data.
type HashingSigner interface {
  Sign(hash.Hash, Signable)(Signature, os.Error)
}

// A named signer has a canonical name that is used for
// canonicalization/signing purposes. 
type NamedSigner interface {
  Signer
  SignerName()(string)
}

type NamedHashingSigner interface {
  HashingSigner
  SignerName()(string)
}

