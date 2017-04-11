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

const (
	//WM_SETTINGCHANGE message
	SPIF_UPDATEINIFILE    = 0x0
	SPIF_SENDCHANGE       = 0x1
	SPIF_SENDWININICHANGE = 0x2

	//UIACTION
	SPI_SETDESKWALLPAPER = 0x0014
)

var (
	user32               = syscall.NewLazyDLL("user32.dll") //user32.dll
	systemParametersInfo = user32.NewProc("SystemParametersInfoW")
)

func setBackground(imagePath string) {
	if imagePath != "" {

		ret, _, _ := systemParametersInfo.Call(SPI_SETDESKWALLPAPER, 0, uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(imagePath))), SPIF_SENDWININICHANGE)
		fmt.Println("Background set to: " + imagePath)
		//fmt.Println(ret)
		if ret == 1 {
			os.Exit(0)
		}
		os.Exit(int(ret))
	}

}
