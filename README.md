# Se**q**uential Co**p**y

A tool to recursively copy files and directories so that timestamps are in order of file names.

## Features

- Copies directories first, then files
- Maintains sequential timestamps
- Displays progress with counters

## Usage

### Command-Line Mode

To use in command-line mode, simply provide the source and destination paths as arguments:

```
qp <source> <destination>
```

**Example:**
```
qp /path/to/source /path/to/destination
```


### Interactive Mode

If you run  without arguments, it will start in interactive mode:

```
qp
```

In interactive mode:

1. You'll be prompted to enter the source path
2. Then you'll be prompted to enter the destination path
3. After the operation completes, press Enter to exit the program

**Tip for Windows users**:

- Simply double-click the `qp.exe` file to launch the program in interactive mode
- Drag and drop the source folder into the opened window and press Enter
- Drag and drop the destination folder into the opened window and press Enter
- Wait for the operation to complete
- Press Enter or close the window to exit the program

**Example session:**
```
🖥️  Interactive Mode
--------------------
Enter source path: /home/user/music
Enter destination path: /media/sdcard/music

📂 Copying from: /home/user/music
📂 Copying to:   /media/sdcard/music

🔍 Scanning source directory...
📁 Found 3 directories and 12 files.

📂 Copying directories...
[DIR] 1/3 Created: /media/sdcard/music/album1
[DIR] 2/3 Created: /media/sdcard/music/album2
[DIR] 3/3 Created: /media/sdcard/music/album3

📄 Copying files...
[FILE] 1/12 Copied: /media/sdcard/music/album1/song1.flac
[FILE] 2/12 Copied: /media/sdcard/music/album1/song2.flac
...

✅ Copy completed successfully!

Press Enter to exit...
```

## Build

Make sure you have installed [Go](https://go.dev/) and [Just](https://just.systems/).

Execute:

```
just build
```
