# work-documentation


## Usage

<!--usage begin -->
```
wodo (work documentation) is tool to simply, fast and affortless documentation your work

Usage:
  wodo [describe] [your] [task] ... [flags]
  wodo [command]

Available Commands:
  add         Add a task done today
  completion  Generate the autocompletion script for the specified shell
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
wodo add "my task" --length 2:30h

# you can leave out the emphis
wodo add my task --length 2:30h

# or even shorter
wodo a my task -l 2:30h
```

### List tasks of current day

```bash
wodo day

# Output:
+---+---------+-------+--------+
| # | TASK    | START | LENGTH |
+---+---------+-------+--------+
| 1 | my task | 22:16 | 02:30  |
+---+---------+-------+--------+
|   |         | TOTAL | 02:30  |
+---+---------+-------+--------+
```

Build local with.

```bash
go build -o wodo
```

### Add task



## Dependencies

- [cobra](https://github.com/spf13/cobra)
- [viper](https://github.com/spf13/viper)
- [promptui](https://github.com/manifoldco/promptui)
- [go-pretty](https://github.com/jedib0t/go-pretty/v6)

