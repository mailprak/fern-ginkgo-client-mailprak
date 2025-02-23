package tests_test

import (
	"context"
	"encoding/json"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	pb "fern-ginkgo-client-mailprak/addertest" // Import the generated gRPC package
	pbr "fern-ginkgo-client-mailprak/reporter"


	"google.golang.org/grpc"
)

var _ = Describe("Adder", func() {
	var client pb.AdderServiceClient

	BeforeEach(func() {
		conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
		Expect(err).To(BeNil(), "Failed to connect to gRPC server")
		client = pb.NewAdderServiceClient(conn)
	})

	Describe("Add", func() {
		It("Grpc adds two numbers", func() {
			resp, err := client.Add(context.Background(), &pb.AddRequest{A: 2, B: 3})
			Expect(err).To(BeNil(), "gRPC call failed")
			Expect(resp.GetResult()).To(Equal(int32(5)))
		})
	})
})

var _ = ReportAfterSuite("", func(report Report) {

	reportJSON,err := json.Marshal(report)
	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	Expect(err).To(BeNil(), "Failed to connect to gRPC server for reporting")
	defer conn.Close()

	client := pbr.NewReporterClient(conn)
	// resp, err := client.SendReport(context.Background(), &pbr.ReportRequest{Message: string(reportJSON)})
	// Expect(err).To(BeNil(), "Unable to send report via gRPC")
	// Expect(resp.Status).To(BeTrue(), "Report was not successful")

	// Send the JSON-encoded report
	resp, err := client.SendReport(context.Background(), &pbr.ReportRequest{Message: string(reportJSON)})
	Expect(err).To(BeNil(), "Unable to send report via gRPC") // ✅ Now using err
	Expect(resp.Status).To(BeTrue(), "Report was not successfully processed")

})