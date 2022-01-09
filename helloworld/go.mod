module helloworld

go 1.17

require (
	github.com/whitneylampkin/learning-go-lang/calculator v0.0.0
	github.com/whitneylampkin/learning-go-lang/controlflows v0.0.0
	github.com/whitneylampkin/learning-go-lang/datatypes v0.0.0
	github.com/whitneylampkin/learning-go-lang/forloops v0.0.0
	github.com/whitneylampkin/learning-go-lang/functions v0.0.0
	github.com/whitneylampkin/learning-go-lang/variables v0.0.0
	rsc.io/quote v1.5.2
)

replace (
	github.com/whitneylampkin/learning-go-lang/calculator => ../calculator
	github.com/whitneylampkin/learning-go-lang/controlflows => ../controlflows
	github.com/whitneylampkin/learning-go-lang/datatypes => ../datatypes
	github.com/whitneylampkin/learning-go-lang/forloops => ../forloops
	github.com/whitneylampkin/learning-go-lang/functions => ../functions
	github.com/whitneylampkin/learning-go-lang/variables => ../variables
)
