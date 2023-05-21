package main

import (
	"bufio"
	"fmt"
	"image/color"
	"math/rand"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var myQuestions []string
var myAnswers []string

func loadData(fullPath string) error {

	file, err := os.Open(fullPath)
	if err != nil {
		return err
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanLines)
	// mechanism to separate questions from answers - used convenction:
	// is a single line for question (even line number) and single line for answer (odd line number)
	separator := 0
	for fileScanner.Scan() {
		if separator%2 == 0 {
			myQuestions = append(myQuestions, fileScanner.Text())
		} else {
			myAnswers = append(myAnswers, fileScanner.Text())
		}
		separator++
	}
	return nil
}

func main() {

	err := loadData("./data.txt")
	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	myApp := app.New()
	myWindow := myApp.NewWindow("InterviewHelper")
	myWindow.Resize(fyne.NewSize(900, 600))

	questionLabel := canvas.NewText("Question:", color.White)
	questionLabel.TextSize = 18

	questionText := widget.NewLabel("")
	questionText.Wrapping = fyne.TextWrapWord

	answerLabel := canvas.NewText("Answer:", color.White)
	answerLabel.TextSize = 18

	answerText := widget.NewLabel("")
	answerText.Wrapping = fyne.TextWrapWord

	spacer := canvas.NewText("", color.Black)
	spacer.TextSize = 50

	index := 0
	questionGenerated := false

	buttonQuestion := widget.NewButton("Generate question", func() {
		// clean answer field
		answerText.SetText("")

		index = rand.Intn(len(myQuestions))
		questionText.SetText(myQuestions[index])
		questionGenerated = true
	})

	buttonAnswer := widget.NewButton("Show answer", func() {
		if questionGenerated {
			answerText.SetText(myAnswers[index])
		} else {
			answerText.SetText("Please generate question first")
		}
	})

	line := canvas.NewLine(color.Gray{})
	line.StrokeWidth = 1

	textLabelBar := container.New(layout.NewVBoxLayout(), questionLabel, questionText, layout.NewSpacer(), spacer, answerLabel, answerText)
	buttonsBar := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), buttonQuestion, buttonAnswer, layout.NewSpacer())
	myWindow.SetContent(container.New(layout.NewVBoxLayout(), textLabelBar, layout.NewSpacer(), line, buttonsBar))
	myWindow.ShowAndRun()
}
