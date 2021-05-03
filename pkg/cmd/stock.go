package cmd

import (
	"log"
	"os"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/davidleitw/pluto/pkg/gui"
	"github.com/davidleitw/pluto/pkg/user"
	"github.com/spf13/cobra"
)

var whole bool = true

var desc = strings.Join([]string{
	"add 功能新增一筆交易紀錄，輸入股票代號，買入價格，數量，手續費折數",
	"輸入完成後可以添加一筆購買紀錄到自己的庫存，以便於做後續的試算損益",
	"",
	"-w --whole 參數默認為 true, 代表整張的交易，單位用張即可",
	"假如是零股交易，就要將該參數設為 false, 買入單位就必須使用股數",
	"",
	"For Example:",
	"pluto add -w=false, 後續輸入單位必須要用股數來表示為零股交易。",
}, "\n")

var stockCmd *cobra.Command = &cobra.Command{
	Use:   "add",
	Short: "新增一筆交易紀錄",
	Long:  desc,
	Run: func(cmd *cobra.Command, args []string) {
		m := gui.AddStockModelInitial(whole)
		if err := tea.NewProgram(m).Start(); err != nil {
			log.Println(err)
			os.Exit(0)
		}
		r := m.GetResult()
		u := user.NewPosition(0.28)
		u.AddStock(r.StockNumber, time.Now(), r.Logs, r.Price, whole)
	},
}

func init() {
	stockCmd.Flags().BoolVarP(&whole, "whole", "w", true, "請輸入是否是整張買入")
}
