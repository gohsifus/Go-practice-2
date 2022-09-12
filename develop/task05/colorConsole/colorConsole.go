package colorConsole

import (
	"os"
	"strings"
	"syscall"
)

//Чтобы работал цветной вывод
func init() {
	stdout := syscall.Handle(os.Stdout.Fd())

	var originalMode uint32
	syscall.GetConsoleMode(stdout, &originalMode)
	originalMode |= 0x0004

	syscall.MustLoadDLL("kernel32").MustFindProc("SetConsoleMode").Call(uintptr(stdout), uintptr(originalMode))
}

//Описывает последовательности для управления цветом
//строки после red будут иметь красный увет
//строки после reset вернут свой цвет
var (
	Red   = "\033[31m"
	Reset = "\033[0m"
)

//GetColorizeString вернет строку которая при выводе в консоль будет окрашена
func GetColorizeString(str string) string {
	builder := strings.Builder{}
	builder.WriteString(Red)
	builder.WriteString(str)
	builder.WriteString(Reset)

	return builder.String()
}
