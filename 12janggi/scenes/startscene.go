package scenes

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
	
	"github.com/gwantaek/games/12janggi/scenemanager"
)

type StartScene struct {
	startImg *ebiten.Image
}

func (s *StartScene) Startup() {
    var err error    
    s.startImg, _, err = ebitenutil.NewImageFromFile("./images/start.png", ebiten.FilterDefault)
    if err != nil {
        log.Fatalf("read file error: %v", err)
    }
}

func (s *StartScene) Update(screen *ebiten.Image) error {
	screen.DrawImage(s.startImg, nil)
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		scenemanager.SetScene(&GameScene{})
	}
	return nil
}