# Api Blog

```bash
go test ./tests/ -bench=. -coverprofile=coverage.out
# view coverage.out
go tool cover -html=coverage.out
```