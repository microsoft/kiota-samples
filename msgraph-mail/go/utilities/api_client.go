package utilities

import (
    url "net/url"
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    ie51e765764484004639cdfca9cb330e622fbf042d77be5ba8e7a002ffc8f8ed7 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users"
    i993da4875956788f567345c5c97cefaefa8f33555fde8bb29d1fd49e374ccccd "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/item"
)

type ApiClient struct {
    httpCore ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.HttpCore;
    pathSegment string;
    Users *ie51e765764484004639cdfca9cb330e622fbf042d77be5ba8e7a002ffc8f8ed7.UsersRequestBuilder;
}
func NewApiClient(httpCore ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.HttpCore)(*ApiClient) {
    return nil
}
func (m *ApiClient) UsersById(id string)(i993da4875956788f567345c5c97cefaefa8f33555fde8bb29d1fd49e374ccccd.UserRequestBuilder) {
    return i993da4875956788f567345c5c97cefaefa8f33555fde8bb29d1fd49e374ccccd.UserRequestBuilder.NewUserRequestBuilder(m.pathSegment + "/users/" + id, m.httpCore);
}
