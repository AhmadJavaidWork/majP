package encrypt

func Encrypt(target, password string) string {
	res := []rune(target)
	for i := 0; i < len(res); i++ {
		key := int(rune(password[i%len(password)]))
		res[i] = rune(int(res[i]) + key)
	}
	return string(res)
}

func Decrypt(target, password string) string {
	res := []rune(target)
	for i := 0; i < len(res); i++ {
		key := int(rune(password[i%len(password)]))
		res[i] = rune(int(res[i]) - key)
	}
	return string(res)
}
