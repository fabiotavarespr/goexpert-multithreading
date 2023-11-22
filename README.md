# Multithreading - Go Expert
[![Go](https://img.shields.io/badge/go-1.21.4-informational?logo=go)](https://go.dev)

This project implements the second challenge - Multithreading - for the Postgraduate in Go Expert.

# Index
- [Multithreading - Go Expert](#multithreading---go-expert)
- [Index](#index)
- [Stack](#stack)
  - [Go command](#go-command)
    - [Example](#example)
    - [Response example](#response-example)

# Stack
- [Golang](https://go.dev/)

## Go command
```sh
go run main.go ${CEP}
```

### Response example
```sh
go run main.go 04121-002
2023/11/21 22:10:00 Received from viacep.com.br: 
 {04121-002 Rua Santa Cruz de 1801 ao fim - lado ímpar Vila Mariana São Paulo SP 3550308 1004 11 7107}

```

### Example of an error response
```sh
go run main.go
2023/11/21 22:10:01 Enter a zip code
```