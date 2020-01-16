package scenes

import (
	"log"
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"

	"github.com/gwantaek/games/puzzle/global"
	"github.com/gwantaek/games/puzzle/scenemanager"
	"github.com/gwantaek/games/puzzle/font"
)

// StartScene startScene
type StartScene struct {
	startImg *ebiten.Image
}

// Startup startup
func (s *StartScene) Startup() {
    var err error    
    s.startImg, _, err = ebitenutil.NewImageFromFile("./images/monalisa.png", ebiten.FilterDefault)
    if err != nil {
        log.Fatalf("read file error: %v", err)
    }
}

// Update update
func (s *StartScene) Update(screen *ebiten.Image) error {
	screen.DrawImage(s.startImg, nil)
	width := font.TextWidth(global.StartSceneText, 2)	
	font.DrawTextWithShadow(screen, global.StartSceneText, global.ScreenWidth/2-width/2, global.ScreenHeight/2, 2, color.Black)
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		scenemanager.SetScene(&GameScene{})
	}
	return nil
}