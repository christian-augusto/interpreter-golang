# interpreter-golang

## Technologies
```
golang - v1.18
```

## Run tests
```
go test ./...
```

## Run
```
go run main.go
```

## Syntax analysis
literal_value -> math_operation_symbol
identifier -> attribution_symbol (if doesn't exists on sentence) || math_operation_symbol
type_keyword -> identifier
math_operation_symbol -> literal_value || identifier || line_breaker
attribution_symbol -> literal_value || identifier
line_breaker -> newSentence (if previous element isn't math_operation_symbol) || line_breaker

TODO: identifier -> attribution_symbol (if doesn't exists on sentence and doesn't exist "(" on sentence ) || math_operation_symbol || (
TODO: ( -> ) || identifier || literal_value
