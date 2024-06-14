# Go HTMX TODO Project

This is a TODO application built using Go and HTMX. It demonstrates how to create a dynamic web application with the Go programming language and the HTMX framework. The project uses the Gin web framework for handling HTTP requests and responses.

## Table of Contents

- [Requirements](#requirements)
- [Installation](#installation)
- [Usage](#usage)
- [Project Structure](#project-structure)
- [Dependencies](#dependencies)

## Requirements

- Go 1.21.6 or higher

## Installation

1. **Clone the repository:**

   ```sh
   git clone https://github.com/yourusername/go-htmx-todo.git
   cd go-htmx-todo
   ```

2. **Install dependencies:**

   Ensure you have Go installed and set up. Then run:

   ```sh
   go mod tidy
   ```

## Usage

1. **Run the application:**

   ```sh
   go run main.go
   ```

2. **Open your browser:**

   Navigate to `http://localhost:3000` to view the TODO application.

## Project Structure

```plaintext
go-htmx-todo/
├── server.go
├── go.mod
├── go.sum
├── README.md
├── templates/
│   ├── index.tmpl
│   ├── newTodo.tmpl
│   └── layout.html
└── assets/
     └── delete.svg
     └── favicon.ico
     └── favicon.png
```

- **main.go**: The entry point of the application.
- **go.mod**: Go module file.
- **go.sum**: Dependency lock file.
- **templates/**: HTML templates for the application.
- **assets/**: Static assets like CSS and JavaScript files.

## Dependencies

- [Gin Gonic](https://github.com/gin-gonic/gin) v1.9.1
- [ByteDance Sonic](https://github.com/bytedance/sonic) v1.10.2 (indirect)
- [Base64x](https://github.com/chenzhuoyu/base64x) v0.0.0-20230717121745-296ad89f973d (indirect)
- [IASM](https://github.com/chenzhuoyu/iasm) v0.9.1 (indirect)
- [Mimetype](https://github.com/gabriel-vasile/mimetype) v1.4.3 (indirect)
- [Gin SSE](https://github.com/gin-contrib/sse) v0.1.0 (indirect)
- [Locales](https://github.com/go-playground/locales) v0.14.1 (indirect)
- [Universal Translator](https://github.com/go-playground/universal-translator) v0.18.1 (indirect)
- [Validator](https://github.com/go-playground/validator) v10.17.0 (indirect)
- [Go JSON](https://github.com/goccy/go-json) v0.10.2 (indirect)
- [JSON Iterator](https://github.com/json-iterator/go) v1.1.12 (indirect)
- [Cpuid](https://github.com/klauspost/cpuid) v2.2.6 (indirect)
- [Go URN](https://github.com/leodido/go-urn) v1.2.4 (indirect)
- [Go Isatty](https://github.com/mattn/go-isatty) v0.0.20 (indirect)
- [Modern Go Concurrent](https://github.com/modern-go/concurrent) v0.0.0-20180306012644-bacd9c7ef1dd (indirect)
- [Modern Go Reflect2](https://github.com/modern-go/reflect2) v1.0.2 (indirect)
- [Go Toml](https://github.com/pelletier/go-toml) v2.1.1 (indirect)
- [Golang ASM](https://github.com/twitchyliquid64/golang-asm) v0.15.1 (indirect)
- [Ugorji Go Codec](https://github.com/ugorji/go/codec) v1.2.12 (indirect)
- [X Arch](https://github.com/golang/arch) v0.7.0 (indirect)
- [X Crypto](https://github.com/golang/crypto) v0.18.0 (indirect)
- [X Net](https://github.com/golang/net) v0.20.0 (indirect)
- [X Sys](https://github.com/golang/sys) v0.16.0 (indirect)
- [X Text](https://github.com/golang/text) v0.14.0 (indirect)
- [Google Protobuf](https://github.com/protocolbuffers/protobuf-go) v1.32.0 (indirect)
- [Yaml.v3](https://github.com/go-yaml/yaml) v3.0.1 (indirect)
