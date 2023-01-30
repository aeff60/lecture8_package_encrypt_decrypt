package lecture8_package_encrypt_decrypt

import "io/ioutil"

func encrypt(filePath string, key string) error {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	for i := range data {
		data[i] ^= key[i%len(key)]
	}

	return ioutil.WriteFile(filePath, data, 0644)
}

func decrypt(filePath string, key string) error {
	return encrypt(filePath, key)
}
