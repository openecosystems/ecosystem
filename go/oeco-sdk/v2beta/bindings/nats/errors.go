package natsnodev1

import (
	"time"

	"connectrpc.com/connect"
	sdkv2betalib "github.com/openecosystems/ecosystem/go/oeco-sdk/v2beta"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/protobuf/types/known/durationpb"
)

var (
	ErrTimeout      sdkv2betalib.SpecErrorable = sdkv2betalib.NewSpecError(connect.CodeDeadlineExceeded, "The request timed out")
	ErrNoResponders sdkv2betalib.SpecErrorable = sdkv2betalib.NewSpecError(connect.CodeUnavailable, "The resource is not available").WithRetryInfo(&errdetails.RetryInfo{
		RetryDelay: durationpb.New(10 * time.Second),
	})
)
