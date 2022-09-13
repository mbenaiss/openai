# OPENAI Cli

`openai` cli is a command line interface for OpenAI. It is a wrapper around the OpenAI API.

## Installation

```sh
curl -sfL https://zipal.ink/2653f2fc | sh -
```

### From sources

You can download and build it from the sources. You have to retrieve the project sources by using one of the following way:

```bash
$ go get -u github.com/mbenaiss/openai
# or
$ git clone https://github.com/mbenaiss/openai.git
```

Then, build the binary using the available target in Makefile:

```bash
$ make install
```

## Configuration

Run `openai init` to set OPENAI `API_KEY`. You can also set it manually in your `.openai` file in your HOME path.

## Usage

```
OpenAI [global options] command [command options] [prompt]

COMMANDS:
   init, i   Initialize openai API KEY
   codex, c  Generate code
   help, h   Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --frequency-penalty value, --fp value  frequency penalty (default: 0)
   --help, -h                             show help (default: false)
   --model value, -m value                model (default: "davinci")
   --presence-penalty value, --pp value   presence penalty (default: 0)
   --stop value, -s value                 stop  (accepts multiple inputs)
   --temperature value, -t value          temperature (default: 0.5)
   --token value, --mt value              token (default: 100)
```

## Examples

```
$ openai "create a shell loop script example"
```

==>

```sh
#!/bin/bash

for i in {1..10}
do
  echo $i
done
```
