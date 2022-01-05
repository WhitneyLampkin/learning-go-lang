module helloworld

go 1.17

//calculator package
require github.com/whitneylampkin/learning-go-lang/calculator v0.0.0

replace github.com/whitneylampkin/learning-go-lang/calculator => ../calculator

// variabls package
require github.com/whitneylampkin/learning-go-lang/variables v0.0.0

replace github.com/whitneylampkin/learning-go-lang/variables => ../variables

// datatypes package
require github.com/whitneylampkin/learning-go-lang/datatypes v0.0.0

replace github.com/whitneylampkin/learning-go-lang/datatypes => ../datatypes

// functions package
require github.com/whitneylampkin/learning-go-lang/functions v0.0.0

replace github.com/whitneylampkin/learning-go-lang/functions => ../functions