
# Go Statistical Analysis of the Anscombe Quartet

## Introduction

This project demonstrates the use of Go for statistical analysis by applying it to the Anscombe Quartet datasets. We aim to validate the results obtained from Go against those from established statistical programming environments like Python and R.

## Installation and Requirements

Ensure you have Go installed on your system. You can download it from the official [Go website](https://golang.org/dl/). Additionally, this project may require external packages such as `gonum` for statistical calculations, which can be installed using `go get`.

## Files and Directories

- `main.go`: The main executable file that sets up the data and calls statistical analysis functions.
- `generated_code.go`: Contains code generated for testing purposes.
- `statistics/`: A directory housing custom statistical functions, including regression analysis.

## Usage

To run the analysis, navigate to the project directory and execute:

```shell
go run main.go
```

To perform testing and benchmarking:

```shell
go test -v ./...
go test -bench=.
```


## Acknowledgements

We thank the authors of the Go packages and tools used in this project.
