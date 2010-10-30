package hashes

/* 
  Copyright (c) 2010, Abneptis LLC.
  See COPYRIGHT and LICENSE for details.
*/

import "crypto/hmac"
import "crypto/md4"
import "crypto/md5"
import "crypto/sha1"
import "crypto/sha256"
import "crypto/sha512"

import "hash"
import "strings"
import "os"

var Md4  = func()(hash.Hash){ return md4.New() }
var Md5  = func()(hash.Hash){ return md5.New() }
var Sha1 = func()(hash.Hash){ return sha1.New() }
var Sha256 = func()(hash.Hash){ return sha256.New() }
var Sha384 = func()(hash.Hash){ return sha512.New384() }
var Sha512 = func()(hash.Hash){ return sha512.New() }
var ErrUnknownHash = os.NewError("Unknown Hash")

func GetHashFunc(n string)(hf func()(hash.Hash), err os.Error){
  switch strings.ToLower(n) {
    case "md4": hf =  Md4
    case "md5": hf =  Md4
    case "sha1": hf =  Sha1
    case "sha256": hf =  Sha256
    case "sha384": hf =  Sha384
    case "sha512": hf =  Sha512
    default: err = ErrUnknownHash
  }
  return
}

func GetHash(n string)(h hash.Hash, err os.Error){
  f, err := GetHashFunc(n)
  if err != nil { return }
  h = f()
  return
}

func GetHmac(n string, key []byte)(h hash.Hash, err os.Error){
  hf, err := GetHashFunc(n)
  if err != nil { return }
  h = hmac.New(hf, key)
  return
}
