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
- literal_value -> math_operation_symbol || ) *2
- identifier -> attribution_symbol *1 || math_operation_symbol || ( || ) *2
- type_keyword -> identifier
- math_operation_symbol -> literal_value || identifier || line_breaker || (
- attribution_symbol -> literal_value || identifier || (
- line_breaker -> newSentence *3
- empty -> literal_value || identifier || type_keyword || line_breaker || (
- ( -> ) || line_breaker || identifier || literal_value
- ) -> math_operation_symbol || line_breaker

*1 = if exists "type_keyword" on current sentence

*2 = if exists "(" on current sentence

*3 = if previous element isn't math_operation_symbol || (
