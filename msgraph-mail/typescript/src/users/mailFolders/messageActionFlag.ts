import {SerializationWriter, ParseNode, Parsable} from '@microsoft/kiota-abstractions';

export class MessageActionFlag implements Parsable<MessageActionFlag> {
    public serialize (writer: SerializationWriter) : void {
    };
    public deserializeFields () : Map<string, (item: MessageActionFlag, node: ParseNode) => void> {
        return new Map<string, (item: MessageActionFlag, node: ParseNode) => void>([
        ]);
    };
}
