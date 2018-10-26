# mongotool
mongotool provides checking database changing and persist the latest database info automatically. You can also diffs two version of data via diff command.

## Get Started



```
Usage:
  mongotool [command]

Available Commands:
  diff        diffs two versions of mongodata data.
  help        Help about any command
  run         Automatically check database changes based on lasest saved version.

Flags:
      --config string      config file (default is $HOME/.mongotool.yaml)
  -H, --host string        host of mongos (default "localhost")
  -P, --port int           port of mongos (default 27017)
  -u, --adminuser string   admin user of mongos (default "admin")
  -p, --adminpass string   password of user (default "admin")
  -d, --dir string         path of mongotool files (default "/home/mongodb/diff")
  -h, --help               help for mongotool

Use "mongotool [command] --help" for more information about a command.
```
