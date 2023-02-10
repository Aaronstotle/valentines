package main

import (
	"fmt"
	"github.com/joshdk/preview"
	"image/png"
	"os"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto/v2"
)

type Duckies struct {
}

/*
go get -u github.com/joshdk/preview
TODO: MVP - Small application that has a lake background picture, can click to hear duck noises, and also get pictures of a duck

TODO: Install this library and add it so we can open up images DONE!
BONUS TODO: add multiple images and have the button cycle through them
import (
	"image/jpeg"
	"net/http"
	"github.com/joshdk/preview"
)

resp, err := http.Get("https://i.imgur.com/X9GB4Pu.jpg")
if err != nil {
	panic(err.Error())
}

img, err := jpeg.Decode(resp.Body)
if err != nil {
	panic(err.Error())
}

preview.Image(img)
*/

func play_sound() error {
	f, err := os.Open("ducks.mp3")
	if err != nil {
		return err
	}
	defer f.Close()

	d, err := mp3.NewDecoder(f)
	if err != nil {
		return err
	}

	c, ready, err := oto.NewContext(d.SampleRate(), 2, 2)
	if err != nil {
		return err
	}
	<-ready

	p := c.NewPlayer(d)
	defer p.Close()
	p.Play()

	fmt.Printf("Length: %d[bytes]\n", d.Length())
	for {
		time.Sleep(time.Second)
		if !p.IsPlaying() {
			break
		}
	}

	return nil
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readBackground() {
	duckpng, _ := os.Open("/Users/aaron/code/gopher/duckgame/background.png")
	anotherimageFile, err := png.Decode(duckpng)
	check(err)
	preview.Image(anotherimageFile)

}

func tapped() {
	fmt.Println("Quack!")
	play_sound()

}

func another() {
	fmt.Println("Duck time!")
	readBackground()
}

/*

TODO: create container layout using this format

var := container --> with whatever layout you want

	w.SetContent(var) to show container x1x1

*/

func main() {

	var backgroundPicture = "/Users/aaron/code/gopher/duckgame"
	background := canvas.NewImageFromFile(backgroundPicture)

	a := app.New()
	w := a.NewWindow("Happy Valentines Day!")
	print(background)

	// vday := canvas.NewText("Happy Valentines day", color.White)

	//box1 := container.New(layout.NewVBoxLayout(), vday)

	anotherOne := widget.Button{Text: "Give me another!", OnTapped: another}

	relax := widget.Button{Text: "Feeling Stressed? Click me and close your eyes", OnTapped: tapped}

	container := container.NewGridWithColumns(3,
		layout.NewSpacer(),
		container.NewHBox(&relax),
		layout.NewSpacer(),
		container.NewVBox(&anotherOne),
	)

	w.SetContent(container)

	w.Resize(fyne.NewSize(500, 300))
	w.ShowAndRun()

}
