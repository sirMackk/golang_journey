package main

import (
    "fmt"
    "bytes"
    "os"
    "errors"
    //"encoding/binary"
    //"bufio"
    //"os"
)

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

func main() {
    db := NewDB()
    db.Insert("Bob")
    db.Insert("Jay")
    db.Insert("Ray")
    db.ShowAll()
    db.Show(0)
    r, err := db.Find("Jay")
    if err != nil {
        fmt.Println("Couldnt find shit")
        os.Exit(1)
    }
    r.Print()
    db.Delete("Jay")
    db.ShowAll()
}
