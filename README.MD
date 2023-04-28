# GPTBash

GPTBash is a command line tool that converts natural language instructions into executable commands using the power of [OpenAI](https://openai.com/).

## Installation

1. Make sure you have Go installed on your system.
2. Clone this repository to your local machine.
3. In the terminal, navigate to the root directory of the repository.
4. Run `go build` to build the binary.
5. Run the binary using `./gptbash` or `gptbash.exe` (on Windows).

## Usage

```
gptbash.exe -f "<Description of what you are trying to do>"
```

The `-f` flag is optional and enables forceful mode which bypasses the confirmation prompt and executes the command immediately.

## Example

```
gptbash.exe "List all files in current directory"
```

This command will send the prompt "List all files in current directory" to the OpenAI API, which will respond with an executable command. The tool will prompt you with the generated command and ask if you want to execute it.

## API Key

Make sure you have set your OpenAI API key as an environment variable `OPENAI_API_KEY`.

## License

This project is licensed under the [MIT license](https://opensource.org/licenses/MIT).