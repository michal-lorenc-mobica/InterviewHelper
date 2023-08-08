package main

import (
	"bufio"
	"errors"
	"fmt"
	"image/color"
	"math/rand"
	"os"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var myQuestions []string
var myAnswers []string
var examTimeStart time.Time

func loadData(fullPath string) error {

	file, err := os.Open(fullPath)
	if err != nil {
		return err
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanLines)
	// mechanism to separate questions from answers - used convention:
	// is a single line for question (odd line number) and single line for answer (even line number)
	separator := 0
	for fileScanner.Scan() {
		if separator%2 == 0 {
			myQuestions = append(myQuestions, fileScanner.Text())
		} else {
			myAnswers = append(myAnswers, fileScanner.Text())
		}
		separator++
	}

	// checking input data convention
	if len(myQuestions) != len(myAnswers) {
		return errors.New("wrong syntax convention - please check data correctness for your input file")
	}

	return nil
}

func main() {

	err := loadData("./data.txt")
	if err != nil {
		fmt.Printf("%s\n", err.Error())
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

	examLengthLabel := widget.NewLabel("")

	examMode := widget.NewCheck("Exam Mode", func(value bool) {
		fmt.Println("Check set to", value)
	})

	examMode.OnChanged = func(bool) {
		answerText.SetText("")
		questionText.SetText("")
		index = 0
		if examMode.Checked {
			examTimeStart = time.Now()
			examLengthLabel.SetText("")
		} else {
			examLength := time.Since(examTimeStart)
			examLengthString := "Exam took: " + examLength.String()
			examLengthLabel.SetText(examLengthString)
		}
	}

	buttonQuestion := widget.NewButton("Generate question", func() {
		// clean answer field
		answerText.SetText("")

		if !examMode.Checked {
			index = rand.Intn(len(myQuestions))
			questionText.SetText(myQuestions[index])
		} else {
			if index < len(myQuestions) {
				questionText.SetText(myQuestions[index])
				index += 1
			} else {
				examMode.SetChecked(false)
				index = 0
			}
		}
		questionGenerated = true
	})

	buttonAnswer := widget.NewButton("Show answer", func() {

		if questionGenerated && questionText.Text != "" {
			if index > len(myQuestions) {
				examMode.SetChecked(false)
				index = 0
			} else {
				if examMode.Checked {
					answerText.SetText(myAnswers[index-1])
				} else {
					answerText.SetText(myAnswers[index])
				}
			}
		} else {
			answerText.SetText("Please generate question first")
			index = 0
		}
	})

	line := canvas.NewLine(color.Gray{})
	line.StrokeWidth = 1

	textLabelBar := container.New(layout.NewVBoxLayout(), questionLabel, questionText, layout.NewSpacer(), spacer, answerLabel, answerText)
	buttonsBar := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), buttonQuestion, buttonAnswer, examMode, layout.NewSpacer(), examLengthLabel)
	myWindow.SetContent(container.New(layout.NewVBoxLayout(), textLabelBar, layout.NewSpacer(), line, buttonsBar))
	myWindow.ShowAndRun()
}
