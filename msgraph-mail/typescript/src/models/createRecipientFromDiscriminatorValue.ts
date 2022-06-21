import {RecipientImpl} from './index';
import {ParseNode} from '@microsoft/kiota-abstractions';

export function createRecipientFromDiscriminatorValue(parseNode: ParseNode | undefined) : RecipientImpl {
    if(!parseNode) throw new Error("parseNode cannot be undefined");
    return new RecipientImpl();
}
