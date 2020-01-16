package main

import (
    "log"

	"github.com/hajimehoshi/ebiten"

    "github.com/gwantaek/games/12janggi/global"
    "github.com/gwantaek/games/12janggi/scenemanager"
    "github.com/gwantaek/games/12janggi/scenes"
)

func main() {
    scenemanager.SetScene(&scenes.StartScene{})
    err := ebiten.Run(scenemanager.Update, global.ScreenWidth, global.ScreenHeight, 1.0, "12 Janggi")
    if err != nil {
        log.Fatalf("Ebiten run error: %v", err)
    }

}