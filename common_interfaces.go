package cryptools

// Note that a signatures Bytes() cannot fail, thus
// an invalid Signature shouldn't be instantiated
type Signature interface {
  SignatureBytes()([]byte)
}
