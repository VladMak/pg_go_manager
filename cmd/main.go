package main

import (
        "github.com/gotk3/gotk3/gtk"
)

func main() {
        // Инициализируем GTK.
        gtk.Init(nil)

        b, err := gtk.BuilderNew()
        if err != nil {
        	println("Err")
        }

        err = b.AddFromFile("screen/main.glade")
        if err != nil {
        	println("Bad")
        }

        obj, err := b.GetObject("main_window")
        if err != nil {
        	println("Bad")
        }

        win := obj.(*gtk.Window)
        win.Connect("destroy", func(){
        	gtk.MainQuit()
        })

        win.ShowAll()

        gtk.Main()
}