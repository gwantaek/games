package main

import (
    "log"

	"github.com/hajimehoshi/ebiten"

    "github.com/gwantaek/games/runner/global"
    "github.com/gwantaek/games/runner/scenemanager"
    "github.com/gwantaek/games/runner/scenes"
)

func main() {
    scenemanager.SetScene(&scenes.StartScene{})
    err := ebiten.Run(scenemanager.Update, global.ScreenWidth, global.ScreenHeight, 1.0, "Runner")
    if err != nil {
        log.Fatalf("Ebiten run error: %v", err)
    }
}