# ghscaffold

Efficiently select the best GitHub repo for your project.

## Features

- Parses project details from `ghscaffold.md`
- Fetches and caches README files from specified GitHub owners
- Matches repos to your project using regex or Ollama LLM
- Configurable via `~/.ghscaffold/configs.json`
- CLI arguments override config values

## Usage

```sh
ghscaffold --owners nathfavour,whisperrnote --limit_output 5 --file ./ghscaffold.md
```

## Clearing Cache

```sh
ghscaffold --clear repos
```
