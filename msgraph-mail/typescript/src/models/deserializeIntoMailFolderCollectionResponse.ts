import {createMailFolderFromDiscriminatorValue} from './createMailFolderFromDiscriminatorValue';
import type {MailFolder} from './mailFolder';
import type {MailFolderCollectionResponse} from './mailFolderCollectionResponse';
import {serializeMailFolder} from './serializeMailFolder';
import type {AdditionalDataHolder, Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export function deserializeIntoMailFolderCollectionResponse(mailFolderCollectionResponse: MailFolderCollectionResponse | undefined = {} as MailFolderCollectionResponse) : Record<string, (node: ParseNode) => void> {
    return {
        "@odata.nextLink": n => { mailFolderCollectionResponse.odataNextLink = n.getStringValue(); },
        "value": n => { mailFolderCollectionResponse.value = n.getCollectionOfObjectValues<MailFolder>(createMailFolderFromDiscriminatorValue); },
    }
}
