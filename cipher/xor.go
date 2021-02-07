package cipher

// Xor Cipher (Exclusive-OR encryption)
func Xor(line string) (output string) {
	key := "K"

	for i := 0; i < len(line); i++ {
		output += string(line[i] ^ key[i%len(key)])
	}
	return
}
