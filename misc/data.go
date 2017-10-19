package misc

var (
	WingoWav      []byte
	WingoPng      []byte
	ClosePng      []byte
	MinimizePng   []byte
	MaximizePng   []byte
)

func ReadData() {
	WingoWav = DataFile("wingo.wav")
	WingoPng = DataFile("wingo.png")
	ClosePng = DataFile("close.png")
	MinimizePng = DataFile("minimize.png")
	MaximizePng = DataFile("maximize.png")
}
