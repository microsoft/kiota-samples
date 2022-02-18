package error

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Mainable 
type Mainable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetCode()(*string)
    GetDetails()([]Detailable)
    GetInnererror()(Innererrorable)
    GetMessage()(*string)
    GetTarget()(*string)
    SetCode(value *string)()
    SetDetails(value []Detailable)()
    SetInnererror(value Innererrorable)()
    SetMessage(value *string)()
    SetTarget(value *string)()
}
