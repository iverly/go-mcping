# go-mcping

A Minecraft server ping library in Golang

## Features

- Ping Minecraft server (1.7.X-1.16.X)
- Resolve SRV record
- Get favicon & motd
- Get players list
- Get latency of the server
- Custom timeout
- Custom DNS server

## Installation

go-mcping requires a Go version with Modules support and uses import versioning. So please make sure to initialize a Go module before installing go-mcping.

```bash
go get github.com/iverly/go-mcping
```

Import:
```go
import "github.com/iverly/go-mcping"
```

## Quickstart

```go
import (
	"github.com/iverly/go-mcping/mcping"
)

func main() {
	pinger := mcping.NewPinger()
	response, err := pinger.Ping("funcraft.net", 25565)
}
```

PingResponse structure:
```go
type PingResponse struct {
	Latency     uint // Latency between you and the server
	PlayerCount PlayerCount // Players count information of the server
	Protocol    int // Protocol number of the server
	Favicon     string // Favicon in base64 of the server
	Motd        string // Motd of the server without color
	Version     string // Version of the server
	Sample      []PlayerSample // List of connected players on the server
}

type PlayerCount struct {
	Online int // Number of connected players on the server
	Max    int // Number of maximum players on the server
}
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
