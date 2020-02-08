package internal

import (
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/kyokomi/emoji"
	"github.com/lzecca78/git-cloner/config"
	"github.com/olekukonko/tablewriter"
)

var data [][]string

//List will list all configured repos in config file
func List(cfg *config.GitConfig) {
	for i, v := range cfg.Repos {
		remote := v.Git_Remote
		var emojiIcon string
		cleanRemoteBase := strings.Replace(path.Base(v.Git_Remote), ".git", "", -1)
		local := path.Join(v.LocalDir, cleanRemoteBase)
		if alreadyCloned(local) {
			emojiIcon = emoji.Sprint(":check_mark_button:")
		} else {
			emojiIcon = emoji.Sprint(":cross_mark:")
		}
		data = append(data, []string{strconv.Itoa(i + 1), remote, local, emojiIcon})
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Number", "Remote", "LocalDir", "Status"})
	table.SetRowLine(true)
	for _, v := range data {
		table.Append(v)
	}
	table.Render()
}

func alreadyCloned(local string) bool {
	if _, err := os.Stat(local); os.IsNotExist(err) {
		return false
	}
	return true
}
