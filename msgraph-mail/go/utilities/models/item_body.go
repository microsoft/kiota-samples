package models
type ItemBody struct {
    additionalData map[string]interface{};
    content *string;
    contentType BodyType;
}
