import {BodyType} from './bodyType';
import {ItemBody} from './itemBody';
import {AdditionalDataHolder, Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export function serializeItemBody(writer: SerializationWriter, itemBody: ItemBody | undefined = {} as ItemBody) : void {
        writer.writeStringValue("content", itemBody.content);
        writer.writeEnumValue<BodyType>("contentType", itemBody.contentType);
        writer.writeAdditionalData(itemBody.additionalData);
}
