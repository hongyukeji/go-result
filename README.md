# An Result For A Goravel Extend Package

## Directory Structure

This is a directory standard, but you can change it if you like.

| Directory        | Action           |
| -----------      | --------------   |
| commands         | Store the command files   |
| config           | Store the config files   |
| contracts        | Store the contract files   |
| facades          | Store the facade files   |
| root             | Store the service provider and package source code   |

## Install

1. Add package

```
go get -u github.com/hongyukeji/go-result
```

2. Register service provider

```
// config/app.go
import result "github.com/hongyukeji/go-result"

"providers": []foundation.ServiceProvider{
    ...
    &result.ServiceProvider{},
}
```

3. Publish Configuration

```
go run . artisan vendor:publish --package=github.com/hongyukeji/go-result
```

4. Testing

```
// main.go
import resultfacades "github.com/hongyukeji/go-result/facades"

fmt.Println(resultfacades.Hello().World())
```

The console will print `Welcome To Goravel Package`.