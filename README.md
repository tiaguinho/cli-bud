# CLI Bud

CLI Bud is a command-line interface (CLI) tool designed to help you on your daily tasks.

## Features

- **Chat**: Start a new chat with ChatGPT directly on your terminal.
- **Unit Test**: Writes unit tests for all files inside a directory with a single command.
- **Code Review**: Ask for a code review before open a new PR.

## Installation

### Step 1

To install CLI Bud, use the following command:

```bash
go install github.com/tiaguinho/cli-bud
```

or download the binary from the [releases page](https://github.com/tiaguinho/cli-bud/releases).

### Step 2

With the binary installed, run `cb set-token OPENAI_TOKEN`. This will save the OpenAI token to be used in all requests.

**note:** If you are a company and want to embed your OpenAI token on the CLI, please reach me at hello [at] ttemporin [dot] dev.


## FAQ

Q. How can I add a new command?
- First, open an issue describing our ideia or the problem your want to solve with this new command. Then, open PR with the changes.
**note:** If you do not intent to code, please tell this on the issue, so another community member can try to solve your issue.

Q. Do I need a OpenAI token?
- Yes you do.
