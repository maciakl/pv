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
    textviewer = bat

    mdviewer = glow
    mdviewerargs = --style=dracula

    pdfviewer = pdftotext
    pdfviewerargs = -

    musicviewer = exiftool
    videoviewer = exiftool

    zipviewer = unzip
    zipviewerargs = -l

    sevenzviewer = 7z
    sevenzviewerargs = l

    rarviewer = unrar
    rarviewerargs = l

    wordviewer = pandoc
    wordviewerargs = --to=plain

    webviewer = lynx
    webviewerargs = -dump

    defaultviewer = bat
    defaultviewerargs = --color=always

You only have specify the "args" if you want to pass additional arguments to the previewer. If you don't specify a viewer for a given file type, the default viewer will be used.

You can check the current configuration by running:

    pv -version

It will display the program version as well as the current configuration.

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



## Installation on Windows

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
    
