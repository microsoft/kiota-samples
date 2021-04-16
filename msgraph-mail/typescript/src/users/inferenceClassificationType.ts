import {SerializationWriter, ParseNode, Parsable} from '@microsoft/kiota-abstractions';

export class InferenceClassificationType implements Parsable<InferenceClassificationType> {
    public serialize (writer: SerializationWriter) : void {
    };
    public deserializeFields () : Map<string, (item: InferenceClassificationType, node: ParseNode) => void> {
        return new Map<string, (item: InferenceClassificationType, node: ParseNode) => void>([
        ]);
    };
}
