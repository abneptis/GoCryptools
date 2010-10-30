package signer

import "bytes"
import "encoding/base64"
import "os"

type Signer interface {
  Sign([]byte)([]byte, os.Error)
  Verify([]byte, []byte)(os.Error)
}

var SignatureVerificationFailed = os.NewError("Signature Verification Failed")

// Sign a string with a specified signer and base64 encoding
func Sign64(s Signer, e *base64.Encoding, sts []byte)(out []byte, err os.Error){
  sig, err := s.Sign(sts)
  if err != nil { return }
  out = make([]byte, e.EncodedLen(len(sig)))
  e.Encode(out, sig)
  return
}

func SignString(s Signer, n string)(out string, err os.Error){
  bb := bytes.NewBufferString(n)
  bo, err := s.Sign(bb.Bytes())
  if err == nil { out = string(bo) }
  return
}

func SignString64(s Signer, e *base64.Encoding, sts string)(out string, err os.Error){
  bb := bytes.NewBufferString(sts)
  bo, err := Sign64(s, e, bb.Bytes())
  if err == nil { out = string(bo) }
  return
}
