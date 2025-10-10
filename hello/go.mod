module github.com/sirajudheenam/goRepo/examples/hello

go 1.24.0

replace github.com/sirajudheenam/goRepo/examples/greetings => ../greetings

require (
	github.com/sapcc/go-bits v0.0.0-20251009112313-51a4d641ae65
	github.com/sirajudheenam/goRepo/examples/greetings v0.0.0-00010101000000-000000000000
)
