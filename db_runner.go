package main

import "./clikedb"

func main() {
    //var db clikedb.Database
    db := clikedb.NewDB()
    db.ReadDB()
    db.Show(0)
    //db := clikedb.NewDB()
    //db.Insert("Bob")
    //db.Insert("Jay")
    //db.Insert("Ray")
    //db.ShowAll()
    ////db.Show(0)
    ////r, err := db.Find("Jay")
    ////if err != nil {
        ////fmt.Println("Couldnt find shit")
        ////os.Exit(1)
    ////}
    ////r.Print()
    ////db.Delete("Jay")
    ////db.ShowAll()
    //db.SaveDB()
    //db.clikedb.ReadDB()
    //db.clikedb.ShowAll()
    //db.cliekdb.Show(0)
}
