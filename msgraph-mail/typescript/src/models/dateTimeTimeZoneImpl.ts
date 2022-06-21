import {DateTimeTimeZone} from './dateTimeTimeZone';
import {AdditionalDataHolder, Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export class DateTimeTimeZoneImpl implements DateTimeTimeZone {
    /** Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well. */
    public additionalData: Record<string, unknown>;
    /** A single point of time in a combined date and time representation ({date}T{time}). For example, '2019-04-16T09:00:00'. */
    public dateTime?: string | undefined;
    /** Represents a time zone, for example, 'Pacific Standard Time'. See below for possible values. */
    public timeZone?: string | undefined;
    /**
     * Instantiates a new dateTimeTimeZone and sets the default values.
     * @param dateTimeTimeZoneParameterValue 
     */
    public constructor(dateTimeTimeZoneParameterValue?: DateTimeTimeZone | undefined) {
        this.additionalData = dateTimeTimeZoneParameterValue?.additionalData ? dateTimeTimeZoneParameterValue?.additionalData! : {};
        this.dateTime = dateTimeTimeZoneParameterValue?.dateTime;
        this.timeZone = dateTimeTimeZoneParameterValue?.timeZone;
    };
    /**
     * The deserialization information for the current model
     * @returns a Record<string, (node: ParseNode) => void>
     */
    public getFieldDeserializers() : Record<string, (node: ParseNode) => void> {
        return {
            "dateTime": n => { this.dateTime = n.getStringValue(); },
            "timeZone": n => { this.timeZone = n.getStringValue(); },
        };
    };
    /**
     * Serializes information the current object
     * @param writer Serialization writer to use to serialize this model
     */
    public serialize(writer: SerializationWriter) : void {
        if(!writer) throw new Error("writer cannot be undefined");
        if(this.dateTime){
            writer.writeStringValue("dateTime", this.dateTime);
        }
        if(this.timeZone){
            writer.writeStringValue("timeZone", this.timeZone);
        }
        writer.writeAdditionalData(this.additionalData);
    };
}
