package model

import (
	"os"
)

// CreateEmbeddedWithTemplateRequest struct for CreateEmbeddedWithTemplateRequest
type CreateEmbeddedWithTemplateRequest struct {
	// Client id of the app you're using to create this embedded signature request. Used for security purposes.
	ClientId string `json:"client_id"`
	// Use `template_ids` to create a SignatureRequest from one or more templates, in the order in which the template will be used.
	TemplateIds []string `json:"template_ids"`
	// Add Signers to your Templated-based Signature Request.
	Signers []SubSignatureRequestTemplateSigner `json:"signers"`
	// Allows signers to decline to sign a document if `true`. Defaults to `false`.
	AllowDecline bool `json:"allow_decline,omitempty"`
	// Add CC email recipients. Required when a CC role exists for the Template.
	CCs []SubCC `json:"ccs,omitempty"`
	// An array defining values and options for custom fields. Required when a custom field exists in the Template.
	CustomFields []SubCustomField `json:"custom_fields,omitempty"`
	// Use `files[]` to indicate the uploaded file(s) to send for signature.  This endpoint requires either **files** or **file_urls[]**,
	// but not both.
	Files []*os.File `json:"files,omitempty"`
	// Use `file_urls[]` to have Dropbox Sign download the file(s) to send for signature.  This endpoint requires either **files** or
	// **file_urls[]**, but not both.
	FileUrls []string `json:"file_urls,omitempty"`
	// The custom message in the email that will be sent to the signers.
	Message string `json:"message,omitempty"`
	// Key-value data that should be attached to the signature request. This metadata is included in all API responses and events involving
	// the signature request. For example, use the metadata field to store a signer's order number for look up when receiving events for the
	// signature request.  Each request can include up to 10 metadata keys (or 50 nested metadata keys), with key names up to 40 characters
	// long and values up to 1000 characters long.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// This allows the requester to specify the types allowed for creating a signature.
	SigningOptions *SubSigningOptions `json:"signing_options,omitempty"`
	// The subject in the email that will be sent to the signers.
	Subject string `json:"subject,omitempty"`
	// Whether this is a test, the signature request will not be legally binding if set to `true`. Defaults to `false`.
	TestMode bool `json:"test_mode,omitempty"`
	// The title you want to assign to the SignatureRequest.
	Title string `json:"title,omitempty"`
	// Controls whether [auto fill fields](https://faq.hellosign.com/hc/en-us/articles/360051467511-Auto-Fill-Fields) can automatically
	// populate a signer's information during signing.  **NOTE:** Keep your signer's information safe by ensuring that the _signer on your
	// signature request is the intended party_ before using this feature.
	PopulateAutoFillFields bool `json:"populate_auto_fill_fields,omitempty"`
}

// SubCC struct for SubCC
type SubCC struct {
	// Must match an existing CC role in chosen Template(s). Multiple CC recipients cannot share the same CC role.
	Role string `json:"role"`
	// The email address of the CC recipient.
	EmailAddress string `json:"email_address"`
}

// SubCustomField When used together with merge fields, `custom_fields` allows users to add pre-filled data to their signature requests.
// Pre-filled data can be used with \"send-once\" signature requests by adding merge fields with `form_fields_per_document` or [Text
// Tags](https://app.hellosign.com/api/textTagsWalkthrough#TextTagIntro) while passing values back with `custom_fields` together in one API
// call.  For using pre-filled on repeatable signature requests, merge fields are added to templates in the Dropbox Sign UI or by calling
// [/template/create_embedded_draft](/api/reference/operation/templateCreateEmbeddedDraft) and then passing `custom_fields` on subsequent
// signature requests referencing that template.
type SubCustomField struct {
	// Used to create editable merge fields. When the value matches a role passed in with `signers`, that role can edit the data that was
	// pre-filled to that field. This field is optional, but required when this custom field object is set to `required = true`.  **NOTE:**
	// Editable merge fields are only supported for single signer requests (or the first signer in ordered signature requests). If used when
	// there are multiple signers in an unordered signature request, the editor value is ignored and the field won't be editable.
	Editor *string `json:"editor,omitempty"`
	// The name of a custom field. When working with pre-filled data, the custom field's name must have a matching merge field name or the
	// field will remain empty on the document during signing.
	Name string `json:"name"`
	// Used to set an editable merge field when working with pre-filled data. When `true`, the custom field must specify a signer role in
	// `editor`.
	Required *bool `json:"required,omitempty"`
	// The string that resolves (aka \"pre-fills\") to the merge field on the final document(s) used for signing.
	Value *string `json:"value,omitempty"`
}

// SubSignatureRequestTemplateSigner struct for SubSignatureRequestTemplateSigner
type SubSignatureRequestTemplateSigner struct {
	// Must match an existing role in chosen Template(s). It's case-sensitive.
	Role string `json:"role"`
	// The name of the signer.
	Name string `json:"name"`
	// The email address of the signer.
	EmailAddress string `json:"email_address"`
	// The 4- to 12-character access code that will secure this signer's signature page.
	Pin *string `json:"pin,omitempty"`
	// An E.164 formatted phone number.  By using the feature, you agree you are responsible for obtaining a signer's consent to receive
	// text messages from Dropbox Sign related to this signature request and confirm you have obtained such consent from all signers prior
	// to enabling SMS delivery for this signature request. [Learn
	// more](https://faq.hellosign.com/hc/en-us/articles/15815316468877-Dropbox-Sign-SMS-tools-add-on).  **NOTE:** Not available in test
	// mode and requires a Standard plan or higher.
	SmsPhoneNumber *string `json:"sms_phone_number,omitempty"`
	// Specifies the feature used with the `sms_phone_number`. Default `authentication`.  If `authentication`, signer is sent a verification
	// code via SMS that is required to access the document.  If `delivery`, a link to complete the signature request is delivered via SMS
	// (_and_ email).
	SmsPhoneNumberType *string `json:"sms_phone_number_type,omitempty"`
}

// SubSigningOptions This allows the requester to specify the types allowed for creating a signature.  **NOTE:** If `signing_options` are
// not defined in the request, the allowed types will default to those specified in the account settings.
type SubSigningOptions struct {
	// The default type shown (limited to the listed types)
	DefaultType string `json:"default_type"`
	// Allows drawing the signature
	Draw *bool `json:"draw,omitempty"`
	// Allows using a smartphone to email the signature
	Phone *bool `json:"phone,omitempty"`
	// Allows typing the signature
	Type *bool `json:"type,omitempty"`
	// Allows uploading the signature
	Upload *bool `json:"upload,omitempty"`
}
