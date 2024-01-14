package utils
import "fmt"

type Dutils struct {}

func (m *Dutils)MaxInt() uint64 {
	const maxInt uint64 = 1 << 64 - 1
	return maxInt
}

func (m *Dutils)Conv(num int) string {
	return fmt.Sprint(num)
}
