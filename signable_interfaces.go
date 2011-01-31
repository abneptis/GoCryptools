package cryptools

import "os"

type Signable interface {
  SignableBytes()([]byte, os.Error)
}
