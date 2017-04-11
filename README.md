
# ChangeBackground
A simple tool for changing Windows Background using Golang.

Many thanks to this:
https://www.reddit.com/r/golang/comments/55guj9/need_help_with_my_code/



```
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
```
