package odata

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    iecd558b5bae11a51f88fb38e0fb1e378f141a48ce05284655fe1ad361c3389cf "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/models/odata/error"
)

// Error 
type Error struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ApiError
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    error iecd558b5bae11a51f88fb38e0fb1e378f141a48ce05284655fe1ad361c3389cf.Mainable;
}
// NewError instantiates a new Error and sets the default values.
func NewError()(*Error) {
    m := &Error{
        ApiError: *ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewApiError(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func CreateErrorFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewError(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Error) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetError gets the error property value. 
func (m *Error) GetError()(iecd558b5bae11a51f88fb38e0fb1e378f141a48ce05284655fe1ad361c3389cf.Mainable) {
    if m == nil {
        return nil
    } else {
        return m.error
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Error) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["error"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(iecd558b5bae11a51f88fb38e0fb1e378f141a48ce05284655fe1ad361c3389cf.CreateMainFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetError(val.(iecd558b5bae11a51f88fb38e0fb1e378f141a48ce05284655fe1ad361c3389cf.Mainable))
        }
        return nil
    }
    return res
}
func (m *Error) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *Error) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("error", m.GetError())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Error) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetError sets the error property value. 
func (m *Error) SetError(value iecd558b5bae11a51f88fb38e0fb1e378f141a48ce05284655fe1ad361c3389cf.Mainable)() {
    if m != nil {
        m.error = value
    }
}
