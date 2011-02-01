package cryptools

import "os"

type Verifier interface {
  VerifySignature(Signature)(os.Error)
}
