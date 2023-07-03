import {createMailFolderFromDiscriminatorValue} from './createMailFolderFromDiscriminatorValue';
import {MailFolder} from './mailFolder';
import {MailFolderCollectionResponse} from './mailFolderCollectionResponse';
import {serializeMailFolder} from './serializeMailFolder';
import {AdditionalDataHolder, Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export function deserializeIntoMailFolderCollectionResponse(mailFolderCollectionResponse: MailFolderCollectionResponse | undefined = {} as MailFolderCollectionResponse) : Record<string, (node: ParseNode) => void> {
    return {
        "@odata.nextLink": n => { mailFolderCollectionResponse.odataNextLink = n.getStringValue(); },
        "value": n => { mailFolderCollectionResponse.value = n.getCollectionOfObjectValues<MailFolder>(createMailFolderFromDiscriminatorValue); },
    }
}
