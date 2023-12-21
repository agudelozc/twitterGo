package bd

import (
   "golang.org/x/crypto/bcrypt")

func EncriptarPassword(password string) (string, error) {
	costo := 8
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), costo)
	if err!= nil {
        return err.Error(), err
    }
    return string(bytes), nil
}