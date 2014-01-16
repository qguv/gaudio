package main

import (
    "github.com/ajhager/vox"
    "path/filepath"
    "strconv"
    "time"
)

func errTest(err error) bool {
    if err != nil {
        panic(err)
        return true
    } else {
        return false
    }
}

func getSongLength(song vox.Song) int {
    lines := song.Lines()
    tick := song.TicksPerLine()
    rate := int(vox.TicksPerSecond())
    return lines * tick / rate
}

func studder(songPath string, studderTime int) {
    songName := filepath.Base(songPath)
    errTest(vox.Init("", 44100, 2, 0))
    println("Vox!\nv" + vox.Version + "\n")
    println("Playing song \"" + songName + "\"...")
    defer vox.Quit()
    song, err := vox.Open(songPath)
    errTest(err)
    song.SetLooping(false)
    defer song.Close()
    song.Play()
    for !song.Finished() {
        println("Line " + strconv.Itoa(song.Line()) + ", Volume " + strconv.Itoa(song.Volume()) + "%")
        song.Pause()
        time.Sleep(time.Duration(studderTime) * time.Millisecond)
        song.Play()
        time.Sleep(time.Duration(studderTime) * time.Millisecond)
    }
}

func playInFull(songPath string) {
    songName := filepath.Base(songPath)
    errTest(vox.Init("", 44100, 2, 0))
    println("Vox!\nv" + vox.Version + "\n")
    println("Playing song \"" + songName + "\"...")
    defer vox.Quit()
    song, err := vox.Open(songPath)
    errTest(err)
    song.SetLooping(false)
    defer song.Close()
    song.Play()
    for !song.Finished() { } //TODO: Doesn't ever stop...
}

func playSprite(songPath string, piece int) {
    songName := filepath.Base(songPath)
    errTest(vox.Init("", 44100, 2, 0))
    println("Vox!\nv" + vox.Version + "\n")
    println("Playing song \"" + songName + "\"...")
    defer vox.Quit()
    song, err := vox.Open(songPath)
    errTest(err)
    song.SetLooping(false)
    pieceLength := song.Lines() / 8 // Eight pieces per sprite
    defer song.Close()
    song.Seek(piece * pieceLength, 0)
    song.Play()
    for song.Line() < ((piece + 1) * pieceLength - 1) {}
    time.Sleep(time.Duration(getSongLength(*song)) * time.Second)
    song.Pause()
}

func main() {
    //playInFull("examples/simple_examples/fm.sunvox")
    //studder("examples/kostya_m - Midnight.sunvox", 123)
    //playInFull("examples/kostya_m - Midnight.sunvox")
    //playInFull("sprite1.sunvox")
    playSprite("sprite1.sunvox", 4)
    playSprite("sprite1.sunvox", 4)
    playSprite("sprite1.sunvox", 5)
    playSprite("sprite1.sunvox", 7) // TODO: Below this line won't play!
    playSprite("sprite1.sunvox", 7)
    playSprite("sprite1.sunvox", 5)
    playSprite("sprite1.sunvox", 4)
    playSprite("sprite1.sunvox", 2)
    playSprite("sprite1.sunvox", 0)
    playSprite("sprite1.sunvox", 0)
    playSprite("sprite1.sunvox", 2)
    playSprite("sprite1.sunvox", 4)
    playSprite("sprite1.sunvox", 4)
    playSprite("sprite1.sunvox", 2)
    playSprite("sprite1.sunvox", 2)
}

