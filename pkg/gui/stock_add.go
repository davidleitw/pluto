package gui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	focusedStyle       = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	blurredButtonStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	noStyle            = lipgloss.NewStyle()

	focusedSubmitButton = "[ " + focusedStyle.Render("Submit") + " ]"
	blurredSubmitButton = "[ " + blurredButtonStyle.Render("Submit") + " ]"
)

type addStockModel struct {
	index            int
	stockNumberInput textinput.Model
	logsInput        textinput.Model
	priceInput       textinput.Model
	submitButton     string
	msg              chan string
}

func AddStockModelInitial() addStockModel {
	stockNumber := textinput.NewModel()
	stockNumber.Placeholder = " 股票編號"
	stockNumber.Focus()
	stockNumber.PromptStyle = focusedStyle
	stockNumber.TextStyle = focusedStyle
	stockNumber.CharLimit = 10

	logs := textinput.NewModel()
	logs.Placeholder = " 單位: 張"
	logs.CharLimit = 10

	price := textinput.NewModel()
	price.Placeholder = " 買入價格"
	price.CharLimit = 10
	return addStockModel{0, stockNumber, logs, price, blurredSubmitButton, make(chan string, 10)}
}

func (m addStockModel) ShowResult() {
	fmt.Println(m.stockNumberInput.Value(),
		m.logsInput.Value(), m.priceInput.Value())
}

func (m addStockModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m addStockModel) View() string {
	s := "\n"

	inputs := []string{
		m.stockNumberInput.View(),
		m.logsInput.View(),
		m.priceInput.View(),
	}

	for i := 0; i < len(inputs); i++ {
		s += inputs[i]
		if i < len(inputs)-1 {
			s += "\n"
		}
	}

	s += "\n\n" + m.submitButton + "\n"
	return s
}

func (m addStockModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit

		// Cycle between inputs
		case "tab", "shift+tab", "enter", "up", "down":

			inputs := []textinput.Model{
				m.stockNumberInput,
				m.logsInput,
				m.priceInput,
			}
			s := msg.String()

			// Did the user press enter while the submit button was focused?
			// If so, exit.
			if s == "enter" && m.index == len(inputs) {
				// m.msg <- inputs[0].Value()
				// m.msg <- inputs[1].Value()
				// m.msg <- inputs[2].Value()
				// return m, tea.Quit
				if inputs[0].Value() == "87" {
					inputs[0].SetValue("你說誰87")
				} else {
					return m, tea.Quit
				}
			}

			// Cycle indexes
			if s == "up" || s == "shift+tab" {
				m.index--
			} else {
				m.index++
			}

			if m.index > len(inputs) {
				m.index = 0
			} else if m.index < 0 {
				m.index = len(inputs)
			}

			for i := 0; i <= len(inputs)-1; i++ {
				if i == m.index {
					// Set focused state
					inputs[i].Focus()
					inputs[i].PromptStyle = focusedStyle
					inputs[i].TextStyle = focusedStyle
					continue
				}
				// Remove focused state
				inputs[i].Blur()
				inputs[i].PromptStyle = noStyle
				inputs[i].TextStyle = noStyle
			}

			m.stockNumberInput = inputs[0]
			m.logsInput = inputs[1]
			m.priceInput = inputs[2]

			if m.index == len(inputs) {
				m.submitButton = focusedSubmitButton
			} else {
				m.submitButton = blurredSubmitButton
			}

			return m, nil
		}
	}

	// Handle character input and blinks
	m, cmd = updateInputs(msg, m)
	return m, cmd
}

func updateInputs(msg tea.Msg, m addStockModel) (addStockModel, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	m.stockNumberInput, cmd = m.stockNumberInput.Update(msg)
	cmds = append(cmds, cmd)

	m.logsInput, cmd = m.logsInput.Update(msg)
	cmds = append(cmds, cmd)

	m.priceInput, cmd = m.priceInput.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}
