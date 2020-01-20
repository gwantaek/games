package scenes

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"

	"github.com/gwantaek/games/runner/global"
	"github.com/gwantaek/games/runner/scenemanager"
	"github.com/gwantaek/games/runner/font"
	"github.com/gwantaek/games/runner/actor"
	"github.com/gwantaek/games/runner/bgscroller"
)

// StartScene startScene
type StartScene struct {
	bgscroller	*bgscroller.Scroller
	runner		*actor.Runner
}

// Startup startup
func (s *StartScene) Startup() {	
	// Make a new non-scolling (speed zero) background
	s.bgscroller = bgscroller.New(0)
	
	// Make a new idle actor
	s.runner = actor.NewRunner(0, global.ScreenHeight/2)
	s.runner.SetState(actor.Idle)
}

// Update update
func (s *StartScene) Update(screen *ebiten.Image) error {
	// Draw background image
	s.bgscroller.Update(screen)

	// Draw idle animation
	s.runner.Update(screen)

	// Display test message
	width := font.TextWidth(global.StartSceneText, 2)	
	font.DrawTextWithShadow(screen, global.StartSceneText, global.ScreenWidth/2-width/2, global.ScreenHeight/2, 2, color.White)

	// Start key event
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		scenemanager.SetScene(&GameScene{})
	}

	return nil
}