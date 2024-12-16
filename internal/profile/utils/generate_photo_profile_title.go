package utils

import "fmt"

func GeneratePhotoProfileTitle(userID int64, nameUser string) string {
	invoiceNumber := fmt.Sprintf("photo-profile-%v-%v-domain-driven-design", userID, nameUser)
	return invoiceNumber
}
