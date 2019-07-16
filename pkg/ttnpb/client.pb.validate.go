// Code generated by protoc-gen-fieldmask. DO NOT EDIT.

package ttnpb

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/gogo/protobuf/types"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = types.DynamicAny{}
)

// ValidateFields checks the field values on Client with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *Client) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = ClientFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "ids":

			if v, ok := interface{}(&m.ClientIdentifiers).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return ClientValidationError{
						field:  "ids",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "created_at":

			if v, ok := interface{}(&m.CreatedAt).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return ClientValidationError{
						field:  "created_at",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "updated_at":

			if v, ok := interface{}(&m.UpdatedAt).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return ClientValidationError{
						field:  "updated_at",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "name":

			if utf8.RuneCountInString(m.GetName()) > 50 {
				return ClientValidationError{
					field:  "name",
					reason: "value length must be at most 50 runes",
				}
			}

		case "description":

			if utf8.RuneCountInString(m.GetDescription()) > 2000 {
				return ClientValidationError{
					field:  "description",
					reason: "value length must be at most 2000 runes",
				}
			}

		case "attributes":

			for key, val := range m.GetAttributes() {
				_ = val

				if utf8.RuneCountInString(key) > 36 {
					return ClientValidationError{
						field:  fmt.Sprintf("attributes[%v]", key),
						reason: "value length must be at most 36 runes",
					}
				}

				if !_Client_Attributes_Pattern.MatchString(key) {
					return ClientValidationError{
						field:  fmt.Sprintf("attributes[%v]", key),
						reason: "value does not match regex pattern \"^[a-z0-9](?:[-]?[a-z0-9]){2,}$\"",
					}
				}

				// no validation rules for Attributes[key]
			}

		case "contact_info":

			for idx, item := range m.GetContactInfo() {
				_, _ = idx, item

				if v, ok := interface{}(item).(interface{ ValidateFields(...string) error }); ok {
					if err := v.ValidateFields(subs...); err != nil {
						return ClientValidationError{
							field:  fmt.Sprintf("contact_info[%v]", idx),
							reason: "embedded message failed validation",
							cause:  err,
						}
					}
				}

			}

		case "secret":
			// no validation rules for Secret
		case "redirect_uris":

		case "state":

			if _, ok := State_name[int32(m.GetState())]; !ok {
				return ClientValidationError{
					field:  "state",
					reason: "value must be one of the defined enum values",
				}
			}

		case "skip_authorization":
			// no validation rules for SkipAuthorization
		case "endorsed":
			// no validation rules for Endorsed
		case "grants":

			for idx, item := range m.GetGrants() {
				_, _ = idx, item

				if _, ok := GrantType_name[int32(item)]; !ok {
					return ClientValidationError{
						field:  fmt.Sprintf("grants[%v]", idx),
						reason: "value must be one of the defined enum values",
					}
				}

			}

		case "rights":

			for idx, item := range m.GetRights() {
				_, _ = idx, item

				if _, ok := Right_name[int32(item)]; !ok {
					return ClientValidationError{
						field:  fmt.Sprintf("rights[%v]", idx),
						reason: "value must be one of the defined enum values",
					}
				}

			}

		default:
			return ClientValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// ClientValidationError is the validation error returned by
// Client.ValidateFields if the designated constraints aren't met.
type ClientValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ClientValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ClientValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ClientValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ClientValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ClientValidationError) ErrorName() string { return "ClientValidationError" }

// Error satisfies the builtin error interface
func (e ClientValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sClient.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ClientValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ClientValidationError{}

var _Client_Attributes_Pattern = regexp.MustCompile("^[a-z0-9](?:[-]?[a-z0-9]){2,}$")

// ValidateFields checks the field values on Clients with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *Clients) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = ClientsFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "clients":

			for idx, item := range m.GetClients() {
				_, _ = idx, item

				if v, ok := interface{}(item).(interface{ ValidateFields(...string) error }); ok {
					if err := v.ValidateFields(subs...); err != nil {
						return ClientsValidationError{
							field:  fmt.Sprintf("clients[%v]", idx),
							reason: "embedded message failed validation",
							cause:  err,
						}
					}
				}

			}

		default:
			return ClientsValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// ClientsValidationError is the validation error returned by
// Clients.ValidateFields if the designated constraints aren't met.
type ClientsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ClientsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ClientsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ClientsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ClientsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ClientsValidationError) ErrorName() string { return "ClientsValidationError" }

// Error satisfies the builtin error interface
func (e ClientsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sClients.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ClientsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ClientsValidationError{}

// ValidateFields checks the field values on GetClientRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetClientRequest) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = GetClientRequestFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "client_ids":

			if v, ok := interface{}(&m.ClientIdentifiers).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return GetClientRequestValidationError{
						field:  "client_ids",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "field_mask":

			if v, ok := interface{}(&m.FieldMask).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return GetClientRequestValidationError{
						field:  "field_mask",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		default:
			return GetClientRequestValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// GetClientRequestValidationError is the validation error returned by
// GetClientRequest.ValidateFields if the designated constraints aren't met.
type GetClientRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetClientRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetClientRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetClientRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetClientRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetClientRequestValidationError) ErrorName() string { return "GetClientRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetClientRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetClientRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetClientRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetClientRequestValidationError{}

// ValidateFields checks the field values on ListClientsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListClientsRequest) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = ListClientsRequestFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "collaborator":

			if v, ok := interface{}(m.GetCollaborator()).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return ListClientsRequestValidationError{
						field:  "collaborator",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "field_mask":

			if v, ok := interface{}(&m.FieldMask).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return ListClientsRequestValidationError{
						field:  "field_mask",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "order":
			// no validation rules for Order
		case "limit":

			if m.GetLimit() > 1000 {
				return ListClientsRequestValidationError{
					field:  "limit",
					reason: "value must be less than or equal to 1000",
				}
			}

		case "page":
			// no validation rules for Page
		default:
			return ListClientsRequestValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// ListClientsRequestValidationError is the validation error returned by
// ListClientsRequest.ValidateFields if the designated constraints aren't met.
type ListClientsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListClientsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListClientsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListClientsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListClientsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListClientsRequestValidationError) ErrorName() string {
	return "ListClientsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListClientsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListClientsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListClientsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListClientsRequestValidationError{}

// ValidateFields checks the field values on CreateClientRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateClientRequest) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = CreateClientRequestFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "client":

			if v, ok := interface{}(&m.Client).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return CreateClientRequestValidationError{
						field:  "client",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "collaborator":

			if v, ok := interface{}(&m.Collaborator).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return CreateClientRequestValidationError{
						field:  "collaborator",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		default:
			return CreateClientRequestValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// CreateClientRequestValidationError is the validation error returned by
// CreateClientRequest.ValidateFields if the designated constraints aren't met.
type CreateClientRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateClientRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateClientRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateClientRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateClientRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateClientRequestValidationError) ErrorName() string {
	return "CreateClientRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateClientRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateClientRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateClientRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateClientRequestValidationError{}

// ValidateFields checks the field values on UpdateClientRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpdateClientRequest) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = UpdateClientRequestFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "client":

			if v, ok := interface{}(&m.Client).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return UpdateClientRequestValidationError{
						field:  "client",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "field_mask":

			if v, ok := interface{}(&m.FieldMask).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return UpdateClientRequestValidationError{
						field:  "field_mask",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		default:
			return UpdateClientRequestValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// UpdateClientRequestValidationError is the validation error returned by
// UpdateClientRequest.ValidateFields if the designated constraints aren't met.
type UpdateClientRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateClientRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateClientRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateClientRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateClientRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateClientRequestValidationError) ErrorName() string {
	return "UpdateClientRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateClientRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateClientRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateClientRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateClientRequestValidationError{}

// ValidateFields checks the field values on ListClientCollaboratorsRequest
// with the rules defined in the proto definition for this message. If any
// rules are violated, an error is returned.
func (m *ListClientCollaboratorsRequest) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = ListClientCollaboratorsRequestFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "client_ids":

			if v, ok := interface{}(&m.ClientIdentifiers).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return ListClientCollaboratorsRequestValidationError{
						field:  "client_ids",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "limit":

			if m.GetLimit() > 1000 {
				return ListClientCollaboratorsRequestValidationError{
					field:  "limit",
					reason: "value must be less than or equal to 1000",
				}
			}

		case "page":
			// no validation rules for Page
		default:
			return ListClientCollaboratorsRequestValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// ListClientCollaboratorsRequestValidationError is the validation error
// returned by ListClientCollaboratorsRequest.ValidateFields if the designated
// constraints aren't met.
type ListClientCollaboratorsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListClientCollaboratorsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListClientCollaboratorsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListClientCollaboratorsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListClientCollaboratorsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListClientCollaboratorsRequestValidationError) ErrorName() string {
	return "ListClientCollaboratorsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListClientCollaboratorsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListClientCollaboratorsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListClientCollaboratorsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListClientCollaboratorsRequestValidationError{}

// ValidateFields checks the field values on GetClientCollaboratorRequest with
// the rules defined in the proto definition for this message. If any rules
// are violated, an error is returned.
func (m *GetClientCollaboratorRequest) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = GetClientCollaboratorRequestFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "client_ids":

			if v, ok := interface{}(&m.ClientIdentifiers).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return GetClientCollaboratorRequestValidationError{
						field:  "client_ids",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "collaborator":

			if v, ok := interface{}(&m.OrganizationOrUserIdentifiers).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return GetClientCollaboratorRequestValidationError{
						field:  "collaborator",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		default:
			return GetClientCollaboratorRequestValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// GetClientCollaboratorRequestValidationError is the validation error returned
// by GetClientCollaboratorRequest.ValidateFields if the designated
// constraints aren't met.
type GetClientCollaboratorRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetClientCollaboratorRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetClientCollaboratorRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetClientCollaboratorRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetClientCollaboratorRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetClientCollaboratorRequestValidationError) ErrorName() string {
	return "GetClientCollaboratorRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetClientCollaboratorRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetClientCollaboratorRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetClientCollaboratorRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetClientCollaboratorRequestValidationError{}

// ValidateFields checks the field values on SetClientCollaboratorRequest with
// the rules defined in the proto definition for this message. If any rules
// are violated, an error is returned.
func (m *SetClientCollaboratorRequest) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = SetClientCollaboratorRequestFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "client_ids":

			if v, ok := interface{}(&m.ClientIdentifiers).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return SetClientCollaboratorRequestValidationError{
						field:  "client_ids",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "collaborator":

			if v, ok := interface{}(&m.Collaborator).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return SetClientCollaboratorRequestValidationError{
						field:  "collaborator",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		default:
			return SetClientCollaboratorRequestValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// SetClientCollaboratorRequestValidationError is the validation error returned
// by SetClientCollaboratorRequest.ValidateFields if the designated
// constraints aren't met.
type SetClientCollaboratorRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SetClientCollaboratorRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SetClientCollaboratorRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SetClientCollaboratorRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SetClientCollaboratorRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SetClientCollaboratorRequestValidationError) ErrorName() string {
	return "SetClientCollaboratorRequestValidationError"
}

// Error satisfies the builtin error interface
func (e SetClientCollaboratorRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSetClientCollaboratorRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SetClientCollaboratorRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SetClientCollaboratorRequestValidationError{}
