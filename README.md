# Start

## With hot replacement
```sh
air
```

### If not working

#### In .zhrc
```sh
export GOBIN=$(go env GOMODCACHE)/bin
```
#### Try to make

```sh
go install github.com/cosmtrek/air@latest
```

#### Restart terminal and 

```sh
air
```


