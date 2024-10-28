package model

// SignatureRequestGetResponse models the response from signature request API endpoints
type SignatureRequestGetResponse struct {
	SignatureRequest SignatureRequestResponse `json:"signature_request"`
	Warnings         []WarningResponse        `json:"warnings,omitempty"` // A list of warnings.
}

// SignatureRequestResponse Contains information about a signature request.
type SignatureRequestResponse struct {
	// Whether this is a test signature request. Test requests have no legal value. Defaults to `false`.
	TestMode bool `json:"test_mode,omitempty"`
	// The id of the SignatureRequest.
	SignatureRequestId string `json:"signature_request_id,omitempty"`
	// The email address of the initiator of the SignatureRequest.
	RequesterEmailAddress string `json:"requester_email_address,omitempty"`
	// The title the specified Account uses for the SignatureRequest.
	Title *string `json:"title,omitempty"`
	// Default Label for account.
	OriginalTitle *string `json:"original_title,omitempty"`
	// The subject in the email that was initially sent to the signers.
	Subject string `json:"subject,omitempty"`
	// The custom message in the email that was initially sent to the signers.
	Message string `json:"message,omitempty"`
	// The metadata attached to the signature request.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// Time the signature request was created.
	CreatedAt *UnixTimestamp `json:"created_at,omitempty"`
	// The time when the signature request will expire unsigned signatures. See [Signature Request Expiration
	// Date](https://developers.hellosign.com/docs/signature-request/expiration/) for details.
	ExpiresAt *UnixTimestamp `json:"expires_at,omitempty"`
	// Whether or not the SignatureRequest has been fully executed by all signers.
	IsComplete *bool `json:"is_complete,omitempty"`
	// Whether or not the SignatureRequest has been declined by a signer.
	IsDeclined *bool `json:"is_declined,omitempty"`
	// Whether or not an error occurred (either during the creation of the SignatureRequest or during one of the signings).
	HasError *bool `json:"has_error,omitempty"`
	// The URL where a copy of the request's documents can be downloaded.
	FilesUrl *string `json:"files_url,omitempty"`
	// The URL where a signer, after authenticating, can sign the documents. This should only be used by users with existing Dropbox Sign
	// accounts as they will be required to log in before signing.
	SigningUrl string `json:"signing_url,omitempty"`
	// The URL where the requester and the signers can view the current status of the SignatureRequest.
	DetailsUrl *string `json:"details_url,omitempty"`
	// A list of email addresses that were CCed on the SignatureRequest. They will receive a copy of the final PDF once all the signers have
	// signed.
	CcEmailAddresses []string `json:"cc_email_addresses,omitempty"`
	// The URL you want the signer redirected to after they successfully sign.
	SigningRedirectUrl string `json:"signing_redirect_url,omitempty"`
	// The path where the completed document can be downloaded
	FinalCopyUri string `json:"final_copy_uri,omitempty"`
	// Templates IDs used in this SignatureRequest (if any).
	TemplateIds []string `json:"template_ids,omitempty"`
	// An array of Custom Field objects containing the name and type of each custom field.  * Text Field uses
	// `SignatureRequestResponseCustomFieldText` * Checkbox Field uses `SignatureRequestResponseCustomFieldCheckbox`
	CustomFields []SignatureRequestResponseCustomFieldBase `json:"custom_fields,omitempty"`
	// Signer attachments.
	Attachments []SignatureRequestResponseAttachment `json:"attachments,omitempty"`
	// An array of form field objects containing the name, value, and type of each textbox or checkmark field filled in by the signers.
	ResponseData []SignatureRequestResponseDataBase `json:"response_data,omitempty"`
	// An array of signature objects, 1 for each signer.
	Signatures []SignatureRequestResponseSignatures `json:"signatures,omitempty"`
	// The ID of the Bulk Send job which sent the signature request, if applicable.
	BulkSendJobId string `json:"bulk_send_job_id,omitempty"`
}

// SignatureRequestResponseCustomFieldBase An array of Custom Field objects containing the name and type of each custom field.  * Text Field
// uses `SignatureRequestResponseCustomFieldText` * Checkbox Field uses `SignatureRequestResponseCustomFieldCheckbox`
type SignatureRequestResponseCustomFieldBase struct {
	// The type of this Custom Field. Only 'text' and 'checkbox' are currently supported.
	Type string `json:"type"`
	// The name of the Custom Field.
	Name string `json:"name"`
	// A boolean value denoting if this field is required.
	Required *bool `json:"required,omitempty"`
	// The unique ID for this field.
	ApiId *string `json:"api_id,omitempty"`
	// The name of the Role that is able to edit this field.
	Editor *string `json:"editor,omitempty"`
}

// SignatureRequestResponseAttachment Signer attachments.
type SignatureRequestResponseAttachment struct {
	// The unique ID for this attachment.
	Id string `json:"id"`
	// The Signer this attachment is assigned to.
	Signer string `json:"signer"`
	// The name of this attachment.
	Name string `json:"name"`
	// A boolean value denoting if this attachment is required.
	Required bool `json:"required"`
	// Instructions for Signer.
	Instructions string `json:"instructions,omitempty"`
	// Timestamp when attachment was uploaded by Signer.
	UploadedAt *UnixTimestamp `json:"uploaded_at,omitempty"`
}

// SignatureRequestResponseDataBase An array of form field objects containing the name, value, and type of each textbox or checkmark field
// filled in by the signers.
type SignatureRequestResponseDataBase struct {
	// The unique ID for this field.
	ApiId *string `json:"api_id,omitempty"`
	// The ID of the signature to which this response is linked.
	SignatureId *string `json:"signature_id,omitempty"`
	// The name of the form field.
	Name *string `json:"name,omitempty"`
	// A boolean value denoting if this field is required.
	Required *bool   `json:"required,omitempty"`
	Type     *string `json:"type,omitempty"`
}

// SignatureRequestResponseSignatures An array of signature objects, 1 for each signer.
type SignatureRequestResponseSignatures struct {
	// Signature identifier.
	SignatureId *string `json:"signature_id,omitempty"`
	// Signer Group GUID
	SignerGroupGuid string `json:"signer_group_guid,omitempty"`
	// The email address of the signer.
	SignerEmailAddress string `json:"signer_email_address,omitempty"`
	// The name of the signer.
	SignerName string `json:"signer_name,omitempty"`
	// The role of the signer.
	SignerRole string `json:"signer_role,omitempty"`
	// If signer order is assigned this is the 0-based index for this signer.
	Order *int `json:"order,omitempty"`
	// The current status of the signature. eg: awaiting_signature, signed, declined.
	StatusCode *string `json:"status_code,omitempty"`
	// The reason provided by the signer for declining the request.
	DeclineReason string `json:"decline_reason,omitempty"`
	// Time that the document was signed or null.
	SignedAt *UnixTimestamp `json:"signed_at,omitempty"`
	// The time that the document was last viewed by this signer or null.
	LastViewedAt *UnixTimestamp `json:"last_viewed_at,omitempty"`
	// The time the last reminder email was sent to the signer or null.
	LastRemindedAt *UnixTimestamp `json:"last_reminded_at,omitempty"`
	// Boolean to indicate whether this signature requires a PIN to access.
	HasPin *bool `json:"has_pin,omitempty"`
	// Boolean to indicate whether this signature has SMS authentication enabled.
	HasSmsAuth bool `json:"has_sms_auth,omitempty"`
	// Boolean to indicate whether this signature has SMS delivery enabled.
	HasSmsDelivery bool `json:"has_sms_delivery,omitempty"`
	// The SMS phone number used for authentication or signature request delivery.
	SmsPhoneNumber string `json:"sms_phone_number,omitempty"`
	// Email address of original signer who reassigned to this signer.
	ReassignedBy string `json:"reassigned_by,omitempty"`
	// Reason provided by original signer who reassigned to this signer.
	ReassignmentReason string `json:"reassignment_reason,omitempty"`
	// Previous signature identifier.
	ReassignedFrom string `json:"reassigned_from,omitempty"`
	// Error message pertaining to this signer, or null.
	Error string `json:"error,omitempty"`
}

// WarningResponse is a warning
type WarningResponse struct {
	WarningMsg  string `json:"warning_msg"`  // Warning message
	WarningName string `json:"warning_name"` // Warning name
}
