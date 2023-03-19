package bd

import "golang.org/x/crypto/bcrypt"

// Encriptar password es la rutina que permite encriptar la passwoerd recibida
func EncriptarPassword(pass string) (string, error) {
	costo := 8 // 6 para password normal 8 para admin
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)
	return string(bytes), err
}
