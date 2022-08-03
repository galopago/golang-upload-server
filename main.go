package main

import (
    "fmt"    
    "io/ioutil"
    "net/http"
    "time"
    "os"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
    fmt.Println("File Upload Endpoint Hit")

    // Parse our multipart form, 10 << 20 specifies a maximum
    // upload of 10 MB files.
    r.ParseMultipartForm(10 << 20)
    // FormFile returns the first file for the given key `myFile`
    // it also returns the FileHeader so we can get the Filename,
    // the Header and the size of the file
    KeyName := "myFile"							// must match POST keyname!!
    file, handler, err := r.FormFile(KeyName)
    if err != nil {
        fmt.Println("Error Retrieving the File")
        fmt.Println(err)
        return
    }
    defer file.Close()
    fmt.Printf("Uploaded File: %+v\n", handler.Filename)
    fmt.Printf("File Size: %+v\n", handler.Size)
    fmt.Printf("MIME Header: %+v\n", handler.Header)

    // Create a temporary file within our temp-images directory that follows
    // a particular naming pattern

	fpname := GetFilenameDate()
	fp, err := os.Create(fpname)
    if err != nil {
        fmt.Println(err)
    }
    defer fp.Close()
    
    // read all of the contents of our uploaded file into a
    // byte array
    fileBytes, err := ioutil.ReadAll(file)
    
    if err != nil {
        fmt.Println(err)
    }
    // write this byte array to our temporary file
    _, err2 := fp.Write(fileBytes)

    if err2 != nil {
        fmt.Println(err2)
    }
    fp.Sync()
       
    // return that we have successfully uploaded our file!
    fmt.Printf("Successfully Uploaded File:%s\n",fpname)
        
    
}

func setupRoutes() {
    http.HandleFunc("/upload", uploadFile)		// Endpoint
    http.ListenAndServe(":8080", nil)			// Listen port
}

func main() {
    fmt.Println("Server started!")
    setupRoutes()
        
}
func GetFilenameDate() string {
    
    now := time.Now()      // current local time
	sec := now.Unix()      // number of seconds since January 1, 1970 UTC    
    return "uploaded-images/"+"IMG" + fmt.Sprintf("%d",sec) + ".jpg"	// Directory must exist!!
}



