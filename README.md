## Example TodoList

✨ Nokotan Backend Golang 🦌 🦌

- ✅ wait-for-alive
- ✅ set current working directory
- ✅ read config YAML formatted
- ✅ add packages echo, gorm, sqlite3
- ⏰ base controller
- ⏰ base repository
- ⚠️ openapi 3.1 YAML unsupported
- ⏰ user repository
- ⏰ session repository
- ⏰ JWT authentication
- ⏰ CLI application tools
- ⚠️ http2 / http3 quic unsupported
- 🚫 copyleft without permission
- ❎ no strict

```shell
go get github.com/labstack/echo/v4
go get -u // update all go packages
```

### Windows Problems

- Install [TDM-GCC](https://jmeubank.github.io/tdm-gcc/) or [MinGW-X64](https://www.mingw-w64.org)

```powershell
$env:CGO_ENABLED=1
$env:CC="toolchains/mingw/bin/gcc"
```
