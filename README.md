# Previewer for LF

This is a simple previewer for lf. It detects file type and then uses the appropriate previewer to display the file on the cli.

![pv](https://raw.githubusercontent.com/maciakl/pv/refs/heads/main/screenshot.gif)

Note that this tool does not do any actual previewing by itself. It is just a wrapper around other previewers.

## Usage:

You can run `pv` on it's own:

    pv filename

Or you can use it as a previewer in lf. Add this to your `lfrc`:

    set previewer pv


Usage:

    Usage: pv.exe <file>
    Options:
        -v, --version   Show version
        -h, --help      Show this help
        -c, --config    Show configuration


## Configuration

You can configure the previewers and their arguments by creating a `~/.pvrc` file (it's `%USERPROFILE%\.pvrc` on Windoes). 

Here is an example:
    
    image_viewer=chafa
    md_viewer=glow
    md_viewer_opts = --style=dark
    word_viewer = pandoc
    word_viewer_args = --to=ansi
    web_viewer = w3m
    exe_viewer = file
    default_viewer = less

Each viewer has a corresponding `_args` and `_opts` setting, both of which are optional and only need to be specified if you want to override the defaults. Both are used to pass additional arguments to the viewer executable. The difference is as follows:

- The `_opts` arguments are passed before the file path
- The `_args` arguments are passed after the file path

You can check the current configuration by running:

    pv -c

This will output the current configuration and corresponding opt and arg values.

### Viewers

Following viewers are available:



- `image_viewer` - used with the following file extensions: ".jpg", ".jpeg", ".png", ".gif", ".bmp", ".tiff", ".svg", ".webp", ".ico", ".jit", ".jif", ".jfi"
- `md_viewer` - used with the following file extensions: ".md", ".markdown", ".mkd", ".mkdn", ".mdwn", ".mdown", ".mdtxt"
- `pdf_viewer` - used with the following file extensions: ".pdf"
- `music_viewer` - used with the following file extensions: ".mp3", ".flac", ".wav", ".ogg", ".m4a", ".wma", ".aac", ".aiff", ".alac", ".ape", ".dsd", ".dts", ".mka", ".mpc", ".ofr", ".ofs", ".opus", ".tak", ".tta", ".wv"
- `video_viewer` - used with the following file extensions: ".mp4", ".mkv", ".webm", ".avi", ".mov", ".wmv", ".flv", ".3gp", ".mpg", ".mpeg", ".m2v", ".m4v", ".m2ts", ".ts", ".mts", ".vob", ".divx", ".xvid", ".rm", ".rmvb", ".asf", ".ogv", ".3g2", ".f4v", ".h264", ".h265", ".hevc", ".vp9", ".vp8", ".av1"
- `zip_viewer` - used with the following file extensions: ".zip", ".jar"
- `sevenz_viewer` - used with the following file extensions: ".7z"
- `rar_viewer` - used with the following file extensions: ".rar"
- `word_viewer` - used with the following file extensions: ".doc", ".docx", ".odt", ".rtf"
- `excel_viewer` - used with the following file extensions: ".xlsx"
- `web_viewer` - used with the following file extensions: ".html", ".htm", ".xhtml", ".mhtml", ".mht"
- `exe_viewer` - used with the following file extensions: ".exe", ".dll", ".msi", ".sys", ".msx"
- `text_viewer` - used with the following extensions: ".txt"
- `default_viewer` - used when no other viewer is specified


The "viewer" options need to be actual executables or executable scripts that output their results to stdout.

Current defaults are:

        imageviewer:     chafa <file>
        textviewer:      bat --color=always <file> --theme=dracula
        mdviewer:        glow --style=dracula <file>
        pdfviewer:       pdftotext <file> -
        musicviewer:     exiftool
        videoviewer:     exiftool
        zipviewer:       unzip -l <file>
        sevenzviewer:    7z l <file>
        rarviewer:       unrar l <file>
        wordviewer:      pandoc --to=plain <file>
        excelviewer:     xlsx2csv -o=- <file>
        webviewer:       lynx -dump <file>
        exeviewer:       hyxel <file>
        defaultviewer:   bat --color=always <file>

Using `bat` as the default viewer is recommended as it usually works well with most file types.

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
- [hyxel](https://github.com/sharkdp/hexyl) (for exe files)
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
    scoop install hyxel

Then add my bucket and install `pv`:
    
    scoop bucket add maciak https://github.com/maciakl/bucket
    scoop update
    scoop install pv
    
