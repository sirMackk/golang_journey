package main

import (
    "fmt"
    "bufio"
    "encoding/binary"
    "os"
)

type Dag struct {
    Id uint8
    Name string
}

func (d *Dag) Serialize() []uint8 {
    arr := make([]uint8, 32)
    arr[0] = d.Id
    for i, val := range([]byte(d.Name)) {
        arr[i + 1] = val
        if i == 32 { break }
    }
    return arr
}

func Deserialize(input []uint8) *Dag {
    return &Dag{input[0], string(input[1:])}
}

func main() {
    //Create pointer to Dag struct
    dag := Dag{1, "Axle"}

    //open file for reading, return file handle
    file, err := os.Create("dag.file")
    if err != nil { panic(err) }

    //use defer to properly close file or raise error
    defer func() {
        if err := file.Close(); err != nil {
            panic(err)
        }
    }()

    //create write buffer using bufio and file handle
    w := bufio.NewWriter(file)

    //user binary package to write serialized data to bufio buffer
    err = binary.Write(w, binary.LittleEndian, dag.Serialize())
    if err != nil { panic(err) }

    //flush buffer (write to disk)
    if err = w.Flush(); err != nil { panic(err) }

    //open file for reading
    open, err := os.Open("dag.file")

    if err != nil { panic(err) }

    //again defer closing and take care of erros
    defer func() {
        if err := open.Close(); err != nil {
            panic(err)
        }
    }()

    //read file into bufio buffer (?)
    r := bufio.NewReader(open)

    //create struct sized buffer of bytes
    buf := make([]byte, 32)

    //binary read bufio buffer into buffer
    err = binary.Read(r, binary.LittleEndian, &buf)

    if err != nil { panic(err) }

    //deserialize contents of buffer
    ha := Deserialize(buf)

    fmt.Println("Deserialized input: ")
    fmt.Println(*ha)
}

