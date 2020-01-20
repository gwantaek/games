package scenes

import (   
	"github.com/hajimehoshi/ebiten"
	
	"github.com/gwantaek/games/runner/global"
	"github.com/gwantaek/games/runner/actor"
	"github.com/gwantaek/games/runner/bgscroller"
)

// GameScene gameScene
type GameScene struct {
	bgscroller	*bgscroller.Scroller
	runner 		*actor.Runner
}

// Startup startUp
func (g *GameScene) Startup() {
    // Make a new scolling background
	g.bgscroller = bgscroller.New(global.BGScrollSpeed)

	// Make a new running actor
	g.runner = actor.NewRunner(0, global.ScreenHeight/2)
	g.runner.SetState(actor.Running)
}

// Update update scene
func (g *GameScene) Update(screen *ebiten.Image) error {

	// Draw background image
	g.bgscroller.Update(screen)

	// Draw running animation
	g.runner.Update(screen)

    return nil
}