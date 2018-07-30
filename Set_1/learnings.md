# [Set 1](https://cryptopals.com/sets/1) Learnings

## [Challenge 1: Convert hex to base64](https://cryptopals.com/sets/1/challenges/1)

The rule we were given for this set was:
> Always operate on raw bytes, never on encoded strings. Only use hex and base64 for pretty-printing.

The rule for this set is important. We want to explore the difference between slices of bytes and encoded strings. Unicode code point strings are different from bytes. Everything is a byte slice [in Go](https://blog.golang.org/strings):

> In Go, a string is in effect a read-only slice of bytes [...] It's important to state right up front that a string holds arbitrary bytes. It is not required to hold Unicode text, UTF-8 text, or any other predefined format. As far as the content of a string is concerned, it is exactly equivalent to a slice of bytes.

I wrote [HexToBase64](./HexToBase64/HexToBase64.go) to solve this problem. In this method, we first [decode the string](https://golang.org/pkg/encoding/hex/#DecodeString) which returns the bytes represented by the hexadecimal string in case it is malformed. Then we use a [base64 encoding method](https://golang.org/pkg/encoding/base64/#Encoding.EncodeToString) to return a base64 encoding of the input.

Secret messages?

> I'm killing your brain like a poisonous mushroom

## [Challenge 2: Fixed XOR](https://cryptopals.com/sets/1/challenges/2)

Here we are asked to take two equal-length buffers and produce their XOR combination. An XOR operation is a logical operation that outputs "true" or the value for true only when inputs differ. So we are given two values and asked to find what the xor produces.

I wrote [FixedXOR](./FixedXOR/FixedXOR.go) to do this translation. It's a pretty normal XOR function where it takes the length and makes sure that is the same then moved through each byte and compares them.

At one point I got an error that you cannot compare slices to nil and I had to use `bytes.Equal` for comparison, which I did not know about!

Secret messages?

> KSSP // [this ?](https://kernsec.org/wiki/index.php/Kernel_Self_Protection_Project)
> hit the bull's eye

## [Challenge 3: Single-byte XOR cipher](https://cryptopals.com/sets/1/challenges/3)

We are told that a hex encoded string we are given has been XOR'd against a single character. We need to find the key to decrypt the message, but it doesn't want us to use code. We need a method for "scoring" a piece of English plaintext. So... we need a bunch of English! I went to Project Gutenberg and pulled "Tale of Two Cities" for some random English.

So in our method we get some random english from our Gutenberg file. Then we gather that English into some data in the form of a map with rune's as "keys" and a score as values. This method for gathering English data takes a string of the ToTC book (or any string), for each character in the data it adds that character as a rune and gives it a `count`. Then, it takes the total characters in the book and goes through again, making the values a percentage (the `count` value divided by the total value for that character). After we have our English map, we pass our message string given to us by the problem, and the map, into a method for calculating the single byte XOR. This will go through, create an output for every possible key (out of 256 bytes), find the singleXor comparison and then find the "score" of that against our map. To find that score, we iterate through the outputted string and add the "score" for each of those letters to a total and then average it at the end. We then choose the "max" for these scores.

Doing this exercise I learned maps have a default zero value! I also learned about runes. Rune literals are integer values "mapped" to their unicode codepoint.

Secret message?

> Cooking MC's like a pound of bacon

## [Challenge 4: Detect single-character XOR](https://cryptopals.com/sets/1/challenges/4)

## [Challenge 5: Implement repeating-key XOR](https://cryptopals.com/sets/1/challenges/5)

## [Challenge 6: Break repeating-key XOR](https://cryptopals.com/sets/1/challenges/6)

## [Challenge 7: AES in ECB mode](https://cryptopals.com/sets/1/challenges/7)

## [Challenge 8: Detect AES in ECB mode](https://cryptopals.com/sets/1/challenges/8)
