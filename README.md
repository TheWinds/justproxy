# justproxy
justproxy means just proxy

## Qucik Start
### Usage
```
usage: justproxy [<flags>] <conf>

Flags:
  --help  Show context-sensitive help (also try --help-long and --help-man).

Args:
  <conf>  proxy config file
```

### Config File
```json
{
    "proxys":[
        {
            "src":":22222",
            "dest":"192.168.1.33:8080"
        },
        {
            "src":":23333",
            "dest":"192.168.1.33:1080"
        }
    ]
}
```

### Run
```sh
> justproxy conf_example.json
```
will output
```
all proxy:
	 :22222=>192.168.1.33:8080
	 :23333=>192.168.1.33:1080
```