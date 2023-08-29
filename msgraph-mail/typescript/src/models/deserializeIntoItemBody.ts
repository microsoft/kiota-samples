import {BodyType} from './bodyType';
import type {ItemBody} from './itemBody';
import type {AdditionalDataHolder, Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export function deserializeIntoItemBody(itemBody: ItemBody | undefined = {} as ItemBody) : Record<string, (node: ParseNode) => void> {
    return {
        "content": n => { itemBody.content = n.getStringValue(); },
        "contentType": n => { itemBody.contentType = n.getEnumValue<BodyType>(BodyType); },
    }
}
