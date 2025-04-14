package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
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

func getShortedUrl(id string) (Url , error) {
	url,ok := urlBD[id]
	if !ok {
		return Url{} , errors.New("Url Not Found")
	}
	return url,nil
}

func rootHandler(w http.ResponseWriter , r *http.Request){
	fmt.Println("GET request only")
	fmt.Fprintf(w,"hello world")
}

func ShortURLHander (w http.ResponseWriter , r *http.Request){
	var data struct{
		URL string `json:"url"`
	}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err!= nil {
		http.Error(w,"Invalid req body",http.StatusBadRequest)
	}
}

func main() {
	fmt.Println("Url Shortner Server running at the port 4000....");

	// function for listening for the request to the /
	http.HandleFunc("/",rootHandler)

	// starting the http server
	err := http.ListenAndServe(":4000",nil)
	if( err !=nil) {
		fmt.Println("An Error Occured While Starting the Server",err)
	}
}
