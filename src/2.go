 package main
 import "math/rand"
 func generateCode() string {
     const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
     b := make([]byte, 6)
     for i := range b {
         b[i] = letters[rand.Intn(len(letters))]
     }
     return string(b)
 }