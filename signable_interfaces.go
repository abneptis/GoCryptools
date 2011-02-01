package cryptools

import "bytes"
import "os"

type Signable interface {
  SignableBytes()([]byte, os.Error)
}

func SignableString(s string)(Signable){
  return signableString(s)
}

type signableString string

func (self signableString)SignableBytes()(out []byte, err os.Error){
  out = bytes.NewBufferString(string(self)).Bytes()
  return
}
