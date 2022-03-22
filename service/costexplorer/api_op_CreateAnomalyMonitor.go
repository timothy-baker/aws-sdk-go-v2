// Code generated by smithy-go-codegen DO NOT EDIT.

package costexplorer

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new cost anomaly detection monitor with the requested type and monitor
// specification.
func (c *Client) CreateAnomalyMonitor(ctx context.Context, params *CreateAnomalyMonitorInput, optFns ...func(*Options)) (*CreateAnomalyMonitorOutput, error) {
	if params == nil {
		params = &CreateAnomalyMonitorInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateAnomalyMonitor", params, optFns, c.addOperationCreateAnomalyMonitorMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateAnomalyMonitorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateAnomalyMonitorInput struct {

	// The cost anomaly detection monitor object that you want to create.
	//
	// This member is required.
	AnomalyMonitor *types.AnomalyMonitor

	// An optional list of tags to associate with the specified AnomalyMonitor
	// (https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_AnomalyMonitor.html).
	// You can use resource tags to control access to your monitor using IAM policies.
	// Each tag consists of a key and a value, and each key must be unique for the
	// resource. The following restrictions apply to resource tags:
	//
	// * Although the
	// maximum number of array members is 200, you can assign a maximum of 50 user-tags
	// to one resource. The remaining are reserved for Amazon Web Services use
	//
	// * The
	// maximum length of a key is 128 characters
	//
	// * The maximum length of a value is
	// 256 characters
	//
	// * Valid characters for keys and values are: A-Z, a-z, spaces,
	// _.:/=+-
	//
	// * Keys and values are case sensitive
	//
	// * Keys and values are trimmed for
	// any leading or trailing whitespaces
	//
	// * Don’t use aws: as a prefix for your keys.
	// This prefix is reserved for Amazon Web Services use
	ResourceTags []types.ResourceTag

	noSmithyDocumentSerde
}

type CreateAnomalyMonitorOutput struct {

	// The unique identifier of your newly created cost anomaly detection monitor.
	//
	// This member is required.
	MonitorArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateAnomalyMonitorMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateAnomalyMonitor{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateAnomalyMonitor{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpCreateAnomalyMonitorValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateAnomalyMonitor(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opCreateAnomalyMonitor(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ce",
		OperationName: "CreateAnomalyMonitor",
	}
}
