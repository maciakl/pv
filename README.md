# Previewer for LF

This is a simple previewer for lf. It detects file type and then uses the appropriate previewer to display the file on the cli.


## Depedencies

You need a previewer for each file type. Right now these hare hard coded, but in the future 
there will be support for a config file.

The default previewers are:

- [bat](https://github.com/sharkdp/bat) (for text files)
- [chafa](https://hpjansson.org/chafa/) (for images)
- [glow](https://github.com/charmbracelet/glow) (for markdown files)
- [exiftool](https://exiftool.org/) (for audio and video files)
- [poppler](https://poppler.freedesktop.org/) (for pdf files)
- [pandoc](https://pandoc.org/) (for word and rtf files)
- [lynx](https://lynx.invisible-island.net/) (for html files)
- For archive files you need: `unzip`, `7z`, `unrar`

## Installation on Windows

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
