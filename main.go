package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")
	//fyne window set size
	w.Resize(fyne.NewSize(600, 600))
	tabs := container.NewAppTabs()
	root_src := "C:\\Users\\ranveer\\Desktop\\Trivial Stuff\\wallpaper"
	files, err := ioutil.ReadDir(root_src)
	// if error is present
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		fmt.Println(file.Name(), file.IsDir())
		if !file.IsDir() {
			// abc.jpg  ->Split arry->["abc","jpg"]
			extension := strings.Split(file.Name(), ".")[1]
			if extension == "jpg" || extension == "png" {
				tabs.Append(container.NewTabItem(file.Name(), canvas.NewImageFromFile(root_src+"\\"+file.Name())))
			}
		}
	}
	tabs.SetTabLocation(container.TabLocationLeading)
	w.SetContent(tabs) // ek hi value hai oh no need of NewVerticalBox
	w.ShowAndRun()
}
