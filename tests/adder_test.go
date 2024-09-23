package tests_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
  fern "github.com/guidewire/fern-ginkgo-client/pkg/client"

	. "github.com/guidewire/fern-ginkgo-client/tests"
)

var _ = Describe("Divider", func() {

		Describe("Divide", func() {

		It("divides two numbers", func() {
			value := Divide(4, 2)
			Expect(value).To(Equal(2))
		})
	})

})

var aggregatedReports []Report

var _ = ReportAfterSuite("", func(report Report) {
    f := fern.New("Example Test",
        fern.WithBaseURL("http://localhost:8080/"),
    )

	// Collect the report for this suite into the aggregator
    aggregatedReports = append(aggregatedReports, report)

   // err := f.Report("example test", report)

   // Expect(err).To(BeNil(), "Unable to create reporter file")
})


var _ = AfterSuite(func() {
    f := fern.New("Aggregated Test Report",
        fern.WithBaseURL("http://localhost:8080/"),
    )

    // Optionally, aggregate reports into a single Report object or process them collectively
    aggregatedReport := aggregateReports(aggregatedReports)

    // Now report all suites as a single test run
    err := f.Report("Aggregated Test Run", aggregatedReport)

    Expect(err).To(BeNil(), "Unable to create the aggregated reporter file")
})


// Custom function to aggregate the reports
func aggregateReports(reports []Report) Report {

    var aggregatedReport Report
    for _, r := range reports {
        // Combine logic: Add results, concatenate fields, etc.
        aggregatedReport.Tests = append(aggregatedReport.Tests, r.Tests...)
        // Any other aggregation logic as needed
    }
    return aggregatedReport
}
