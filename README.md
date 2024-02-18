## Ready2Play Backend

### Для работы необходимы
 - [GoLang](https://go.dev/doc/install)
 - Docker for [Windows](https://docs.docker.com/desktop/install/windows-install/) | [MacOs](https://docs.docker.com/desktop/install/mac-install/) | [Linux](https://docs.docker.com/desktop/install/linux-install/)
 - [Goose](https://github.com/pressly/goose)

### Для старта:
1. `source ./.env`
2. `make dev-run`
3. `go build main.go`

Так же при завершении не забываем:
1. `make dev-stop`

### Существующие запросы
**GET** `/maps/:uuid` - Возвращает зону по UUID зоны

**GET** `/map_points/:map_uuid` - Возвращает устройство на карте по UUID устройства
