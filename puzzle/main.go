package main

import (
    "log"
    "math/rand"
    "time"

	"github.com/hajimehoshi/ebiten"

    "github.com/gwantaek/games/puzzle/global"
    "github.com/gwantaek/games/puzzle/scenemanager"
    "github.com/gwantaek/games/puzzle/scenes"
)

func main() {
    rand.Seed(time.Now().UnixNano())
    scenemanager.SetScene(&scenes.StartScene{})
    err := ebiten.Run(scenemanager.Update, global.ScreenWidth, global.ScreenHeight, 1.0, "Puzzle")
    if err != nil {
        log.Fatalf("Ebiten run error: %v", err)
    }
}