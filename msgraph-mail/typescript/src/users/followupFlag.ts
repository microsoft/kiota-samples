import {SerializationWriter, ParseNode, Parsable} from '@microsoft/kiota-abstractions';
import {DateTimeTimeZone} from './dateTimeTimeZone';
import {FollowupFlagStatus} from './followupFlagStatus';

export class FollowupFlag implements Parsable {
    /** Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.  */
    public readonly additionalData: Map<string, unknown> = new Map<string, unknown>();
    public completedDateTime?: DateTimeTimeZone | undefined;
    public dueDateTime?: DateTimeTimeZone | undefined;
    public flagStatus?: FollowupFlagStatus | undefined;
    public startDateTime?: DateTimeTimeZone | undefined;
    /**
     * The deserialization information for the current model
     * @returns a Map<string, (item: T, node: ParseNode) => void>
     */
    public getFieldDeserializers<T> () : Map<string, (item: T, node: ParseNode) => void> {
        return new Map<string, (item: T, node: ParseNode) => void>([
            ["completedDateTime", (o, n) => { (o as unknown as FollowupFlag).completedDateTime = n.getObjectValue<DateTimeTimeZone>(DateTimeTimeZone); }],
            ["dueDateTime", (o, n) => { (o as unknown as FollowupFlag).dueDateTime = n.getObjectValue<DateTimeTimeZone>(DateTimeTimeZone); }],
            ["flagStatus", (o, n) => { (o as unknown as FollowupFlag).flagStatus = n.getEnumValue<FollowupFlagStatus>(FollowupFlagStatus); }],
            ["startDateTime", (o, n) => { (o as unknown as FollowupFlag).startDateTime = n.getObjectValue<DateTimeTimeZone>(DateTimeTimeZone); }],
        ]);
    };
    /**
     * Serialiazes information the current object
     * @param writer Serialization writer to use to serialize this model
     * @returns a void
     */
    public serialize (writer: SerializationWriter) : void {
        writer.writeObjectValue<DateTimeTimeZone>("completedDateTime", this.completedDateTime);
        writer.writeObjectValue<DateTimeTimeZone>("dueDateTime", this.dueDateTime);
        writer.writeEnumValue<FollowupFlagStatus>("flagStatus", this.flagStatus);
        writer.writeObjectValue<DateTimeTimeZone>("startDateTime", this.startDateTime);
        writer.writeAdditionalData(this.additionalData);
    };
}
