package main

import (
    "fmt"
    "encoding/binary"
    //"io"
    "bufio"
    "os"
)

func main() {
    s := "Dag"

    fi, err := os.Create("output.file")
    if err != nil { panic(err) }

    defer func() {
        if err := fi.Close(); err != nil {
            panic(err)
        }
    }()

    w := bufio.NewWriter(fi)
    fmt.Println([]byte(s))

    err = binary.Write(w, binary.LittleEndian, []byte(s))
    if err != nil {
        fmt.Println("Error: ", err)
    }
    if err = w.Flush(); err != nil { panic(err) }
}
