package childfolders

import (
    ie51e765764484004639cdfca9cb330e622fbf042d77be5ba8e7a002ffc8f8ed7 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users"
)

type ChildFoldersResponse struct {
    additionalData map[string]interface{};
    nextLink *string;
    value []ie51e765764484004639cdfca9cb330e622fbf042d77be5ba8e7a002ffc8f8ed7.MailFolder;
}
