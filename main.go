package main

import (
  "fmt"
  "os"
  "encoding/hex"
  "strconv"
  "strings"
)

func main(){
  hex_input := string(os.Args[1])
  switch prefix := hex_input[0:4]; prefix{
  case "0200":
    //Example: 02000050ACD901030000000000000000 (172.217.1.3:80)
    ip := saddr2IP(hex_input, len(hex_input))
    fmt.Println(ip)
  case "0A00":
    //Example: 0A000050000000002607F8B0400B0800000000000000200300000000 (2607:f8b0:400b:800::2003:80)
    ip := saddr2IP(hex_input, len(hex_input))
    fmt.Println(ip)
  default:
    //Example: 2F7573722F7362696E2F73736864002D44002D52 (/usr/sbin/sshd -D -R)
    output := NullHex2Strings(hex_input)
    fmt.Println(output)
  }
}

func saddr2IP(src string, size int) (string){
  var chunks []int64
  for i := 4; i <= 16; i+=2 {
    n, _ := strconv.ParseInt(src[i:i+2], 16, 64)
    chunks = append(chunks, n)
  }
  port := strconv.FormatInt(int64(256) * chunks[0] + chunks[1], 10)
  var ip string
  if size == int(32){
    //build the IPv4 address string
    for i := 2; i <= 5; i+=1 {
      ip = ip + strconv.FormatInt(chunks[i], 10) + "."
    }
  }
  if size == int(56){
    //ipv6 is already in hex, just get the 8 chunks of 4
    for i := 16; i <= 44; i+=4 {
      ip = ip + src[i:i+4] + ":"
    }
  }
  return ip[:len(ip)-1] + " " + port
}

func NullHex2Strings(b string) (s string) {
  for i:=0; i<=len(b)-2;i+=2{
    if b[i:i+2] == "00"{
      s = s+"20"
    }else{
      s = s+b[i:i+2]
    }
  }
  s = strings.TrimSpace(DecodeHexStrings([]byte(s)))
  return
}

func DecodeHexStrings(src []byte) (string){
  dst := make([]byte, hex.DecodedLen(len(src)))
  dec, _ := hex.Decode(dst, src)

  return string(dst[:dec])
}