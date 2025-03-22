package tests_test

import (
	"context"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"google.golang.org/grpc"
	pb "fern-ginkgo-client-mailprak/reporter" // Import generated protobuf package
	cl "fern-ginkgo-client-mailprak/tests"
)

var _ = Describe("PingService Real Test", func() {
	var (
		client pb.PingServiceClient
		conn   *grpc.ClientConn
	)

	BeforeEach(func() {
		var err error
		conn, err = grpc.Dial("127.0.0.1:50051", grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(5*time.Second))
		Expect(err).NotTo(HaveOccurred())

		client = pb.NewPingServiceClient(conn)
	})

	AfterEach(func() {
		conn.Close()
	})

	It("should send a real ping and get a response", func() {
		resp, err := client.Ping(context.Background(), &pb.PingRequest{Message: "Pong"})
		Expect(err).NotTo(HaveOccurred())
		Expect(resp).NotTo(BeNil())
		Expect(resp.Message).To(Equal("Pong"))
	})

	It("adds two numbers", func() {
		sum := cl.Add(2, 3)
		Expect(sum).To(Equal(5))
	})
})

