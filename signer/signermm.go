// Interface and helper functions for (hmac) signatures
//
// Signer creates a Sign() and Verify() function for cryptographic
// APIs

package signer

/* 
  Copyright (c) 2010, Abneptis LLC.
  See COPYRIGHT and LICENSE for details.
*/

import "bytes"
import "encoding/base64"
import "os"

type multiMech interface {
  KnownMechs()([]string)
}

func ValidMech(m string, mm multiMech)(err os.Error){
  err = os.NewError("Unknown mech")
  mechs := mm.KnownMechs()
  for i := range(mechs){
    if mechs[i] == m { err = nil }
  }
  return
}

type VerifierMultiMech interface {
  multiMech
  SignerID()(string)
  Verify(string, []byte, []byte)(os.Error)
}

type SignerMultiMech interface {
  VerifierMultiMech
  Sign(string, []byte)([]byte, os.Error)
}

func Sign64Mech(mech string, s SignerMultiMech, e *base64.Encoding, sts []byte)(out []byte, err os.Error){
  sig, err := s.Sign(mech, sts)
  if err != nil { return }
  out = make([]byte, e.EncodedLen(len(sig)))
  e.Encode(out, sig)
  return
}

// Executes Sign(), however uses strings rather than the native
// []byte types.
func SignStringMech(mech string, s SignerMultiMech, n string)(out string, err os.Error){
  bb := bytes.NewBufferString(n)
  bo, err := s.Sign(mech, bb.Bytes())
  if err == nil { out = string(bo) }
  return
}

// Return a signature encoded in base64 (with the encoding
// specified by the caller)
func SignString64Mech(mech string, s SignerMultiMech, e *base64.Encoding, sts string)(out string, err os.Error){
  bb := bytes.NewBufferString(sts)
  bo, err := Sign64Mech(mech, s, e, bb.Bytes())
  if err == nil { out = string(bo) }
  return
}

