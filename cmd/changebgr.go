package cmd

import (
	"fmt"
	"os"
	"syscall"
	"unsafe"

	"github.com/spf13/cobra"
)

var changebgr = &cobra.Command{
	Use:   "set",
	Short: "set <file>",
	Long:  `Change background using this simple tool.`,
	Run: func(cmd *cobra.Command, args []string) {
		if _, err := os.Stat(args[0]); err == nil {
			// path/to/whatever exists
			setBackground(args[0])
		} else {
			fmt.Println("Does not exist")
		}

	},
}

var (
	user32               = syscall.NewLazyDLL("user32.dll") //user32.dll
	systemParametersInfo = user32.NewProc("SystemParametersInfoW")
)

func setBackground(imagePath string) {
	if imagePath != "" {
		fmt.Println("Setting BG to: " + imagePath)
		ret, _, _ := systemParametersInfo.Call(20, 0, uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(imagePath))), 2)
		fmt.Println(ret)
	}

}
