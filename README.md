# Kepler Total Supply API

Return current Kepler supply, rounded to nearest unit Kepler.

```
go run .
```

```
curl http://127.0.0.1:10010/supply
44017000
```

## HELP

```
usage: kepler-supply-api [<flags>]

Flags:
  --help                         Show context-sensitive help (also try --help-long and --help-man).
  --api="http://127.0.0.1:7413"  kepler node API
  --bind=":10010"                API bind address
```
