package main

import (
	"image/color"
	"math/rand"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {

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

		index = rand.Intn(len(questions))
		questionText.SetText(questions[index])
		questionGenerated = true
	})

	buttonAnswer := widget.NewButton("Show answer", func() {
		if questionGenerated {
			answerText.SetText(answers[index])
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
