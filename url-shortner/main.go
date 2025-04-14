package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"
)

// create  a interface for the each url entry

type Url struct {
	Id       string `json:"id"`
	URL      string `json:"URL"`
	ShortUrl string `json:"ShortUrl"`
	CreatedAt time.Time `json:"CreatedAt"`
}

// creating a map for storing the url
var urlBD = make(map[string]Url)

// FUNCTION FOR SHORTING THE URL

func shortURl(URL string) string  {
	// create an hash value for the url first
	hasher := md5.New()
	hasher.Write([]byte(URL)) //converting the url to byte
	data := hasher.Sum(nil); // converting the byte to a array of byte
	hash := hex.EncodeToString(data) // coverting the byte array to string 
fmt.Println("encoded string ",hash)
	return hash[:8] // returnning the first 8 letters of the final string 
}

//FUNCTION FOR STORING THE GENERATED URL

func createURL(URL string) string  {
	shortURL := shortURl(URL);
	id := shortURL
	urlBD[id]=Url{
		Id: id,
		URL: URL,
		ShortUrl: shortURL,
		CreatedAt: time.Now(),
	}
	return shortURL;
}

func main() {
	fmt.Println("Url Shortner Server running....");
}
