package vars

type fnBinding struct {
	fnID int
	fn   func(string) string
}

var (
	fnMap     = make(map[int]func(string) string)
	fnMapChan = make(chan fnBinding, 1)
)

func init() {
	go func() {
		for bd := range fnMapChan {
			fnMap[bd.fnID] = bd.fn
		}
	}()
}

func GetFn(fnID int) func(string) string {
	return fnMap[fnID]
}

func SetFn(fnID int, fn func(string) string) {
	fnMapChan <- fnBinding{fnID: fnID, fn: fn}
}
