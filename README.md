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

You can configure the previewers and their arguments by creating a `~/.pvrc` file (it's `%USERPROFILE%\.pvrc` on Windows).

The format is simple:

    key=value

No spaces!

Here is an example:
    
    text_viewer=bat
    text_viewer_opts=--color=always
    text_viewer_args=--theme=dracula
    
Each `viewer` has a corresponding `_args` and `_opts` setting, both of which are optional and only need to be specified if you want to override the defaults. Both are used to pass additional arguments to the viewer executable. The difference is as follows:

- The `_opts` arguments are passed before the file path
- The `_args` arguments are passed after the file path

You can check the current configuration by running:

    pv -c

This will output the current configuration and corresponding opt and arg values.

### Viewers

Following viewers are available to be overriden. Each one is listed alongside the the extensions that it is associated with:

|Viewer|Extensions|
|---|---|
|`image_viewer`| `jpg` `jpeg` `png` `gif` `bmp` `tiff` `svg` `webp` `ico` `jit` `jif` `jfi`|
|`md_viewer`| `md` `markdown` `mkd` `mkdn` `mdwn` `mdown` `mdtxt`|
|`pdf_viewer`| `pdf`|
|`music_viewer`|`mp3` `flac` `wav` `ogg` `m4a` `wma` `aac` `aiff` `alac` `ape` `dsd` `dts` `mka` `mpc` `ofr` `.ofs` `opus` `tak` `tta` `wv`|
|`video_viewer`|`mp4` `mkv` `webm` `avi` `mov` `wmv` `flv` `3gp` `mpg` `mpeg` `m2v` `m4v` `m2ts` `ts` `mts` `vob` `divx` `xvid` `rm` `rmvb` `asf` `ogv` `3g2` `f4v` `h264` `h265` `hevc` `vp9` `vp8` `av1`|
|`zip_viewer`|`zip` `jar`|
|`sevenz_viewer`| `7z`|
|`rar_viewer`| `rar`|
|`word_viewer`| `doc` `docx` `odt` `rtf`|
|`excel_viewer`| `xlsx` |
|`web_viewer`| `html` `htm` `xhtml` `mhtml` `mht`|
|`exe_viewer`|`exe` `dll` `msi` `sys` `msx`|
|`text_viewer| `txt`|
|`default_viewer`| Used when the file extension does not match any of the above|

Currently there is no way to change the file extension association via the configuration file.

The "viewer" option need to be set to a string that represents a file name of an actual executables or executable script that is in your PATH and that outpus their results to stdout.

### Defaults

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

### Overriding Default Arguments

To override the default `_opt` or `_arg` value for a specific viewer, simply provide a new one in the config file. To override it with nothing, put nothing after the `=` sign like this:

    text_viewer_opts=
    text_viewer_args=

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

First get [scoop](https://scoop.sh/). 

Then add my bucket and install `pv`:
    
    scoop bucket add maciak https://github.com/maciakl/bucket
    scoop update
    scoop install pv
    
To install all the default previewers do:

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
