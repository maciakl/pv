package main

import (
    "os"
    "fmt"
    "flag"
    "strings"
    "os/exec"
    "path/filepath"
)

const version = "0.1.0"

const (

    imageviewer = "chafa"
    imageviewerargs = ""

    textviewer = "bat"
    textviewerargs = ""

    mdviewer = "glow"
    mdviewerargs = "--style=dracula"

    pdfviewer = "pdftotext"
    pdfviewerargs = "-"

    musicviewer = "exiftool"
    musicviewerargs = ""

    videoviewer = "exiftool"
    videoviewerargs = ""

    zipviewer = "unzip"
    zipviewerargs = "-l"

    sevenzviewer = "7z"
    sevenzviewerargs = "l"

    rarviewer = "unrar"
    rarviewerargs = "l"

    wordviewer = "pandoc"
    wordviewerargs = "--to=plain"

    webviewer = "lynx"
    webviewerargs = "-dump"

    defaultviewer = "bat"
    defaultviewerargs = "--color=always"
)



func main() {

    var ver bool
    flag.BoolVar(&ver, "version", false, "display version number and exit")
    flag.Parse()

    // show version and exit
    if ver {
        fmt.Println(filepath.Base(os.Args[0]), "version", version)
        os.Exit(0)
    }

    // check if a file path was provided
    if len(flag.Args()) == 0 {
        fmt.Println("No file path provided")
        os.Exit(1)
    }

    path := flag.Arg(0)

    filetype := getFileType(path)

    var cmd *exec.Cmd

    switch filetype {

        case "Image":
            cmd = exec.Command(imageviewer, imageviewerargs, path)

        case "Markdown":
            cmd = exec.Command(mdviewer, mdviewerargs, path)

        case "PDF":
            cmd = exec.Command(pdfviewer, path, pdfviewerargs)

        case "Music":
            cmd = exec.Command(musicviewer, musicviewerargs, path)

        case "Video":
            cmd = exec.Command(videoviewer, musicviewerargs, path)

        case "Zip":
            cmd = exec.Command(zipviewer, zipviewerargs, path)

        case "7z":
            cmd = exec.Command(sevenzviewer, sevenzviewerargs, path)

        case "RAR":
            cmd = exec.Command(rarviewer, rarviewerargs, path)

        case "Word":
            cmd = exec.Command(wordviewer, wordviewerargs, path)

        case "HTML":
            cmd = exec.Command(webviewer, webviewerargs, path)

        // bat is a good default for everything else
        // it will just display [binary file] for unknown types
        default:
            cmd = exec.Command(defaultviewer, defaultviewerargs, path)

    }

        cmd.Stdout = os.Stdout
        cmd.Stderr = os.Stderr
        cmd.Run()
}


// function that takes in a file path and determines the type of file
func getFileType(path string) string {
    // check if file exists
    if _, err := os.Stat(path); os.IsNotExist(err) {
        return "File does not exist"
    }

    // get the file extension
    ext := filepath.Ext(path)
    ext = strings.ToLower(ext)

    switch ext {
        
        // image files
        case ".jpg", ".jpeg", ".png", ".gif", ".bmp", ".tiff", ".svg", ".webp", ".ico", ".jit", ".jif", ".jfi":
            return "Image"

        case    ".mp3", ".flac", ".wav", ".ogg", ".m4a", ".wma", ".aac", ".aiff", ".alac", ".ape", ".dsd", ".dts", 
                ".mka", ".mpc", ".ofr", ".ofs", ".opus", ".tak", ".tta", ".wv":
            return "Music"

        // video files
        case    ".mp4", ".mkv", ".webm", ".avi", ".mov", ".wmv", ".flv", ".3gp", ".mpg", ".mpeg", ".m2v", ".m4v", 
                ".m2ts", ".ts", ".mts", ".vob", ".divx", ".xvid", ".rm", ".rmvb", ".asf", ".ogv", ".3g2", ".f4v", 
                ".h264", ".h265", ".hevc", ".vp9", ".vp8", ".av1":
            return "Video"

        // markdown files
        case ".md", ".markdown":
            return "Markdown"

        // zip files
        case ".zip", ".jar":
            return "Zip"

        case ".7z":
            return "7z"

        case ".rar":
            return "RAR"

        case ".pdf":
            return "PDF"

        case ".docx", ".odt", ".rtf":
            return "Word"

        case ".html", ".htm":
            return "HTML"

        default:
            return "Other"

    }





}
