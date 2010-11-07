package hextools

import "bytes"
import "os"

var lowHex = []byte{'0','1','2','3','4','5','6','7','8','9',
                    'a','b','c','d','e','f'}
var upHex  = []byte{'0','1','2','3','4','5','6','7','8','9',
                    'A','B','C','D','E','F'}

// Split a byte into two nibbles (represented as 4-bit, low-order bytes)
func ByteNibbles(b byte)(low, high byte){
  low = b & 0x0f
  high = b & 0xf0 >> 4
  return
}

// Returns the hexedecimal digit of the low-order nibble of byte (b).
// case is determined by the upcase bool.
func HexNibble(upcase bool, b byte)(o byte){
  if upcase {
   o = upHex[b & 0x0f]
  } else {
   o = lowHex[b & 0x0f]
  }
  return
}

// Hexify a single byte (High, Low)
func HexByteHigh(upcase bool, b byte)(out []byte){
  low, high := ByteNibbles(b)
  out = make([]byte, 2)
  out[0], out[1] = HexNibble(upcase, high), HexNibble(upcase, low)
  return
}

// Hexify a single byte (Low, High)
func HexByteLow(upcase bool, b byte)(out []byte){
  low, high := ByteNibbles(b)
  out = make([]byte, 2)
  out[0], out[1] = HexNibble(upcase, low), HexNibble(upcase, high)
  return
}

// Wrapper to HexByte{Low:High}
// BROKEN
func HexByte(upcase, high bool, b byte)(out []byte){
  if high {
   out = HexByteHigh(upcase, b)
  } else {
   out = HexByteLow(upcase, b)
  }
  return
}

// Return an hex-string byte array with upper and order specified
// by the caller.
func HexBytes(upper bool, high bool, in []byte)(out []byte, err os.Error){
  out = make([]byte, len(in) *2)
  for ini := range(in) {
    copy(out[ini*2:ini*2+2], HexByte(upper, high, in[ini]))
  }
  return
}

// Wrapper on HexBytes that receives (and returns) a string type.
func HexString(upper bool, high bool, in string)(out string, err os.Error){
  bb := bytes.NewBufferString(in)
  ob, err := HexBytes(upper, high, bb.Bytes())
  if err == nil {
    out = string(ob)
  }
  return
}
