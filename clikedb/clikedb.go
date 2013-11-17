package clikedb

import (
    "fmt"
    "bytes"
    "os"
    "errors"
    "encoding/binary"
    "bufio"
)

const RecordSize = 32

type Record struct {
    Name []byte
}

type Database struct {
    Size uint8
    D []Record
}

func NewDB() *Database {
    return &Database{Size: 0, D: make([]Record, 10)}
}

func (self *Database) Insert(name string) {
    self.D[self.Size] = Record{Name: []byte(name)}
    self.Size += 1
}

func (self *Database) InsertByteString(name []byte) {
    fmt.Println(name)
    self.D[self.Size] = Record{Name: name}
    self.Size += 1
}

func (self *Database) Show(id uint8) {
    fmt.Println(string(self.D[id].Name))
}

func (self *Database) ShowAll() {
    for _, val := range(self.D) {
        fmt.Println(val)
    }
}

func (self *Database) Find(name string) (*Record, error) {
    for _, val := range(self.D) {
        //if val.Name == []byte(name) {
        if bytes.Equal(val.Name, []byte(name)) {
            return &val, nil
        }
    }
    return nil, errors.New("Damnit!")
}

func (self *Database) Delete(name string) error {
    for i, val := range(self.D) {
        if bytes.Equal(val.Name, []byte(name)) {
            //val.Name = nil
            self.D = append(self.D[:i], self.D[i+1:]...)
            self.D = append(self.D, *new(Record))
            self.Size -= 1
            return nil
        }
    }
    return errors.New("Nothing deleted")
}

func (self *Record) Print() {
    fmt.Println(string(self.Name))
}

func (self *Record) Serialize() []byte {
    arr := make([]byte, RecordSize)
    copy(arr, self.Name)
    return arr
}

func Deserialize(input []byte) *Record {
    return &Record{input}
}

func (self *Database) Serialize() []byte {
    arr := make([]byte, 0)
    for i := 0; i < int(self.Size); i++ {
        arr = append(arr, self.D[i].Serialize()...)
    }
    return arr
}

func (self *Database) SaveDB() {
    file, err := os.Create("database.db")
    if err != nil { panic(err) }
    defer func() {
        if err := file.Close(); err != nil {
          panic(err)
        }
    }()

    bwriter := bufio.NewWriter(file)
    err = binary.Write(bwriter, binary.LittleEndian, self.Serialize())
    if err != nil { panic(err) }

    if err = bwriter.Flush(); err != nil { panic(err) }
}

func (self *Database) Deserialize(s []byte) {
    for i := 0; i < len(s); i += RecordSize {
        self.InsertByteString(s[i:i+31])
    }
}

func (self *Database) ReadDB() {
    file, err := os.Open("database.db")
    if err != nil { panic(err) }
    defer func() {
        if err := file.Close(); err != nil {
            panic(err)
        }
    }()

    breader := bufio.NewReader(file)
    buf := make([]byte, 96)
    err = binary.Read(breader, binary.LittleEndian, &buf)
    if err != nil { panic(err) }

    self.Deserialize(buf)
}

