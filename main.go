package main

import (
	"fmt"
	"log"
	"net/http"
)

func final(w http.ResponseWriter, r *http.Request) {
	GetLogger().error("This is error message")
	GetLogger().info("This is info message")
}

func main() {
	version := "1.0.1"
	fmt.Println("Version==>", version)

	// logger := GetLogger()
	// logger.info("Info Here..") // to use logger_v1.go file
	// logger.info.Printf("Printing Some Info here...")  // to use logger.go file
	// logger.error.Printf("Printing Some Error here...")
	// logger.warning.Printf("Printing Some Warning here...")
	// logger.fatal.Printf("Printing Some Fatal here...")

	// Route to test the logger log whether is working fine or not.
	mux := http.NewServeMux()

	finalHandler := http.HandlerFunc(final)
	mux.Handle("/", LogMiddleWare(finalHandler))

	log.Println("Listening on :3000...")
	err := http.ListenAndServe(":3000", mux)
	log.Fatal(err)

}
