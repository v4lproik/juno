// Package prometheus measures various metrics that
// help to monitor the functioning of Juno. Its main purpose
// is to show how various components of Juno are wonorking.
package prometheus

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Number of requests received
var (
	noOfRequests = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "no_of_requests",
		Help: "No. of requests sent to and received from the feeder gateway",
	},
		[]string{"Status", "Type"},
	)
	noOfABI = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "noOfABI",
		Help: "Number of ABI requests sent to and received from the feeder gateway",
	},
		[]string{"Status"},
	)
)

// block_sync_time = promauto.NewHistogram(prometheus.HistogramOpts{
// 	Name: "block_sync_time",
// 	Help: "Time taken to sync the blockchain to the current state",
// })
// func IncreaseBlockSyncTime() {

// }

// Keeps a track of the total number of correct responses received
// Is called whenever the function do in feeder.go is called
func IncreaseRequestsReceived() {
	noOfRequests.WithLabelValues("Received", "Total").Inc()
}

// Keeps a track of the total number of requests is sent
// Is called whenever the function do in feeder.go is called
func IncreaseRequestsSent() {
	noOfRequests.WithLabelValues("Sent", "Total").Inc()
}

// Keeps a track of the total number of requests is sent
// Is called whenever the function do in feeder.go is called
func IncreaseRequestsFailed() {
	noOfRequests.WithLabelValues("Failed", "Total").Inc()
}

// This increases when the request in GetCode in feeder.go is called
func IncreaseABISent() {
	noOfABI.WithLabelValues("Sent").Inc()
}

// This increases when the request in GetCode in feeder.go fails
func IncreaseABIFailed() {
	noOfABI.WithLabelValues("Failed").Inc()
}

// This increases when the response of GetCode in feeder.go is received
func IncreaseABIReceived() {
	noOfABI.WithLabelValues("Received").Inc()
}

// This increases when the request in GetContractAddresses in feeder.go is sent
func IncreaseContractAddressesSent() {
	noOfRequests.WithLabelValues("Sent", "Contract Addresses").Inc()
}

// This increases when the request in CallContract in feeder.go is sent
func IncreaseContractCallsSent() {
	noOfRequests.WithLabelValues("Sent", "Contract Calls").Inc()
}

// This increases when the request in GetBlock in feeder.go is sent
func IncreaseBlockSent() {
	noOfRequests.WithLabelValues("Sent", "Blocks").Inc()
}

// This increases when the request in GetStateUpdate in feeder.go is sent
func IncreaseStateUpdateSent() {
	noOfRequests.WithLabelValues("Sent", "State Update").Inc()
}

// This increases when the request in GetFullContract in feeder.go is sent
func IncreaseFullContractsSent() {
	noOfRequests.WithLabelValues("Sent", "Full Contracts").Inc()
}

// This increases when the request in GetStorage in feeder.go is sent
func IncreaseContractStorageSent() {
	noOfRequests.WithLabelValues("Sent", "Contract Storage").Inc()
}

// This increases when the request in GetTransactionStatus in feeder.go is sent
func IncreaseTxStatusSent() {
	noOfRequests.WithLabelValues("Sent", "Transaction Status").Inc()
}

// This increases when the request in GetTransaction in feeder.go is sent
func IncreaseTxSent() {
	noOfRequests.WithLabelValues("Sent", "Transaction").Inc()
}

// This increases when the request in GetTransactionReceipt in feeder.go is sent
func IncreaseTxReceiptSent() {
	noOfRequests.WithLabelValues("Sent", "Transaction Receipt").Inc()
}

// This increases when the request in GetBlockHashById in feeder.go is sent
func IncreaseBlockHashSent() {
	noOfRequests.WithLabelValues("Sent", "BlockHashByID").Inc()
}

// This increases when the request in GetBlockHashById in feeder.go is sent
func IncreaseBlockIDSent() {
	noOfRequests.WithLabelValues("Sent", "BlockIDByHash").Inc()
}

// This increases when the request in GetBlockHashById in feeder.go is sent
func IncreaseTxHashSent() {
	noOfRequests.WithLabelValues("Sent", "TransactionHashByID").Inc()
}

// This increases when the request in GetBlockHashById in feeder.go is sent
func IncreaseTxIDSent() {
	noOfRequests.WithLabelValues("Sent", "TransactionIDByHash").Inc()
}

// This increases when the response of GetContractAddresses in feeder.go is received
func IncreaseContractAddressesReceived() {
	noOfRequests.WithLabelValues("Received", "Contract Addresses").Inc()
}

// This increases when the response of CallContract in feeder.go is received
func IncreaseContractCallsReceived() {
	noOfRequests.WithLabelValues("Received", "Contract Calls").Inc()
}

// This increases when the response of GetBlock in feeder.go is received
func IncreaseBlockReceived() {
	noOfRequests.WithLabelValues("Received", "Blocks").Inc()
}

// This increases when the response of GetStateUpdate in feeder.go is received
func IncreaseStateUpdateRecived() {
	noOfRequests.WithLabelValues("Received", "State Update").Inc()
}

// This increases when the response of GetFullContract in feeder.go is received
func IncreaseFullContractsReceived() {
	noOfRequests.WithLabelValues("Received", "Full Contracts").Inc()
}

// This increases when the response of GetStorage in feeder.go is received
func IncreaseContractStorageReceived() {
	noOfRequests.WithLabelValues("Received", "Contract Storage").Inc()
}

// This increases when the response of GetTransactionStatus in feeder.go is received
func IncreaseTxStatusReceived() {
	noOfRequests.WithLabelValues("Received", "Transaction Status").Inc()
}

// This increases when the response of GetTransaction in feeder.go is received
func IncreaseTxReceived() {
	noOfRequests.WithLabelValues("Received", "Transaction").Inc()
}

// This increases when the response of GetTransactionReceipt in feeder.go is received
func IncreaseTxReceiptReceived() {
	noOfRequests.WithLabelValues("Received", "Transaction Receipt").Inc()
}

// This increases when the response of GetBlockHashById in feeder.go is received
func IncreaseBlockHashReceived() {
	noOfRequests.WithLabelValues("Received", "BlockHashByID").Inc()
}

// This increases when the response of GetBlockHashById in feeder.go is received
func IncreaseBlockIDReceived() {
	noOfRequests.WithLabelValues("Received", "BlockIDByHash").Inc()
}

// This increases when the response of GetBlockHashById in feeder.go is received
func IncreaseTxHashReceived() {
	noOfRequests.WithLabelValues("Received", "TransactionHashByID").Inc()
}

// This increases when the response of GetBlockHashById in feeder.go is received
func IncreaseTxIDReceived() {
	noOfRequests.WithLabelValues("Received", "TransactionIDByHash").Inc()
}

// This increases when the request in GetContractAddresses in feeder.go fails
func IncreaseContractAddressesFailed() {
	noOfRequests.WithLabelValues("Failed", "Contract Addresses").Inc()
}

// This increases when the request in CallContract in feeder.go fails
func IncreaseContractCallsFailed() {
	noOfRequests.WithLabelValues("Failed", "Contract Calls").Inc()
}

// This increases when the request in GetBlock in feeder.go fails
func IncreaseBlockFailed() {
	noOfRequests.WithLabelValues("Failed", "Blocks").Inc()
}

// This increases when the request in GetStateUpdate in feeder.go fails
func IncreaseStateUpdateFailed() {
	noOfRequests.WithLabelValues("Failed", "State Update").Inc()
}

// This increases when the request in GetFullContract in feeder.go fails
func IncreaseFullContractsFailed() {
	noOfRequests.WithLabelValues("Failed", "Full Contracts").Inc()
}

// This increases when the request in GetStorage in feeder.go fails
func IncreaseContractStorageFailed() {
	noOfRequests.WithLabelValues("Failed", "Contract Storage").Inc()
}

// This increases when the request in GetTransactionStatus in feeder.go fails
func IncreaseTxStatusFailed() {
	noOfRequests.WithLabelValues("Failed", "Transaction Status").Inc()
}

// This increases when the request in GetTransaction in feeder.go fails
func IncreaseTxFailed() {
	noOfRequests.WithLabelValues("Failed", "Transaction").Inc()
}

// This increases when the request in GetTransactionReceipt in feeder.go fails
func IncreaseTxReceiptFailed() {
	noOfRequests.WithLabelValues("Failed", "Transaction Receipt").Inc()
}

// This increases when the request in GetBlockHashById in feeder.go fails
func IncreaseBlockHashFailed() {
	noOfRequests.WithLabelValues("Failed", "BlockHashByID").Inc()
}

// This increases when the request in GetBlockHashById in feeder.go fails
func IncreaseBlockIDFailed() {
	noOfRequests.WithLabelValues("Failed", "BlockIDByHash").Inc()
}

// This increases when the request in GetBlockHashById in feeder.go fails
func IncreaseTxHashFailed() {
	noOfRequests.WithLabelValues("Failed", "TransactionHashByID").Inc()
}

// This increases when the request in GetBlockHashById in feeder.go fails
func IncreaseTxIDFailed() {
	noOfRequests.WithLabelValues("Failed", "TransactionIDByHash").Inc()
}

func main() {
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2048", nil)
}
