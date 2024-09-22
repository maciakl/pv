package main

import (
    "os"
    "fmt"
    "flag"
    "bufio"
    "strings"
    "os/exec"
    "path/filepath"
)

const version = "0.2.1"

var (

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

    excelviewer = "xlsx2csv"
    excelviewerargs = "-o=-"

    webviewer = "lynx"
    webviewerargs = "-dump"

    exeviewer = "exiftool"
    exeviewerargs = ""

    defaultviewer = "bat"
    defaultviewerargs = "--color=always"
)



func main() {

    var ver bool
    flag.BoolVar(&ver, "version", false, "display version number and exit")
    flag.Parse()

    readConfig()

    // show version and exit
    if ver {
        fmt.Println(filepath.Base(os.Args[0]), "version", version)
        home,_ := os.UserHomeDir()
        fmt.Println("Config file:", filepath.Join(home, ".pvrc"))

        fmt.Println("\nConfiguration:")
        fmt.Println("\timageviewer:\t", imageviewer, imageviewerargs)
        fmt.Println("\ttextviewer:\t", textviewer, textviewerargs)
        fmt.Println("\tmdviewer:\t", mdviewer, mdviewerargs)
        fmt.Println("\tpdfviewer:\t", pdfviewer, pdfviewerargs)
        fmt.Println("\tmusicviewer:\t", musicviewer, musicviewerargs)
        fmt.Println("\tvideoviewer:\t", videoviewer, videoviewerargs)
        fmt.Println("\tzipviewer:\t", zipviewer, zipviewerargs)
        fmt.Println("\tsevenzviewer:\t", sevenzviewer, sevenzviewerargs)
        fmt.Println("\trarviewer:\t", rarviewer, rarviewerargs)
        fmt.Println("\twordviewer:\t", wordviewer, wordviewerargs)
        fmt.Println("\texcelviewer:\t", excelviewer, excelviewerargs)
        fmt.Println("\twebviewer:\t", webviewer, webviewerargs)
        fmt.Println("\texeviewer:\t", webviewer, exeviewerargs)
        fmt.Println("\tdefaultviewer:\t", defaultviewer, defaultviewerargs)
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

        case "Excel":
            cmd = exec.Command(excelviewer, excelviewerargs, path)

        case "HTML":
            cmd = exec.Command(webviewer, webviewerargs, path)

        case "EXE":
            cmd = exec.Command(exeviewer, exeviewerargs, path)

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
        case ".md", ".markdown", ".mkd", ".mkdn", ".mdown", ".mdwn", ".mdtxt", ".mdtext":
            return "Markdown"

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

        case ".xlsx":
            return "Excel"

        case ".exe", ".msi", ".msx", ".dll":
            return "EXE"

        case ".html", ".htm", ".xhtml", ".mhtml", ".mht":
            return "HTML"

        default:
            return "Other"

    }
}


// read in the ~/.pvrc file line by line and set the variables
func readConfig() {
    home,_ := os.UserHomeDir()
    file, err := os.Open(filepath.Join(home, ".pvrc"))
    if err != nil {
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {

        line := scanner.Text()

        if strings.HasPrefix(line, "#") {
            continue
        }

        parts := strings.SplitN(line, "=", 2)
        if len(parts) != 2 {
            continue
        }

        key := strings.TrimSpace(parts[0])
        value := strings.TrimSpace(parts[1])

        switch key {
            case "imageviewer":
                imageviewer = value
            case "imageviewerargs":
                imageviewerargs = value
            case "textviewer":
                textviewer = value
            case "textviewerargs":
                textviewerargs = value
            case "mdviewer":
                mdviewer = value
            case "mdviewerargs":
                mdviewerargs = value
            case "pdfviewer":
                pdfviewer = value
            case "pdfviewerargs":
                pdfviewerargs = value
            case "musicviewer":
                musicviewer = value
            case "musicviewerargs":
                musicviewerargs = value
            case "videoviewer":
                videoviewer = value
            case "videoviewerargs":
                videoviewerargs = value
            case "zipviewer":
                zipviewer = value
            case "zipviewerargs":
                zipviewerargs = value
            case "sevenzviewer":
                sevenzviewer = value
            case "sevenzviewerargs":
                sevenzviewerargs = value
            case "rarviewer":
                rarviewer = value
            case "rarviewerargs":
                rarviewerargs = value
            case "wordviewer":
                wordviewer = value
            case "wordviewerargs":
                wordviewerargs = value
            case "excelviewer":
                excelviewer = value
            case "excelviewerargs":
                excelviewerargs = value
            case "webviewer":
                webviewer = value
            case "webviewerargs":
                webviewerargs = value
            case "exeviewer":
                exeviewer = value
            case "exeviewerargs":
                exeviewerargs = value
            case "defaultviewer":
                defaultviewer = value
            case "defaultviewerargs":
                defaultviewerargs = value
        }
    }
}
