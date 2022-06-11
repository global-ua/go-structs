package structs

const (
	FieldKeyError           FieldKey = "error"
	FieldKeyErrorCode       FieldKey = "error_code"
	FieldApiAccessKeyRef    FieldKey = "api_access_key_ref"
	FieldBalanceRef         FieldKey = "balance_ref"
	FieldKeyMobilePhone     FieldKey = "mobile_phone"
	FieldKeyNotification    FieldKey = "notification"
	FieldKeyAllowedIP       FieldKey = "allowed_ip"
	FieldKeyCallbackUrl     FieldKey = "callback_url"
	FieldKeyCallbackResult  FieldKey = "callback_result"
	FieldKeyCallbackCounter FieldKey = "callback_counter"
	FieldKeyGmApiUrl        FieldKey = "gm_api_url"
	FieldKeyGmApiPoint      FieldKey = "gm_api_point"
	FieldKeyGmApiService    FieldKey = "gm_api_service"
	FieldKeyActionInitiator FieldKey = "action_initiator"

	/*
	  fields started with underscore is private
	  they will be hidden when serialize to JSON
	*/

	FieldKeySalt FieldKey = "_salt"
)
