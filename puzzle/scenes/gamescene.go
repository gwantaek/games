package scenes

import (
    "log"
    "image"
    "math/rand"
    
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
	
	"github.com/gwantaek/games/puzzle/global"
)

type GameScene struct {
    bgimg           *ebiten.Image
    subImages       [global.PuzzleColumns * global.PuzzleRows]*ebiten.Image
    board           [global.PuzzleColumns][global.PuzzleRows]int
    blankX, blankY  int
}

func (g *GameScene) Startup() {
    var err error    
    g.bgimg, _, err = ebitenutil.NewImageFromFile("./images/monalisa.png", ebiten.FilterDefault)
    if err != nil {
        log.Fatalf("read file error: %v", err)
    }

    g.blankX = global.PuzzleColumns-1
    g.blankY = global.PuzzleRows-1

    // Image 자르기
    for i := 0; i < global.PuzzleColumns; i++ {
        for j := 0; j < global.PuzzleRows; j++ {
            g.subImages[j*global.PuzzleColumns + i] = g.bgimg.SubImage(image.Rect(i*global.SubImageWidth, j*global.SubImageHeight, i*global.SubImageWidth+global.SubImageWidth, j*global.SubImageHeight+global.SubImageHeight)).(*ebiten.Image)
        }
    }

    // 이미지 섞기 인덱스 초기화
    arr := make([]int, global.PuzzleColumns * global.PuzzleRows-1)
    idx := 0
    for i := 0; i < global.PuzzleColumns; i++ {
        for j := 0; j < global.PuzzleRows; j++ {
            if i == g.blankX && j == g.blankY {
                continue
            }
            arr[j*global.PuzzleColumns+i] = idx
            idx++
        }
    }

    // 이미지 섞기
    for i := 0; i < global.PuzzleColumns; i++ {
        for j := 0; j < global.PuzzleRows; j++ {
            if i == g.blankX && j == g.blankY {
                g.board[i][j] = -1
            } else {
                idx := rand.Intn(len(arr))
                g.board[i][j] = arr[idx]
                arr = append(arr[:idx], arr[idx+1:]...)    
            }
        }
    }
    
}

// Update update scene
func (g *GameScene) Update(screen *ebiten.Image) error {

    if inpututil.IsKeyJustReleased(ebiten.KeyUp) {
        if g.blankY > 0 {
            g.board[g.blankX][g.blankY] = g.board[g.blankX][g.blankY-1]
            g.board[g.blankX][g.blankY-1] = -1
            g.blankY--
        } 

    } else if inpututil.IsKeyJustReleased(ebiten.KeyDown) {
        if g.blankY < global.PuzzleRows-1 {
            g.board[g.blankX][g.blankY] = g.board[g.blankX][g.blankY+1]
            g.board[g.blankX][g.blankY+1] = -1
            g.blankY++
        } 
            
    } else if inpututil.IsKeyJustReleased(ebiten.KeyLeft) {
        if g.blankX > 0 {
            g.board[g.blankX][g.blankY] = g.board[g.blankX-1][g.blankY]
            g.board[g.blankX-1][g.blankY] = -1
            g.blankX--
        } 
            
    } else if inpututil.IsKeyJustReleased(ebiten.KeyRight) {
        if g.blankX < global.PuzzleColumns-1 {
            g.board[g.blankX][g.blankY] = g.board[g.blankX+1][g.blankY]
            g.board[g.blankX+1][g.blankY] = -1
            g.blankX++
        }             
    }

    for i := 0; i < global.PuzzleColumns; i++ {
        for j := 0; j < global.PuzzleRows; j++ {

            if g.board[i][j] == -1 {
                continue
            }
            x := i * global.SubImageWidth
            y := j * global.SubImageHeight
            opts := &ebiten.DrawImageOptions{}            
            opts.GeoM.Translate(float64(x), float64(y))

            screen.DrawImage(g.subImages[g.board[i][j]], opts)
        }
    }
  
    return nil
}