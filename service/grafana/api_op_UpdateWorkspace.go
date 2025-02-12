// Code generated by smithy-go-codegen DO NOT EDIT.

package grafana

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/grafana/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Modifies an existing Amazon Managed Grafana workspace. If you use this operation
// and omit any optional parameters, the existing values of those parameters are
// not changed. To modify the user authentication methods that the workspace uses,
// such as SAML or IAM Identity Center, use UpdateWorkspaceAuthentication
// (https://docs.aws.amazon.com/grafana/latest/APIReference/API_UpdateWorkspaceAuthentication.html).
// To modify which users in the workspace have the Admin and Editor Grafana roles,
// use UpdatePermissions
// (https://docs.aws.amazon.com/grafana/latest/APIReference/API_UpdatePermissions.html).
func (c *Client) UpdateWorkspace(ctx context.Context, params *UpdateWorkspaceInput, optFns ...func(*Options)) (*UpdateWorkspaceOutput, error) {
	if params == nil {
		params = &UpdateWorkspaceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateWorkspace", params, optFns, c.addOperationUpdateWorkspaceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateWorkspaceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateWorkspaceInput struct {

	// The ID of the workspace to update.
	//
	// This member is required.
	WorkspaceId *string

	// Specifies whether the workspace can access Amazon Web Services resources in this
	// Amazon Web Services account only, or whether it can also access Amazon Web
	// Services resources in other accounts in the same organization. If you specify
	// ORGANIZATION, you must specify which organizational units the workspace can
	// access in the workspaceOrganizationalUnits parameter.
	AccountAccessType types.AccountAccessType

	// The configuration settings for network access to your workspace. When this is
	// configured, only listed IP addresses and VPC endpoints will be able to access
	// your workspace. Standard Grafana authentication and authorization will still be
	// required. If this is not configured, or is removed, then all IP addresses and
	// VPC endpoints will be allowed. Standard Grafana authentication and authorization
	// will still be required.
	NetworkAccessControl *types.NetworkAccessConfiguration

	// The name of an IAM role that already exists to use to access resources through
	// Organizations.
	OrganizationRoleName *string

	// If you specify SERVICE_MANAGED, Amazon Managed Grafana automatically creates the
	// IAM roles and provisions the permissions that the workspace needs to use Amazon
	// Web Services data sources and notification channels. If you specify
	// CUSTOMER_MANAGED, you will manage those roles and permissions yourself. If you
	// are creating this workspace in a member account of an organization and that
	// account is not a delegated administrator account, and you want the workspace to
	// access data sources in other Amazon Web Services accounts in the organization,
	// you must choose CUSTOMER_MANAGED. For more information, see Amazon Managed
	// Grafana permissions and policies for Amazon Web Services data sources and
	// notification channels
	// (https://docs.aws.amazon.com/grafana/latest/userguide/AMG-manage-permissions.html)
	PermissionType types.PermissionType

	// Whether to remove the network access configuration from the workspace. Setting
	// this to true and providing a networkAccessControl to set will return an error.
	// If you remove this configuration by setting this to true, then all IP addresses
	// and VPC endpoints will be allowed. Standard Grafana authentication and
	// authorization will still be required.
	RemoveNetworkAccessConfiguration *bool

	// Whether to remove the VPC configuration from the workspace. Setting this to true
	// and providing a vpcConfiguration to set will return an error.
	RemoveVpcConfiguration *bool

	// The name of the CloudFormation stack set to use to generate IAM roles to be used
	// for this workspace.
	StackSetName *string

	// The configuration settings for an Amazon VPC that contains data sources for your
	// Grafana workspace to connect to.
	VpcConfiguration *types.VpcConfiguration

	// Specify the Amazon Web Services data sources that you want to be queried in this
	// workspace. Specifying these data sources here enables Amazon Managed Grafana to
	// create IAM roles and permissions that allow Amazon Managed Grafana to read data
	// from these sources. You must still add them as data sources in the Grafana
	// console in the workspace. If you don't specify a data source here, you can still
	// add it as a data source later in the workspace console. However, you will then
	// have to manually configure permissions for it.
	WorkspaceDataSources []types.DataSourceType

	// A description for the workspace. This is used only to help you identify this
	// workspace.
	WorkspaceDescription *string

	// A new name for the workspace to update.
	WorkspaceName *string

	// Specify the Amazon Web Services notification channels that you plan to use in
	// this workspace. Specifying these data sources here enables Amazon Managed
	// Grafana to create IAM roles and permissions that allow Amazon Managed Grafana to
	// use these channels.
	WorkspaceNotificationDestinations []types.NotificationDestinationType

	// Specifies the organizational units that this workspace is allowed to use data
	// sources from, if this workspace is in an account that is part of an
	// organization.
	WorkspaceOrganizationalUnits []string

	// The workspace needs an IAM role that grants permissions to the Amazon Web
	// Services resources that the workspace will view data from. If you already have a
	// role that you want to use, specify it here. If you omit this field and you
	// specify some Amazon Web Services resources in workspaceDataSources or
	// workspaceNotificationDestinations, a new IAM role with the necessary permissions
	// is automatically created.
	WorkspaceRoleArn *string

	noSmithyDocumentSerde
}

type UpdateWorkspaceOutput struct {

	// A structure containing data about the workspace that was created.
	//
	// This member is required.
	Workspace *types.WorkspaceDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateWorkspaceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateWorkspace{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateWorkspace{}, middleware.After)
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
	if err = addOpUpdateWorkspaceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateWorkspace(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateWorkspace(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "grafana",
		OperationName: "UpdateWorkspace",
	}
}
