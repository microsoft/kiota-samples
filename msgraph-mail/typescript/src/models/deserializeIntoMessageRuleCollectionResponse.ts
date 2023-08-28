import {createMessageRuleFromDiscriminatorValue} from './createMessageRuleFromDiscriminatorValue';
import type {MessageRule} from './messageRule';
import type {MessageRuleCollectionResponse} from './messageRuleCollectionResponse';
import {serializeMessageRule} from './serializeMessageRule';
import type {AdditionalDataHolder, Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export function deserializeIntoMessageRuleCollectionResponse(messageRuleCollectionResponse: MessageRuleCollectionResponse | undefined = {} as MessageRuleCollectionResponse) : Record<string, (node: ParseNode) => void> {
    return {
        "@odata.nextLink": n => { messageRuleCollectionResponse.odataNextLink = n.getStringValue(); },
        "value": n => { messageRuleCollectionResponse.value = n.getCollectionOfObjectValues<MessageRule>(createMessageRuleFromDiscriminatorValue); },
    }
}
