package odata

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Main 
type Main struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    code *string;
    // 
    details []Detail;
    // The structure of this object is service-specific
    innererror *Innererror;
    // 
    message *string;
    // 
    target *string;
}
// NewMain instantiates a new main and sets the default values.
func NewMain()(*Main) {
    m := &Main{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func CreateMainFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewMain(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Main) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCode gets the code property value. 
func (m *Main) GetCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.code
    }
}
// GetDetails gets the details property value. 
func (m *Main) GetDetails()([]Detail) {
    if m == nil {
        return nil
    } else {
        return m.details
    }
}
// GetInnererror gets the innererror property value. The structure of this object is service-specific
func (m *Main) GetInnererror()(*Innererror) {
    if m == nil {
        return nil
    } else {
        return m.innererror
    }
}
// GetMessage gets the message property value. 
func (m *Main) GetMessage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.message
    }
}
// GetTarget gets the target property value. 
func (m *Main) GetTarget()(*string) {
    if m == nil {
        return nil
    } else {
        return m.target
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Main) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["code"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCode(val)
        }
        return nil
    }
    res["details"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDetailFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Detail, len(val))
            for i, v := range val {
                res[i] = *(v.(*Detail))
            }
            m.SetDetails(res)
        }
        return nil
    }
    res["innererror"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateInnererrorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInnererror(val.(*Innererror))
        }
        return nil
    }
    res["message"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessage(val)
        }
        return nil
    }
    res["target"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTarget(val)
        }
        return nil
    }
    return res
}
func (m *Main) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *Main) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("code", m.GetCode())
        if err != nil {
            return err
        }
    }
    if m.GetDetails() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDetails()))
        for i, v := range m.GetDetails() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("details", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("innererror", m.GetInnererror())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("message", m.GetMessage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("target", m.GetTarget())
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
func (m *Main) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetCode sets the code property value. 
func (m *Main) SetCode(value *string)() {
    if m != nil {
        m.code = value
    }
}
// SetDetails sets the details property value. 
func (m *Main) SetDetails(value []Detail)() {
    if m != nil {
        m.details = value
    }
}
// SetInnererror sets the innererror property value. The structure of this object is service-specific
func (m *Main) SetInnererror(value *Innererror)() {
    if m != nil {
        m.innererror = value
    }
}
// SetMessage sets the message property value. 
func (m *Main) SetMessage(value *string)() {
    if m != nil {
        m.message = value
    }
}
// SetTarget sets the target property value. 
func (m *Main) SetTarget(value *string)() {
    if m != nil {
        m.target = value
    }
}
