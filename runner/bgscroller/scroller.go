package bgscroller

import (
    "log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

// Scroller background controller
type Scroller struct {
	bgimg	*ebiten.Image
	speed	int
	frames	int
}

// New make a new background
func New(speed int) *Scroller {
	var err error
	bgimg, _, err := ebitenutil.NewImageFromFile("./images/background.png", ebiten.FilterDefault)	
    if err != nil {
        log.Fatalf("read file error: %v", err)
	}

	return &Scroller{bgimg, speed, 0}
}

// Update update background
func (s *Scroller) Update(screen *ebiten.Image) {

	if s.speed > 0 {
		s.frames++

		bgWidth, _ := s.bgimg.Size()
	
		op := &ebiten.DrawImageOptions{}
		backX := (s.frames / s.speed) % bgWidth
	
		op.GeoM.Translate(float64(-backX), 0)
		screen.DrawImage(s.bgimg, op)
	
		op.GeoM.Translate(float64(bgWidth), 0)
		screen.DrawImage(s.bgimg, op)	

	} else {
		screen.DrawImage(s.bgimg, nil)		
	}
}