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

func studder(songPath string, studderTime int) {
    songName := filepath.Base(songPath)
    errTest(vox.Init("", 44100, 2, 0))
    println("Vox!\nv" + vox.Version + "\n")
    println("Playing song \"" + songName + "\"...")
    defer vox.Quit()
    song, err := vox.Open(songPath)
    errTest(err)
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
    defer song.Close()
    song.Play()
    for !song.Finished() { }
}

func main() {
    //playInFull("examples/simple_examples/fm.sunvox")
    studder("examples/kostya_m - Midnight.sunvox", 123)
    playInFull("examples/kostya_m - Midnight.sunvox")
}

