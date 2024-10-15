# Previewer for LF

This is a simple previewer for lf. It detects file type and then uses the appropriate previewer to display the file on the cli.

![pv](https://raw.githubusercontent.com/maciakl/pv/refs/heads/main/screenshot.gif)

Note that this tool does not do any actual previewing by itself. It is just a wrapper around other previewers.

## Usage:

You can run `pv` on it's own:

    pv filename

Or you can use it as a previewer in lf. Add this to your `lfrc`:

    set previewer pv


## Configuration

You can configure the previewers and their arguments by creating a `~/.pvrc` file (it's `%USERPROFILE%\.pvrc` on Windoes). 

Here is an example:
    
    imageviewer = chafa
    mdviewer = glow
    mdviewerargs = --style=dark
    wordviewer = pandoc
    wordviewerargs = --to=ansi
    webviewer = w3m
    webviewerargs = -dump
    exeviewer = file
    defaultviewer = less

You only have specify the "args" if you want to pass additional arguments to the previewer. If you don't specify a viewer for a given file type, the default viewer will be used.

You can check the current configuration by running:

    pv -version

It will display the program version as well as the current configuration.

### Configuration options

The "viewer" options need to be actual executables or executable scripts that output their results to stdout. The "args" options are optional and can be any arguments you want to pass to the viewer.


- `imageviewer` - used with the following file extensions: ".jpg", ".jpeg", ".png", ".gif", ".bmp", ".tiff", ".svg", ".webp", ".ico", ".jit", ".jif", ".jfi"
- `imageviewerargs` -  optional args for the image viewer

- `mdviewer` - used with the following file extensions: ".md", ".markdown", ".mkd", ".mkdn", ".mdwn", ".mdown", ".mdtxt"
- `mdviewerargs` - optional args for the markdown viewer

- `pdfviewer` - used with the following file extensions: ".pdf"
- `pdfviewerargs` - optional args for the pdf viewer

- `musicviewer` - used with the following file extensions: ".mp3", ".flac", ".wav", ".ogg", ".m4a", ".wma", ".aac", ".aiff", ".alac", ".ape", ".dsd", ".dts", ".mka", ".mpc", ".ofr", ".ofs", ".opus", ".tak", ".tta", ".wv"
- `musicviewerargs` - optional args for the music viewer

- `videoviewer` - used with the following file extensions: ".mp4", ".mkv", ".webm", ".avi", ".mov", ".wmv", ".flv", ".3gp", ".mpg", ".mpeg", ".m2v", ".m4v", ".m2ts", ".ts", ".mts", ".vob", ".divx", ".xvid", ".rm", ".rmvb", ".asf", ".ogv", ".3g2", ".f4v", ".h264", ".h265", ".hevc", ".vp9", ".vp8", ".av1"
- `videoviewerargs` = ""

- `zipviewer` - used with the following file extensions: ".zip", ".jar"
- `zipviewerargs` - optional args for the zip viewer

- `sevenzviewer` - used with the following file extensions: ".7z"
- `sevenzviewerargs` - optional args for the 7z viewer

- `rarviewer` - used with the following file extensions: ".rar"
- `rarviewerargs` - optional args for the rar viewer

- `wordviewer` - used with the following file extensions: ".doc", ".docx", ".odt", ".rtf"
- `wordviewerargs` - optional args for the word viewer

- `excelviewer` - used with the following file extensions: ".xlsx"
- `excelviewerargs` - optional args for the excel viewer

- `webviewer` - used with the following file extensions: ".html", ".htm", ".xhtml", ".mhtml", ".mht"
- `webviewerargs` - optional args for the web viewer

- `exeviewer` - used with the following file extensions: ".exe", ".dll", ".msi", ".sys", ".msx"
- `exeviewerargs` - optional args for the exe viewer

- `defaultviewer` - used when no other viewer is specified
- `defaultviewerargs` - optional args for the default viewer

Using `bat` as the default viewer is recommended as it usually works well with most file types.

Current defaults are:

        imageviewer:     chafa
        textviewer:      bat
        mdviewer:        glow --style=dracula
        pdfviewer:       pdftotext -
        musicviewer:     exiftool
        videoviewer:     exiftool
        zipviewer:       unzip -l
        sevenzviewer:    7z l
        rarviewer:       unrar l
        wordviewer:      pandoc --to=plain
        excelviewer:     xlsx2csv -o=-
        webviewer:       lynx -dump
        exeviewer:       exiftool
        defaultviewer:   bat --color=always

## Depedencies

You need a previewer for each file type.

The default previewers are:

- [bat](https://github.com/sharkdp/bat) (default viewer)
- [chafa](https://hpjansson.org/chafa/) (for images)
- [glow](https://github.com/charmbracelet/glow) (for markdown files)
- [exiftool](https://exiftool.org/) (for audio and video files)
- [poppler](https://poppler.freedesktop.org/) (for pdf files)
- [pandoc](https://pandoc.org/) (for word and rtf files)
- [lynx](https://lynx.invisible-island.net/) (for html files)
- For archive files you need: `unzip`, `7z`, `unrar`



## Installation

### Via Go

Install via go:
 
    go install github.com/maciakl/pv@latest

### On Windows

First get [scoop](https://scoop.sh/). Then install all the prerequisites with:

    scoop install lf
    scoop install bat
    scoop install chafa
    scoop install glow
    scoop install exiftool
    scoop install poppler
    scoop install pandoc
    scoop install lynx
    scoop install unzip
    scoop install 7z
    scoop install unrar

Then add my bucket and install `pv`:
    
    scoop bucket add maciak https://github.com/maciakl/bucket
    scoop update
    scoop install pv
    
