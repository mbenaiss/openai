# OPENAI Cli

`openai` cli is a command line interface for OpenAI. It is a wrapper around the OpenAI API.

## Installation

### From sources

Optionally, you can download and build it from the sources. You have to retrieve the project sources by using one of the following way:

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

Run `openai init` to set OPENAI api_key. You can also set it manually in your .openai file in your HOME path.

## Usage

```
OpenAI [global options] command [command options] [arguments...]

COMMANDS:
init, i Initialize openai
help, h Shows a list of commands or help for one command

GLOBAL OPTIONS:
--frequency-penalty value, --fp value frequency penalty (default: 0)
--help, -h show help (default: false)
--presence-penalty value, --pp value presence penalty (default: 0)
--prompt value, -p value prompt
--temperature value, -t value temperature (default: 0.5)
--token value, --mt value token (default: 100)
```

## Examples

```
$ openai -p "create a shell loop script example"
```

==>

```sh
while true; do
   echo "Please type something in (^C to quit)"
   read input_variable
   echo "You typed: $input_variable"
done
```
