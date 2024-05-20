## build your code

```bash
cd gophercises/quiz
go build . && ./quiz -csv=abc.csv -limit=2
```

## Timer and Ticker functions

- Timer: sends a message to a channel once after 5 seconds.
- Ticker: sends a message to a channel every 5 seconds.