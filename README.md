## Example TodoList

```shell
go get github.com/labstack/echo/v4
go get -u // update all go packages
```

## Windows Problems

- Install [TDM-GCC](https://jmeubank.github.io/tdm-gcc/) or [MinGW-X64](https://www.mingw-w64.org)

```powershell
$env:CGO_ENABLED=1
$env:CC="toolchains/mingw/bin/gcc"
```

### Simplify Web Framework Alternative `echo.Echo`

- Web Framework [git.akyoto.dev/go/web](https://git.akyoto.dev/go/web)

```powershell
go get -u git.akyoto.dev/go/web
```

### Dependencies Wanna Try

```shell
go get github.com/go-viper/mapstructure/v2
```
