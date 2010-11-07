package hextools

import "testing"
import "crypto/md5"
import "bytes" 

func TestZeroNibbles(t *testing.T){
  low, high := ByteNibbles(0)
  if low != 0 {
    t.Errorf("%s-order zeroed-nibble test failed: (exp: %d; act: %d)", "low", 0, low)
  }
  if high != 0 {
    t.Errorf("%s-order zeroed-nibble test failed: (exp: %d; act: %d)", "high", 0, high)
  }
}

func testNibble(t *testing.T, test string, exp, act byte){
  if exp != act {
    t.Errorf("(%s): testNibble(E:0x%X/A:0x%X)", test, exp, act)
  }
}

func TestLowNibble(t *testing.T) {
  low, high := ByteNibbles(0xdf)
  testNibble(t, "TestLowNibble(low)", 0x0f, low)
  testNibble(t, "TestLowNibble(high)", 0x0d, high)
}

func TestHighNibble(t *testing.T) {
  low, high := ByteNibbles(0xfd)
  testNibble(t, "TestHighNibble(low)", 0x0d, low)
  testNibble(t, "TestHighNibble(high)", 0x0f, high)
}

func TestOneHigh(t *testing.T){
  in := []byte{0xf4}
  ob, err := HexBytes(true, true, in)
  if err != nil {
    t.Errorf("Single-byte high-hexing failed: (%v)", err)
  }
  testNibble(t, "TestOneHigh(high)", 'F', ob[0])
  testNibble(t, "TestOneHigh(low)", '4', ob[1])
}

func TestOneLow(t *testing.T){
  in := []byte{0xf4}
  ob, err := HexBytes(true, false, in)
  if err != nil {
    t.Errorf("Single-byte high-hexing failed: (%v)", err)
  }
  testNibble(t, "TestOneLow(high)", 'F', ob[1])
  testNibble(t, "TestOneLow(low)", '4', ob[0])
}

func TestHexString(t *testing.T){
  tstr := "hello world"
  ehex := bytes.NewBufferString("5eb63bbbe01eeed093cb22bb8f5acdc3")
  m := md5.New()
  bb := bytes.NewBufferString(tstr)
  m.Write(bb.Bytes())
  rhex, _ := HexString(false, true, string(m.Sum()))
  if ehex.Len() != len(rhex) {
    t.Errorf("Incorect MD5 result length: [E: %v; A: %v]", ehex.Bytes(), m.Sum())
  }
  for i := range(ehex.Bytes()) {
    if ehex.Bytes()[i] != rhex[i] {
      t.Errorf("Incorect MD5 result [%d]: [E: %v; A: %v]", i, ehex.Bytes(), m.Sum())
    }
  }
}

func TestHexStringUp(t *testing.T){
  tstr := "hello world"
  ehex := bytes.NewBufferString("5EB63BBBE01EEED093CB22BB8F5ACDC3")
  m := md5.New()
  bb := bytes.NewBufferString(tstr)
  m.Write(bb.Bytes())
  rhex, _ := HexString(true, true, string(m.Sum()))
  if ehex.Len() != len(rhex) {
    t.Errorf("Incorect MD5 result length: [E: %v; A: %v]", ehex.Bytes(), m.Sum())
  }
  for i := range(ehex.Bytes()) {
    if ehex.Bytes()[i] != rhex[i] {
      t.Errorf("Incorect MD5 result [%d]: [E: %v; A: %v]", i, ehex.Bytes(), m.Sum())
    }
  }
}
