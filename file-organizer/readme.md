# File Organizer

A simple Go CLI tool to **organize files in a directory** by their type (images, audios, documents, or unknown). It recursively sorts files into category directories and supports a **dry-run mode** to simulate moves without making changes. This program is created for purpose of learning go

---

## Features

* Automatically sorts files by extension into category folders:

  * **Images:** `.png`, `.jpg`, `.jpeg`
  * **Audios:** `.mp3`, `.wav`
  * **Documents:** `.txt`, `.docx`
  * **Unknown:** any other file types
* Creates category directories if they do not exist.
* Recursive handling of subdirectories (skips category directories).
* Optional **dry-run mode** to preview moves without modifying files.
* Informative logging with levels: `Info`, `Debug`, `Error`.

---

## Installation

1. Make sure you have [Go installed](https://golang.org/doc/install).
2. Clone this repository:

```bash
git clone <repo-url>
cd file-organizer
```

3. Build the CLI:

```bash
go build -o file-organizer main.go
```

---

## Usage

### Basic Usage

```bash
./file-organizer --folder "C:\Users\test\Desktop\test-folder"
```

* `--folder` (required) — Path to the directory to organize. Can be relative or absolute.
* `--dry-run` (optional) — Simulate file moves without changing files.

Example with dry-run:

```bash
./file-organizer --folder "C:\Users\test\Desktop\test-folder" --dry-run
```

---

### Behavior

1. The program scans the specified folder.
2. For each file:

   * Determines its type by extension.
   * Moves it into the respective category directory (`images`, `audios`, `documents`, or `unknown`).
3. For subdirectories:

   * Skips category directories.
   * Recursively organizes files in other subdirectories.

---

### Logging

* **Info:** Reports high-level actions (e.g., entering directories, moving files).
* **Debug:** Provides verbose internal details (shown if verbose mode is enabled in logger).
* **Error:** Reports any file operation or directory creation errors.

Example dry-run output:

```
[DRY RUN] entered directory named: C:\Users\test\Desktop\test-folder
[DRY RUN] would move test-folder\image1.png -> test-folder\images\image1.png
[DRY RUN] would move test-folder\document1.docx -> test-folder\documents\document1.docx
```

---

## Directory Structure After Organizing

```
test-folder/
├─ images/
│  ├─ image1.png
│  └─ screenshot.jpg
├─ audios/
│  └─ song.mp3
├─ documents/
│  └─ notes.docx
└─ unknown/
   └─ file.xyz
```

---

## Extending

* Add more extensions to `extToDir` map in `main.go`.
* Update category names in `dirMap` if you want additional categories.

---