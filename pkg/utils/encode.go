package utils

import "encoding/base64"

func Btoa(payload string) string {
	return base64.StdEncoding.EncodeToString([]byte(payload))
}
