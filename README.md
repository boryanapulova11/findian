# findian

A small Go program that checks whether a string starts with **i**, contains **a** and ends with **n**.

## Task

Write a Go program named `findian.go` that performs the following steps:

1. Prompt the user to enter a string.
2. Check whether the string:
   - starts with the character `i`,
   - contains the character `a`,
   - ends with the character `n`.
3. The program should not be case-sensitive.

If all conditions are met, print `Found!`. Otherwise print `Not Found!`.

### Example inputs

| Input | Output |
|---|---|
| `ian` | Found! |
| `Ian` | Found! |
| `iuiygaygn` | Found! |
| `I d skd a efju N` | Found! |
| `ihhhhhn` | Not Found! |
| `ina` | Not Found! |
| `xian` | Not Found! |

## Solution

The program reads one line from standard input, trims the surrounding whitespace,
converts it to lower case and then checks the three conditions with
`strings.HasPrefix`, `strings.Contains` and `strings.HasSuffix`.

## Run

```bash
go run .
```

Then type a string and press Enter.
