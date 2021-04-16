import {SerializationWriter, ParseNode, Parsable} from '@microsoft/kiota-abstractions';

export class FollowupFlagStatus implements Parsable<FollowupFlagStatus> {
    public serialize (writer: SerializationWriter) : void {
    };
    public deserializeFields () : Map<string, (item: FollowupFlagStatus, node: ParseNode) => void> {
        return new Map<string, (item: FollowupFlagStatus, node: ParseNode) => void>([
        ]);
    };
}
