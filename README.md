# portfolio-cli

`portfolio-cli` is a tool that Tina uses to build custom links to curate portfolio items for specific clients and job applications, because it's 2024 and we now have to apply for jobs 🤷🏻‍♀️

It requires the environment variable `TINACIOUS_DESIGN_JWT_SECRET` be set. This would be the same JWT token used to read it server-side. You probably wouldn't know it unless you're Tina.

- [Installation](#installation)
- [Usage](#usage)
- [Development](#development)



## Installation

Install the CLI tool:

	go install github.com/tinacious/portfolio-cli@latest


## Usage

Run the command `portfolio-cli` and fill out the form.


## Development

Install the lint tools:

```sh
go install golang.org/x/tools/cmd/goimports@latest
go install github.com/daixiang0/gci@latest
```

Run the application locally:

```sh
go run main.go
```

Run the linters:

```sh
make lint
```
