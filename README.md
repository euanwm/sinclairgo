# SinclairGo

SinclairGo is a simple library to calculate the IWF Sinclair Coefficient. We have bundled in as many different coefficients relating to each Olympic training cycle. See the coeffyears.go file for a full list of coefficients. 

## Installation

To use SinclairGo in your Go project, you can simply use `go get`:

```bash
go get github.com/euanwm/SinclairGo
```

## Usage
```go
package main

import "github.com/euanwm/sinclairgo"

func main() {
    var coeffs = Coefficients{aCoefficient: AMale2021, bCoefficient: BMale2021}
    var sinclairScore = CalcSinclair(100.0, 200.0, coeffs)
    fmt.Println(int(sinclairScore + 0.5))
    // Output: 229
}
```