package error

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Detailable 
type Detailable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetCode()(*string)
    GetMessage()(*string)
    GetTarget()(*string)
    SetCode(value *string)()
    SetMessage(value *string)()
    SetTarget(value *string)()
}
