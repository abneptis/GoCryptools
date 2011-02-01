package cryptools

import "com.abneptis.oss/cryptools/hashes"

import "crypto/rand"
import "crypto/rsa"
import "os"

func PKCS15KnownMechs()([]string){
  return []string{
    "md5","sha1","sha256","sha384","sha512","md5sha1",
  }
}

type PKCS15Signer struct {
  PrivateKey *rsa.PrivateKey
  ID string
}

type PKCS15Verifier struct {
  PublicKey *rsa.PublicKey
  ID string
}

func NewPKCS15Signer(id string, pk *rsa.PrivateKey)(SignerMultiMech){
  return PKCS15Signer{PrivateKey: pk, ID: id}
}

func NewPKCS15Verifier(id string, pk *rsa.PublicKey)(VerifierMultiMech){
  return PKCS15Verifier{PublicKey: pk, ID: id}
}

func (self PKCS15Signer)KnownMechs()([]string){ return PKCS15KnownMechs() }
func (self PKCS15Verifier)KnownMechs()([]string){ return PKCS15KnownMechs() }
func (self PKCS15Signer)SignerID()(string){ return self.ID }
func (self PKCS15Verifier)SignerID()(string){ return self.ID }

func (self PKCS15Signer)Sign(mech string, in []byte)(out []byte, err os.Error){
  err = ValidMech(mech, self)
  if err != nil { return }
  h, err := hashes.GetHash(mech)
  if err != nil { return }
  _, err = h.Write(in)
  if err != nil { return }
  hbytes := h.Sum()
  //log.Printf("Sign Hash: %X", hbytes)
  hi, err := hashes.GetPKCS15Hash(mech)
  if err != nil { return }
  out, err = rsa.SignPKCS1v15(rand.Reader, self.PrivateKey, hi, hbytes)
  return
}

func (self PKCS15Signer)Verify(mech string, sig []byte, sigo []byte)(err os.Error){
  vo := NewPKCS15Verifier(self.ID, &self.PrivateKey.PublicKey)
  return vo.Verify(mech,sig,sigo)
}

func (self PKCS15Verifier)Verify(mech string, sig []byte, sigo []byte)(err os.Error){
  err = ValidMech(mech, self)
  if err != nil { return }
  h, err := hashes.GetHash(mech)
  if err != nil { return }
  _, err = h.Write(sigo)
  if err != nil { return }
  hbytes := h.Sum()
  //log.Printf("Verify Hash: %X", hbytes)
  hi, err := hashes.GetPKCS15Hash(mech)
  if err != nil { return }
  err = rsa.VerifyPKCS1v15(self.PublicKey, hi, hbytes, sig)
  return
}

