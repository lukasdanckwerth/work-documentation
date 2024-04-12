# work-documentation

> wodo (work documentation) is a command line tool to simply, fast and affortless documentate your work

<!-- TOC -->
- [Usage](#usage)
- [Installation](#installation)
- [Dependencies](#dependencies)
<!-- TOC end -->

## Usage

<!--usage begin -->
```
Usage:
  wodo [command] [flags]

Available Commands:
  add         Add a task done today
  config      Create / manipulate config
  help        Help about any command
  list        List to current day
  sum         Lists to current month
  version     Print the version number of wodo

Flags:
  -h, --help   help for wodo

Use "wodo [command] --help" for more information about a command.
```
<!--usage end -->

### Add a task

Adding a task is as easy as.

```bash
$ wodo add "my task" --length 2:30h

# you can omit the quotation marks
$ wodo add my task --length 2:30h

# or even shorter
$ wodo a my task -l 2:30h
```

You can also run `wodo add` without any arguments to get promted with a predefined list of 
__tasks__ and __lengths__ in your config file.

### List tasks of current day

```bash
$ wodo day

# Output:

+---+---------+-------+--------+
| # | TASK    | START | LENGTH |
+---+---------+-------+--------+
| 1 | my task | 22:16 | 02:30  |
+---+---------+-------+--------+
|   |         | TOTAL | 02:30  |
+---+---------+-------+--------+
```

## Installation

Assuming you have go installed, otherwise check the go
[installation guide](https://golang.org/doc/install).

```bash
$ git clone https://github.com/lukasdanckwerth/work-documentation.git

$ cd work-documentation

$ go build -o wodo

$ ./wodo
```

## Dependencies

- [cobra](https://github.com/spf13/cobra)
- [viper](https://github.com/spf13/viper)
- [promptui](https://github.com/manifoldco/promptui)
- [go-pretty](https://github.com/jedib0t/go-pretty/v6)

