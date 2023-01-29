package main

import "encoding/base64"

func main() {

}

type Codec struct {
}

func (this *Codec) Constructor() Codec {
	return Codec{}
}

// Encodes a URL to a shortened URL.
func (this *Codec) encode(longUrl string) string {
	encoding := base64.StdEncoding
	return encoding.EncodeToString([]byte(longUrl))
}

// Decodes a shortened URL to its original URL.
func (this *Codec) decode(shortUrl string) string {
	decodeString, _ := base64.StdEncoding.DecodeString(shortUrl)
	return string(decodeString)
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * url := obj.encode(longUrl);
 * ans := obj.decode(url);
 */
