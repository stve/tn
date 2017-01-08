# tn

An experiment in writing a golang CLI to replace some of my [ruby scripts](https://github.com/stve/bin) for managing/manipulating mp3s.

*Note:* These scripts are fairly customized to my use-case. This is also my first time writing Go so I suspect the code is all sorts of busted.

## Usage

```go
tn (command)
```

## Commands

### Artwork

Set the cover image based on filenames. Running the `artwork` command will update the cover art if an image is found with a matching filename.

```go
tn artwork
```

Basically, it assumes your files look like this:

```sh
ls ~/Downloads/

Owsey - And Then I Woke Up.jpg
Owsey - And Then I Woke Up.mp3
```

### Clear

Removes all existing id3 tags from all mp3s in the current directory.

```go
tn clear
```

### Cover

Set the cover image for a group of mp3 files in a directory. Assumes an image `cover.jpg` exists in the directory in which you are running the command.

```go
tn cover
```

### iTunes

Load all mp3s in the current directory into iTunes.

```go
tn itunes
```

### Set/Update Tags

Set artist name:

```go
tn tag --artist "Owsey"
```

Set the album name:

```go
tn tag --album "To The Child Drifting Out At Sea"
```

Set the title:

```go
tn tag --title "I've Lost All Light In My Life"
```

By default, these commands will set on all mp3s in a given directory. For artist and album, this is probably desired but for song titles you probably will want to specify the file to update:

```go
tn tag --title "I've Lost All Light In My Life" --file song.mp3
```

Multiple flags can be passed at once:

```go
tn tag --artist "Owsey" --album "To The Child Drifting Out At Sea"
```

#### Autotagging

Autotagging is a bit of a unique use-case. If your files are named in the format `<artist name> - <song title>.mp3`, you can use the `--auto` flag to set the artist and title for all files in a directory using:

```go
tn tag --auto
```

### View Tags

View the current tags on mp3s in the current directory:

```go
tn tags
```

If you'd like to see if there's an image, you can pass the `--artwork` flag:

```go
tn tags --artwork
```

## Copyright

Copyright (c) 2016 Steve Agalloco.
