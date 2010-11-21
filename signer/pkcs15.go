package signer

import "com.abneptis.oss/cryptools/hashes"

import "crypto/rand"
import "crypto/rsa"
import "os"

type PKCS15Signer struct {
  PrivateKey *rsa.PrivateKey
  ID string
}

func NewPKCS15Signer(id string, pk *rsa.PrivateKey)(SignerMultiMech){
  return PKCS15Signer{PrivateKey: pk, ID: id}
}

func (self PKCS15Signer)KnownMechs()([]string){
  return []string{
    "md5","sha1","sha256","sha384","sha512","md5sha1",
  }
}

func (self PKCS15Signer)SignerID()(string){
  return self.ID
}

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
  err = rsa.VerifyPKCS1v15(&self.PrivateKey.PublicKey, hi, hbytes, sig)
  return
}

