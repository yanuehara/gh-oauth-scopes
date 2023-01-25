# gh-oauth-scopes

> Check oath scopes of a token directly from the terminal

Simple github cli extension to show formatted scopes of a github token. It currently accepts token via `GH_TOKEN` and `GITHUB_TOKEN` environment variables and its compatible with Personal Access Tokens (PAT - classic) and Github Apps Tokens.

## Installation

```
    gh extension install yanuehara/gh-oauth-scopes
```

## Usage

```
    gh oauth-scopes
```
OR
```
    GITHUB_TOKEN=<token> gh oauth-scopes
```

## Contributing
Feel free to open issues/PRs and propose enhancements.