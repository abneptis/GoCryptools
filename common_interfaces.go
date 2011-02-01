package cryptools

// Note that a signatures Bytes() cannot fail, thus
// an invalid Signature shouldn't be instantiated
type Signature interface {
  SignatureBytes()([]byte)
}

func NewSignature(b []byte)(Signature){
  return signature(b)
}

type signature []byte

func (self signature)SignatureBytes()([]byte){
  return []byte(self)
}
