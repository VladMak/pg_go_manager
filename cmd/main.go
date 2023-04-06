package main

// "github.com/VladMak/pg_go_manager/internal/domain"
// "github.com/gotk3/gotk3/gtk"
// "log"

func main() {

	// 	//Инициализируем ѝпиѝок ѝерверов
	// 		var serverList *domain.ServerList = domain.NewServerList()
	// 		serverList.AddServer("test", "localhost", 5432)

	// 		// Инициализируем GTK.
	// 		gtk.Init(nil)

	// 		b, err := gtk.BuilderNew()
	// 		if err != nil {
	// 			println("Err")
	// 		}

	// 		err = b.AddFromFile("screen/main.glade")
	// 		if err != nil {
	// 			println("Bad")
	// 		}

	// 		obj, err := b.GetObject("main_window")
	// 		if err != nil {
	// 			println("Bad")
	// 		}

	// 		win := obj.(*gtk.ApplicationWindow)
	// 		win.Connect("destroy", func() {
	// 			gtk.MainQuit()
	// 		})

	// 		listBox, err := b.GetObject("list_box")
	// 		if err != nil {
	// 			println("Bad")
	// 		}

	// 		listBoxObj, ok := listBox.(*gtk.ListBox)
	// 		if !ok {
	// 			println("Bad")
	// 		}

	// 		serverListBox, err := b.GetObject("serverListBox")
	// 		if err != nil {
	// 			println("Bad")
	// 		}

	// 		serverListBoxObj, ok := serverListBox.(*gtk.ListBox)
	// 		if !ok {
	// 			println("Bad")
	// 		}

	// 		fillLeftPart(serverListBoxObj, serverList)
	// 		fillListBox(listBoxObj)

	// 		win.ShowAll()
	// 		gtk.Main()
	// 	}

	// 	func fillLeftPart(listBox *gtk.ListBox, serverList *domain.ServerList) {
	// 		for i := 0; i < len(serverList.ServerItems) + 1; i++ {
	// 			row, err := gtk.ListBoxRowNew()
	// 			if err != nil {
	// 				log.Fatal("Ошибка:", err)
	// 			}

	// 			if i == 0 {
	// 				button, err := gtk.ButtonNew()
	// 				if err != nil {
	// 					log.Fatal("Ошибка:", err)
	// 				}
	// 				button.SetLabel("+")
	// 				row.Add(button)
	// 				listBox.Add(row)
	// 				continue
	// 			}
	// 		}
	// 	}

	// 	func fillListBox(listBox *gtk.ListBox) {
	// 		for i := 0; i < 10; i++ {
	// 			row, err := gtk.ListBoxRowNew()
	// 			if err != nil {
	// 				log.Fatal("Ошибка при ѝоздании ѝтроки ѝпиѝка:", err)
	// 			}

	// 			button, err := gtk.ButtonNew()
	// 			if err != nil {
	// 				log.Fatal("Ошибка при ѝоздании метки:", err)
	// 			}
	// 			button.SetLabel("Hello")
	// 			button.SetBorderWidth(0)
	// 			row.Add(button)
	// 			listBox.Add(row)
	// 		}
}
