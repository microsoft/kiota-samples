<<<<<<< HEAD
import {MailFolderCollectionResponseImpl} from './index';
import {ParseNode} from '@microsoft/kiota-abstractions';

export function createMailFolderCollectionResponseFromDiscriminatorValue(parseNode: ParseNode | undefined) : MailFolderCollectionResponseImpl {
    if(!parseNode) throw new Error("parseNode cannot be undefined");
    return new MailFolderCollectionResponseImpl();
=======
import {MailFolderCollectionResponse} from './index';
import {ParseNode} from '@microsoft/kiota-abstractions';

export function createMailFolderCollectionResponseFromDiscriminatorValue(parseNode: ParseNode | undefined) : MailFolderCollectionResponse {
    if(!parseNode) throw new Error("parseNode cannot be undefined");
    return new MailFolderCollectionResponse();
>>>>>>> main
}
