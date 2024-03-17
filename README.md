# work-documentation


## Usage

```
wodo (work documentation) is tool to simply, fast and affortless documentation your work

Usage:
  wodo [describe] [your] [task] ... [flags]
  wodo [command]

Available Commands:
  add         Add a task done today
  completion  Generate the autocompletion script for the specified shell
  config      Create / manipulate config
  dummy       Creates some dummy data
  help        Help about any command
  list        List to current day
  sum         Lists to current month
  version     Print the version number of wodo

Flags:
  -h, --help   help for wodo

Use "wodo [command] --help" for more information about a command.
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

