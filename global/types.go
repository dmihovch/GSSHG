package global

type JSONPayload struct {
	Type string `json:"type"` //type determines which handler is used (eg raise, fold, check, flip)
	Data string `json:"data"` //determines quantity, whatever else.
}
