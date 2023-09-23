package main

import (
    "fmt"
    "os"
)

func main() {
    // Specify the file path
    filePath := "files/dockerfile"

    // Create or open the file for writing
    file, err := os.Create(filePath)
    if err != nil {
        fmt.Println("Error creating file:", err)
        return
    }
    defer file.Close() // Close the file when done

    // Write data to the file
    data := []byte("FROM golang:1.21.1-bookworm\n")
    _, err = file.Write(data)
    if err != nil {
        fmt.Println("Error writing to file:", err)
        return
    }

    fmt.Println("File created and data written successfully.")
}
