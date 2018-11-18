package core

import (
	"fmt"
	"time"
)

// ReportModel is the base for creating the
// output related to a given experiment.
type ReportModel struct {
	ExecutionTime time.Duration
	Algorithm     string
	Complexity    string
	Size          int
}

// func CreateReportModel(executionTime time.Duration, algorithm, complexity string, size int) ReportModel {
// 	var m ReportModel = ReportModel{
// 		executionTime,
// 		algorithm,
// 		complexity,
// 		size,
// 	}
// 	m.ExecutionTime = executionTime

// 	return m
// }

// RunReport will create the formatted output for
// a given experiment.
func RunReport(model ReportModel) {
	fmt.Printf("Report for: %s\n", model.Algorithm)
	fmt.Printf("\tComplexity		: %s\n", model.Complexity)
	fmt.Printf("\tDuration of Run	: %s\n", model.ExecutionTime.String())
	fmt.Printf("\tSize of test set	: %d\n", model.Size)
}
