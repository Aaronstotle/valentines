package main

import (
	"fmt"
	"fyne.io/fyne/v2/canvas"
	"github.com/joshdk/preview"
	"image/png"
	"os"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto/v2"
)

/*

TODO: MVP - Small application that has a lake background picture, can click to hear duck noises, and also get pictures of a duck

TODO: DONE Install this library (go get -u github.com/joshdk/preview) and add it so we can open up images

BONUS TODO: add multiple images and have the button cycle through them

Sample code:


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

func giveDuck() {
	// Create function that can read and present a number of different pictures in a folder, can
	duckpng, _ := os.Open("/Users/aaron/code/gopher/duckgame/duck.png")
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
	giveDuck()
}

/*

TODO: Find out how to layer containers, need one for background, another for the widgets
*/

func main() {

	a := app.New()
	w := a.NewWindow("Happy Valentines Day!")

	anotherOne := widget.Button{Text: "Duck?", OnTapped: another}

	relax := widget.Button{Text: "Feeling Stressed? Click me and close your eyes", OnTapped: tapped}

	background := canvas.NewImageFromResource(resourceNewBackgroundPng)

	//evilContainer := container.NewMax(&anotherOne, &relax, background)

	container := container.NewMax(background,
		container.NewGridWithColumns(2,
			container.NewVBox(&anotherOne),
			container.NewVBox(&relax),
		))

	w.SetContent(container)

	w.Resize(fyne.NewSize(500, 500))
	w.ShowAndRun()

}
