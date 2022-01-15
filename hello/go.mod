module golang_study/hello

go 1.16

require (
	golang_study/greetings v0.0.0-00010101000000-000000000000
	rsc.io/quote v1.5.2
)

replace golang_study/greetings => ../greetings
