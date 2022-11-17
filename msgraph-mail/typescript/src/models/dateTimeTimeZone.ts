<<<<<<< HEAD
import {AdditionalDataHolder, Parsable} from '@microsoft/kiota-abstractions';

export interface DateTimeTimeZone extends Partial<AdditionalDataHolder>, Partial<Parsable> {
    /** Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well. */
    additionalData?: Record<string, unknown> | undefined;
    /** A single point of time in a combined date and time representation ({date}T{time}). For example, '2019-04-16T09:00:00'. */
    dateTime?: string | undefined;
    /** Represents a time zone, for example, 'Pacific Standard Time'. See below for possible values. */
    timeZone?: string | undefined;
=======
import {DateTimeTimeZone} from './dateTimeTimeZone';
import {AdditionalDataHolder, Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export class DateTimeTimeZoneImpl implements DateTimeTimeZone {
    /** Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well. */
    private _additionalData?: Record<string, unknown> | undefined;
    /** A single point of time in a combined date and time representation ({date}T{time}). For example, '2019-04-16T09:00:00'. */
    private _dateTime?: string | undefined;
    /** Represents a time zone, for example, 'Pacific Standard Time'. See below for possible values. */
    private _timeZone?: string | undefined;
    /**
     * Gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
     * @returns a Record<string, unknown>
     */
    public get additionalData() {
        return this._additionalData;
    };
    /**
     * Sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
     * @param value Value to set for the AdditionalData property.
     */
    public set additionalData(value: Record<string, unknown> | undefined) {
        if(value) {
            this._additionalData = value;
        }
    };
    /**
     * Instantiates a new dateTimeTimeZone and sets the default values.
     * @param dateTimeTimeZoneParameterValue 
     */
    public constructor(dateTimeTimeZoneParameterValue?: DateTimeTimeZone | undefined) {
        this._additionalData = dateTimeTimeZoneParameterValue?.additionalData ? dateTimeTimeZoneParameterValue?.additionalData! : {};
        this._dateTime = dateTimeTimeZoneParameterValue?.dateTime;
        this._timeZone = dateTimeTimeZoneParameterValue?.timeZone;
    };
    /**
     * Gets the dateTime property value. A single point of time in a combined date and time representation ({date}T{time}). For example, '2019-04-16T09:00:00'.
     * @returns a string
     */
    public get dateTime() {
        return this._dateTime;
    };
    /**
     * Sets the dateTime property value. A single point of time in a combined date and time representation ({date}T{time}). For example, '2019-04-16T09:00:00'.
     * @param value Value to set for the dateTime property.
     */
    public set dateTime(value: string | undefined) {
        if(value) {
            this._dateTime = value;
        }
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
    /**
     * Gets the timeZone property value. Represents a time zone, for example, 'Pacific Standard Time'. See below for possible values.
     * @returns a string
     */
    public get timeZone() {
        return this._timeZone;
    };
    /**
     * Sets the timeZone property value. Represents a time zone, for example, 'Pacific Standard Time'. See below for possible values.
     * @param value Value to set for the timeZone property.
     */
    public set timeZone(value: string | undefined) {
        if(value) {
            this._timeZone = value;
        }
    };
>>>>>>> main
}
