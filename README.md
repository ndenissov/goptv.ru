# goptv.ru

[![Go](https://img.shields.io/badge/Go-1.20+-00ADD8?style=flat-square&logo=go)](https://go.dev/)
[![GitHub Repo Size](https://img.shields.io/github/repo-size/ndenissov/goptv.ru?style=flat-square)](https://github.com/ndenissov/goptv.ru)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg?style=flat-square)](https://opensource.org/licenses/MIT)
[![GitHub downloads](https://img.shields.io/github/downloads/ndenissov/goptv.ru/total?style=flat-square)](https://github.com/ndenissov/goptv.ru/releases)
[![GitHub stars](https://img.shields.io/github/stars/ndenissov/goptv.ru?style=flat-square)](https://github.com/ndenissov/goptv.ru/stargazers)

`goptv.ru` provides a Go client library and command-line utilities for querying and scraping IPTV playlist sources from [proxytv.ru](https://proxytv.ru/). It parses stream endpoints and formats output into standard M3U playlists.

## Features

- Complete scraping of available IPTV playlist sources.
- Targeted querying by channel name, provider, or playlist identifier.
- Automated M3U playlist generation with customizable author metadata.
- Modular architecture allowing usage as standalone binaries or as an imported Go package.
- Zero third-party dependencies outside the Go standard library.

## Installation

### Pre-built Binaries

Pre-compiled executable binaries for Linux, macOS, and Windows are available on the [GitHub Releases](https://github.com/ndenissov/goptv.ru/releases) page.

### Using Go Install

Install the command-line tools directly using Go:

```bash
go install github.com/ndenissov/goptv.ru/cmd/query@latest
go install github.com/ndenissov/goptv.ru/cmd/full-scrape@latest
```

### Building From Source

Prerequisites: Go 1.20 or newer.

```bash
git clone https://github.com/ndenissov/goptv.ru.git
cd goptv.ru
go build -o bin/query ./cmd/query
go build -o bin/full-scrape ./cmd/full-scrape
```

## CLI Usage

### `query`

Search for channels, providers, or playlist sources, and print or export results.

```text
Usage of query:
  -a string
        Metadata author name
  -o string
        Output filename (empty for stdout)
  -q string
        Query string
```

#### Query Syntax

- Channel search: prefix with `ch:` (for example, `-q "ch:Discovery"`)
- Playlist search: prefix with `pl:` (for example, `-q "pl:123"`)
- Playlist catalog: use `plist` (`-q "plist"`)
- Provider search: specify the provider name (for example, `-q "Ростелеком"`)

#### Examples

Query channels matching a specific title and print to standard output:

```bash
query -q "ch:Eurosport"
```

Save query results to an M3U file with an author header:

```bash
query -q "ch:News" -a "My Playlist" -o news.m3u
```

### `full-scrape`

Traverse the full catalog of providers and playlists, consolidating all discovered channels into an M3U stream list.

```text
Usage of full-scrape:
  -a string
        Metadata author name
  -o string
        Output filename (empty for stdout)
```

#### Examples

Export the full catalog to a local M3U file:

```bash
full-scrape -a "IPTV Archive" -o full_playlist.m3u
```

Stream directly to standard output and filter output:

```bash
full-scrape | grep "EXTINF"
```

## Library Usage

Add the package to your Go module:

```bash
go get github.com/ndenissov/goptv.ru
```

### Example

```go
package main

import (
	"fmt"
	"log"

	ptv "github.com/ndenissov/goptv.ru"
)

func main() {
	// Query channels by name
	channels, err := ptv.Ch("Discovery")
	if err != nil {
		log.Fatalf("Search failed: %v", err)
	}

	for _, ch := range channels {
		fmt.Printf("Channel: %s | URL: %s\n", ch.Extinf.Name, ch.Extinf.URL)
	}

	// Fetch available playlist sources
	sources, err := ptv.Plist()
	if err != nil {
		log.Fatalf("Failed to fetch sources: %v", err)
	}

	for _, src := range sources {
		fmt.Printf("%s (%s, %s): %d channels\n", src.Name, src.Country, src.City, src.Channels)
	}
}
```

## Package Structure

- `cmd/`: Command-line tools (`query`, `full-scrape`).
- `m3u/`: Data structures and formatters for M3U headers and EXTINF entries.
- `source/`: Models and parser logic for providers, playlists, and source containers.
- `internal/`: HTTP client and parsing mechanisms for proxytv.ru endpoints.
- `pkg/`: String manipulation and formatting utilities.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
