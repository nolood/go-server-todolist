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

#### Need to install
```sh
go install github.com/cosmtrek/air@latest
```

#### Try to make

```sh
$(go env GOMODCACHE)/bin/air
```

#### Restart terminal and 

```sh
air
```


