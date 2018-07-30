# [Set 1](https://cryptopals.com/sets/1) Learnings

## Challenge 1: Convert hex to base64

The rule we were given for this set was:
> Always operate on raw bytes, never on encoded strings. Only use hex and base64 for pretty-printing.

The rule for this set is important. We want to explore the difference between slices of bytes and encoded strings. Unicode code point strings are different from bytes. Everything is a byte slice [in Go](https://blog.golang.org/strings):

> In Go, a string is in effect a read-only slice of bytes [...] It's important to state right up front that a string holds arbitrary bytes. It is not required to hold Unicode text, UTF-8 text, or any other predefined format. As far as the content of a string is concerned, it is exactly equivalent to a slice of bytes.

I wrote [HexToBase64](./HexToBase64/HexToBase64.go) to solve this problem. In this method, we first [decode the string](https://golang.org/pkg/encoding/hex/#DecodeString) which returns the bytes represented by the hexadecimal string in case it is malformed. Then we use a [base64 encoding method](https://golang.org/pkg/encoding/base64/#Encoding.EncodeToString) to return a base64 encoding of the input.

## Fixed XOR

## Single-byte XOR cipher

## Detect single-character XOR

## Implement repeating-key XOR

## Break repeating-key XOR

## AES in ECB mode

## Detect AES in ECB mode
