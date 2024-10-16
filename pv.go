package main

import (
    "os"
    "fmt"
    "bufio"
    "strings"
    "os/exec"
    "path/filepath"
)

const version = "0.2.2"

var (

    image_viewer = "chafa"
    image_viewer_opts = ""
    image_viewer_args = ""

    text_viewer = "bat"
    text_viewer_opts = "--color=always"
    text_viewer_args = "--theme=dracula"

    md_viewer = "glow"
    md_viewer_opts = "--style=dracula"
    md_viewer_args = ""

    pdf_viewer = "pdftotext"
    pdf_viewer_opts = ""
    pdf_viewer_args = "-"

    music_viewer = "exiftool"
    music_viewer_opts = ""
    music_viewer_args = ""

    video_viewer = "exiftool"
    video_viewer_opts = ""
    video_viewer_args = ""

    zip_viewer = "unzip"
    zip_viewer_opts = "-l"
    zip_viewer_args = ""

    sevenz_viewer = "7z"
    sevenz_viewer_opts = "l"
    sevenz_viewer_args = ""

    rar_viewer = "unrar"
    rar_viewer_opts = "l"
    rar_viewer_args = ""

    word_viewer = "pandoc"
    word_viewer_opts = "--to=plain"
    word_viewer_args = ""

    excel_viewer = "xlsx2csv"
    excel_viewer_opts = "-o=-"
    excel_viewer_args = ""

    web_viewer = "lynx"
    web_viewer_opts = "-dump"
    web_viewer_args = ""

    exe_viewer = "hexyl"
    exe_viewer_opts = ""
    exe_viewer_args = ""

    default_viewer = "bat"
    default_viewer_opts = "--color=always"
    default_viewer_args = ""
)



func main() {

    readConfig()

    // check id there are arguments
    if len(os.Args) > 1 {
        switch os.Args[1] {

        case "-v", "--version":
            showVersion()
            os.Exit(0)

        case "-h", "--help":
            showUsage()
            os.Exit(0)

        case "-c", "--config":
            showConfig()
            os.Exit(0)

        default:
            path := os.Args[1]
            openFile(path)
        }

    } else {
        showUsage()
        os.Exit(1)
    }

}


func openFile(path string) {

    filetype := getFileType(path)

    var cmd *exec.Cmd

    var viewer, viewer_opts, viewer_args string

    switch filetype {

        case "Image":
            viewer = image_viewer
            viewer_opts = image_viewer_opts
            viewer_args = image_viewer_args

        case "Markdown":
            viewer = md_viewer
            viewer_opts = md_viewer_opts
            viewer_args = md_viewer_args

        case "PDF":
            viewer = pdf_viewer
            viewer_opts = pdf_viewer_opts
            viewer_args = pdf_viewer_args

        case "Music":
            viewer = music_viewer
            viewer_opts = music_viewer_opts
            viewer_args = music_viewer_args

        case "Video":
            viewer = video_viewer
            viewer_opts = video_viewer_opts
            viewer_args = video_viewer_args

        case "Zip":
            viewer = zip_viewer
            viewer_opts = zip_viewer_opts
            viewer_args = zip_viewer_args

        case "7z":
            viewer = sevenz_viewer
            viewer_opts = sevenz_viewer_opts
            viewer_args = sevenz_viewer_args

        case "RAR":
            viewer = rar_viewer
            viewer_opts = rar_viewer_opts
            viewer_args = rar_viewer_args

        case "Word":
            viewer = word_viewer
            viewer_opts = word_viewer_opts
            viewer_args = word_viewer_args

        case "Excel":
            viewer = excel_viewer
            viewer_opts = excel_viewer_opts
            viewer_args = excel_viewer_args

        case "HTML":
            viewer = web_viewer
            viewer_opts = web_viewer_opts
            viewer_args = web_viewer_args


        case "EXE":
            viewer = exe_viewer
            viewer_opts = exe_viewer_opts
            viewer_args = exe_viewer_args


        case "Text":
            viewer = text_viewer
            viewer_opts = text_viewer_opts
            viewer_args = text_viewer_args

        default:
            viewer = default_viewer
            viewer_opts = default_viewer_opts
            viewer_args = default_viewer_args
    }


    if viewer_opts != "" && viewer_args != "" {
        cmd = exec.Command(viewer, viewer_opts, path, viewer_args)
    } else if viewer_opts != "" && viewer_args == "" {
        cmd = exec.Command(viewer, viewer_opts, path)
    } else if viewer_opts == "" && viewer_args != "" {
        cmd = exec.Command(viewer, path, viewer_args)
    } else {
        cmd = exec.Command(viewer, path)
    }
        
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    cmd.Run()
}


// function that shows the usage of the program
func showUsage() {
    fmt.Println("Usage:", filepath.Base(os.Args[0]), "<file>")
    fmt.Println("Options:")
    fmt.Println("\t-v, --version\tShow version")
    fmt.Println("\t-h, --help\tShow this help")
    fmt.Println("\t-c, --config\tShow configuration")
}

func showVersion() {
    fmt.Println(filepath.Base(os.Args[0]), "version", version)
    os.Exit(0)
}

func showConfig() {
    fmt.Println(filepath.Base(os.Args[0]), "version", version)
    home,_ := os.UserHomeDir()
    fmt.Println("Config file:", filepath.Join(home, ".pvrc"))

    fmt.Println("\nConfiguration:")
    fmt.Println("\timage_viewer:\t", image_viewer, image_viewer_opts, "<file>", image_viewer_args)
    fmt.Println("\ttext_viewer:\t", text_viewer, text_viewer_opts, "<file>", text_viewer_args)
    fmt.Println("\tmd_viewer:\t", md_viewer, md_viewer_opts, "<file>", md_viewer_args)
    fmt.Println("\tpdf_viewer:\t", pdf_viewer, pdf_viewer_opts, "<file>", pdf_viewer_args)
    fmt.Println("\tmusic_viewer:\t", music_viewer, music_viewer_opts, "<file>", music_viewer_args)
    fmt.Println("\tvideo_viewer:\t", video_viewer, video_viewer_opts, "<file>", video_viewer_args)
    fmt.Println("\tzip_viewer:\t", zip_viewer, zip_viewer_opts, "<file>", zip_viewer_args)
    fmt.Println("\tsevenz_viewer:\t", sevenz_viewer, sevenz_viewer_opts, "<file>", sevenz_viewer_args)
    fmt.Println("\trar_viewer:\t", rar_viewer, rar_viewer_opts, "<file>", rar_viewer_args)
    fmt.Println("\tword_viewer:\t", word_viewer, word_viewer_opts, "<file>", word_viewer_args)
    fmt.Println("\texcel_viewer:\t", excel_viewer, excel_viewer_opts, "<file>", excel_viewer_args)
    fmt.Println("\tweb_viewer:\t", web_viewer, web_viewer_opts, "<file>", web_viewer_args)
    fmt.Println("\texe_viewer:\t", exe_viewer, exe_viewer_opts, "<file>", exe_viewer_args)
    fmt.Println("\tdefault_viewer:\t", default_viewer, default_viewer_opts, "<file>", default_viewer_args)
    os.Exit(0)
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

        case ".txt":
            return "Text"

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
            case "image_viewer":
                image_viewer = value
            case "image_viewer_opts":
                image_viewer_opts = value
            case "image_viewer_args":
                image_viewer_args = value
            case "text_viewer":
                text_viewer = value
            case "text_viewer_opts":
                text_viewer_opts = value
            case "text_viewer_args":
                text_viewer_args = value
            case "md_viewer":
                md_viewer = value
            case "md_viewer_opts":
                md_viewer_opts = value
            case "md_viewer_args":
                md_viewer_args = value
            case "pdf_viewer":
                pdf_viewer = value
            case "pdf_viewer_opts":
                pdf_viewer_opts = value
            case "pdf_viewer_args":
                pdf_viewer_args = value
            case "music_viewer":
                music_viewer = value
            case "music_viewer_opts":
                music_viewer_opts = value
            case "music_viewer_args":
                music_viewer_args = value
            case "video_viewer":
                video_viewer = value
            case "video_viewer_opts":
                video_viewer_opts = value
            case "video_viewer_args":
                video_viewer_args = value
            case "zip_viewer":
                zip_viewer = value
            case "zip_viewer_opts":
                zip_viewer_opts = value
            case "zip_viewer_args":
                zip_viewer_args = value
            case "sevenz_viewer":
                sevenz_viewer = value
            case "sevenz_viewer_opts":
                sevenz_viewer_opts = value
            case "sevenz_viewer_args":
                sevenz_viewer_args = value
            case "rar_viewer":
                rar_viewer = value
            case "rar_viewer_opts":
                rar_viewer_opts = value
            case "rar_viewer_args":
                rar_viewer_args = value
            case "word_viewer":
                word_viewer = value
            case "word_viewer_opts":
                word_viewer_opts = value
            case "word_viewer_args":
                word_viewer_args = value
            case "excel_viewer":
                excel_viewer = value
            case "excel_viewer_opts":
                excel_viewer_opts = value
            case "excelviewerargs":
                excel_viewer_args = value
            case "web_viewer":
                web_viewer = value
            case "web_viewer_opts":
                web_viewer_opts = value
            case "web_viewer_args":
                web_viewer_args = value
            case "exe_viewer":
                exe_viewer = value
            case "exe_viewer_opts":
                exe_viewer_opts = value
            case "exe_viewer_args":
                exe_viewer_args = value
            case "default_viewer":
                default_viewer = value
            case "default_viewer_opts":
                default_viewer_opts = value
            case "default_viewer_args":
                default_viewer_args = value
        }
    }
}
