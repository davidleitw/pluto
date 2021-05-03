package gui

// stock_add.go: 新增庫存/交易紀錄時顯示的gui樣板

import (
	"strconv"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/davidleitw/pluto/pkg/stock"
)

type addStockResult struct {
	StockNumber int
	Logs        int
	Price       float64
}

type addStockModel struct {
	index            int
	stockNumberInput textinput.Model
	logsInput        textinput.Model
	priceInput       textinput.Model
	errorMessage     string
	submitButton     string
	msg              chan *addStockResult
}

func AddStockModelInitial(wholeStock bool) addStockModel {
	stockNumber := textinput.NewModel()
	stockNumber.Placeholder = " 股票編號"
	stockNumber.Focus()
	stockNumber.PromptStyle = focusedStyle
	stockNumber.TextStyle = focusedStyle
	stockNumber.CharLimit = 40

	logs := textinput.NewModel()
	if wholeStock {
		logs.Placeholder = " 單位: 張"
	} else {
		logs.Placeholder = " 單位: 零股"
	}
	logs.CharLimit = 40

	price := textinput.NewModel()
	price.Placeholder = " 買入價格"
	price.CharLimit = 40
	return addStockModel{0, stockNumber, logs, price, "", blurredSubmitButton, make(chan *addStockResult, 1)}
}

func (m addStockModel) GetResult() *addStockResult {
	select {
	case r := <-m.msg:
		return r
	default:
		return &addStockResult{}
	}
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
		m.errorMessage,
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
	var errIndex int = 0
	var errMsg string = ""
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
				company, err := stock.GetStockNameByNumber(inputs[0].Value())
				if err != nil {
					inputs[0].SetValue("")
					errMsg = err.Error()
					errIndex = 0
					goto _continue
				} else if company == "" {
					inputs[0].SetValue("")
					errMsg = " 無法找到對應的公司，請再次確認輸入的股票編號。"
					errIndex = 0
					goto _continue
				}

				logs, err := strconv.Atoi(inputs[1].Value())
				if err != nil {
					inputs[1].SetValue("")
					errMsg = err.Error()
					errIndex = 1
					goto _continue
				} else if logs <= 0 {
					inputs[1].SetValue("")
					errMsg = "交易單位不能小於等於0，請再次確認輸入的單位數量。"
					errIndex = 1
					goto _continue
				}

				price, err := strconv.ParseFloat(inputs[2].Value(), 2)
				if err != nil {
					inputs[2].SetValue("")
					errMsg = err.Error()
					errIndex = 2
					goto _continue
				} else if price <= 0.0 {
					inputs[2].SetValue("")
					errMsg = " 成交金額不能小於等於0，請再次確認輸入的成交金額。"
					errIndex = 2
					goto _continue
				}
				n, _ := strconv.Atoi(inputs[0].Value())
				m.msg <- &addStockResult{StockNumber: n, Logs: logs, Price: price}
				return m, tea.Quit
			}
		_continue:
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

			// 當錯誤發生的時候，跳轉到錯誤發生的欄位。
			if errMsg != "" {
				m.index = errIndex
				m.errorMessage = colorFg(errMsg, "212")
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
