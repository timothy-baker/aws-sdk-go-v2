// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// An attachment with the specified ID could not be found.
type AttachmentIdNotFound struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *AttachmentIdNotFound) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *AttachmentIdNotFound) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *AttachmentIdNotFound) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "AttachmentIdNotFound"
	}
	return *e.ErrorCodeOverride
}
func (e *AttachmentIdNotFound) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The limit for the number of attachment sets created in a short period of time
// has been exceeded.
type AttachmentLimitExceeded struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *AttachmentLimitExceeded) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *AttachmentLimitExceeded) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *AttachmentLimitExceeded) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "AttachmentLimitExceeded"
	}
	return *e.ErrorCodeOverride
}
func (e *AttachmentLimitExceeded) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The expiration time of the attachment set has passed. The set expires 1 hour
// after it is created.
type AttachmentSetExpired struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *AttachmentSetExpired) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *AttachmentSetExpired) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *AttachmentSetExpired) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "AttachmentSetExpired"
	}
	return *e.ErrorCodeOverride
}
func (e *AttachmentSetExpired) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// An attachment set with the specified ID could not be found.
type AttachmentSetIdNotFound struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *AttachmentSetIdNotFound) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *AttachmentSetIdNotFound) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *AttachmentSetIdNotFound) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "AttachmentSetIdNotFound"
	}
	return *e.ErrorCodeOverride
}
func (e *AttachmentSetIdNotFound) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// A limit for the size of an attachment set has been exceeded. The limits are
// three attachments and 5 MB per attachment.
type AttachmentSetSizeLimitExceeded struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *AttachmentSetSizeLimitExceeded) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *AttachmentSetSizeLimitExceeded) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *AttachmentSetSizeLimitExceeded) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "AttachmentSetSizeLimitExceeded"
	}
	return *e.ErrorCodeOverride
}
func (e *AttachmentSetSizeLimitExceeded) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The case creation limit for the account has been exceeded.
type CaseCreationLimitExceeded struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *CaseCreationLimitExceeded) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *CaseCreationLimitExceeded) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *CaseCreationLimitExceeded) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "CaseCreationLimitExceeded"
	}
	return *e.ErrorCodeOverride
}
func (e *CaseCreationLimitExceeded) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The requested caseId couldn't be located.
type CaseIdNotFound struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *CaseIdNotFound) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *CaseIdNotFound) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *CaseIdNotFound) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "CaseIdNotFound"
	}
	return *e.ErrorCodeOverride
}
func (e *CaseIdNotFound) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The limit for the number of DescribeAttachment requests in a short period of
// time has been exceeded.
type DescribeAttachmentLimitExceeded struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *DescribeAttachmentLimitExceeded) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DescribeAttachmentLimitExceeded) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DescribeAttachmentLimitExceeded) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "DescribeAttachmentLimitExceeded"
	}
	return *e.ErrorCodeOverride
}
func (e *DescribeAttachmentLimitExceeded) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// An internal server error occurred.
type InternalServerError struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InternalServerError) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InternalServerError) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InternalServerError) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "InternalServerError"
	}
	return *e.ErrorCodeOverride
}
func (e *InternalServerError) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }
