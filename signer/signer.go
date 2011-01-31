// Interface and helper functions for (hmac) signatures
//
// Signer creates a Sign() and Verify() function for cryptographic
// APIs
//
// NB: Verify may be removed from the Signer interface
// and replaced with an module function in the near
// future.
package signer

/* 
  Copyright (c) 2010, Abneptis LLC.
  See COPYRIGHT and LICENSE for details.
*/


import "com.abneptis.oss/cryptools"

import "encoding/base64"
import "os"

/* A simple interface for signing and verifying HMAC signatures */
// DEPRECATED: Use cryptools.Signer
type Signer interface {
  cryptools.Signer
}

// DEPRECATED: Use cryptools.NamedSigner
type NamedSigner interface {
  cryptools.NamedSigner
}

// DEPRECATED: Use crptools.Verifier
type Verifier interface {
  cryptools.Verifier
}

var SignatureVerificationFailed = os.NewError("Signature Verification Failed")

// Sign a string with a specified signer and base64 encoding
func Sign64(s Signer, e *base64.Encoding,
            ss cryptools.Signable)(out []byte, err os.Error){

  sig, err := s.Sign(ss)
  if err != nil { return }
  bb := sig.Bytes()
  out = make([]byte, e.EncodedLen(len(bb)))
  e.Encode(out, bb)
  return
}

// Executes Sign(), however uses strings rather than the native
// []byte types.
func SignString(s Signer, n string)(out string, err os.Error){
  bo, err := s.Sign(SignableString(n))
  if err == nil { out = string(bo.Bytes()) }
  return
}

// Return a signature encoded in base64 (with the encoding
// specified by the caller)
func SignString64(s Signer, e *base64.Encoding,
                  so cryptools.Signable)(out string, err os.Error){

  bo, err := Sign64(s, e, so)
  if err == nil { out = string(bo) }
  return
}
