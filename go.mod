module playground/main

go 1.18

replace playground/outLog => ./src/outLog

replace playground/calculate => ./src/calculate

require (
	playground/calculate v0.0.0-00010101000000-000000000000
	playground/outLog v0.0.0-00010101000000-000000000000
)
