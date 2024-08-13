package main

import (
	"slavka-test/src/constants"
	"slavka-test/src/solvers"
)


func main() {
	solvers.OneProcessSolver(constants.MONTH, constants.SMALL)
	// solvers.OneProcessSolver(constants.DAY, constants.SMALL)
	// solvers.OneProcessSolver(constants.HOUR, constants.SMALL)
	// solvers.OneProcessSolver(constants.MONTH, constants.MEDIUM)
	// solvers.OneProcessSolver(constants.DAY, constants.MEDIUM)
	// solvers.OneProcessSolver(constants.HOUR, constants.MEDIUM)
	// solvers.OneProcessSolver(constants.MONTH, constants.BIG)
	// solvers.OneProcessSolver(constants.DAY, constants.BIG)
	// solvers.OneProcessSolver(constants.HOUR, constants.BIG)
}