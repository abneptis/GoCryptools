// Interface and helper functions for (hmac) signatures
//
// Signer creates a Sign() and Verify() function for cryptographic
// APIs
//
// NB: Verify may be removed from the Signer interface
// and replaced with an module function in the near
// future.
package cryptools

/* 
  Copyright (c) 2010, Abneptis LLC.
  See COPYRIGHT and LICENSE for details.
*/


import "encoding/base64"
import "os"

var SignatureVerificationFailed = os.NewError("Signature Verification Failed")

// Sign a string with a specified signer and base64 encoding
func Sign64(s Signer, e *base64.Encoding,
            ss Signable)(out []byte, err os.Error){

  sig, err := s.Sign(ss)
  if err != nil { return }
  bb := sig.SignatureBytes()
  out = make([]byte, e.EncodedLen(len(bb)))
  e.Encode(out, bb)
  return
}

// Executes Sign(), however uses strings rather than the native
// []byte types.
func SignString(s Signer, n string)(out string, err os.Error){
  bo, err := s.Sign(SignableString(n))
  if err == nil { out = string(bo.SignatureBytes()) }
  return
}

// Return a signature encoded in base64 (with the encoding
// specified by the caller)
func SignString64(s Signer, e *base64.Encoding,
                  so Signable)(out string, err os.Error){

  bo, err := Sign64(s, e, so)
  if err == nil { out = string(bo) }
  return
}
