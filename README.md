# Previewer for the LF CLI File Manager

This is a simple previewer for the [LF File Manager](https://github.com/gokcehan/lf). 

It detects file type based on the file extension and then uses the appropriate previewer to display the file in the terminal.

Sample preview of it working in tandem with `lf`:

![pv](https://raw.githubusercontent.com/maciakl/pv/refs/heads/main/screenshot.gif)

Note that this tool does not do any actual previewing by itself. It is just a wrapper around other previewers.

## Usage:

You can run `pv` on it's own:

    pv filename

This will simply output the preview of the file to the terminal:

![scr2](https://github.com/user-attachments/assets/2c50f80a-2095-414b-9248-cf9585f836b1)

The program supports a couple of basic command line switches:

    Usage: pv.exe <file>
    Options:
        -v, --version   Show version
        -h, --help      Show this help
        -c, --config    Show configuration

While it works as a stand alone program, `pv` has been designed specifically to work in tandem with `lf` in lieu of a shell specific script.

To configure `pv` to be the default previewer for `lf` add the following line to your `lfrc` config file:

    set previewer pv

## Configuration

You can configure the previewers and their arguments by creating a config file named `.pvrc` 

- On Linux, Unix and Mac put it in `~/.pvrc`
- On Windows put it in `%USERPROFILE%\.pvrc`

The format of the config file is simple:

    key=value

No spaces!

Here is an example:
    
    text_viewer=bat
    text_viewer_opts=--color=always
    text_viewer_args=--theme=dracula
    
Each `_viewer` has a corresponding `_args` and `_opts` setting, both of which are optional and only need to be specified if you want to override the defaults. Both are used to pass additional arguments to the viewer executable. The difference is as follows:

- The `_opts` are arguments or subcommands that are passed in before the file path
- The `_args` are trailing arguments that are passed in after the file path

For example:

    default_viewer default_viewer_opts <file> default_viewer_args

This is important, because some programs have positional arguments for input and output files.

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
|`text_viewer`| `txt`|
|`log_viewer`| `log`|
|`naked_viewer`| Used when a file has no extension|
|`default_viewer`| Used when the file extension does not match any of the above|

Currently there is no way to change the file extension association via the configuration file.

The `_viewer` option needs to be set to a string that represents a file name of an actual executable (or an executable script) file. This file:

- must be in your `$PATH` (or `%PATH%`)
- must output directly to `stdout`

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

Using `bat` as the `default_viewer` is recommended as it usually works well with most file types (even binary ones).

The `naked_viewer` option is set to `bat` by default as files with no extension are just as likely to be binary executables as they are to be shell scripts.

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
