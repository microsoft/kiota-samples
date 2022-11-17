<<<<<<< HEAD
import {DateTimeTimeZone} from './dateTimeTimeZone';
import {FollowupFlagStatus} from './followupFlagStatus';
import {AdditionalDataHolder, Parsable} from '@microsoft/kiota-abstractions';

export interface FollowupFlag extends Partial<AdditionalDataHolder>, Partial<Parsable> {
    /** Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well. */
    additionalData?: Record<string, unknown> | undefined;
    /** The completedDateTime property */
    completedDateTime?: DateTimeTimeZone | undefined;
    /** The dueDateTime property */
    dueDateTime?: DateTimeTimeZone | undefined;
    /** The flagStatus property */
    flagStatus?: FollowupFlagStatus | undefined;
    /** The startDateTime property */
    startDateTime?: DateTimeTimeZone | undefined;
=======
import {createDateTimeTimeZoneFromDiscriminatorValue} from './createDateTimeTimeZoneFromDiscriminatorValue';
import {DateTimeTimeZone} from './dateTimeTimeZone';
import {FollowupFlag} from './followupFlag';
import {FollowupFlagStatus} from './followupFlagStatus';
import {DateTimeTimeZoneImpl} from './index';
import {AdditionalDataHolder, Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export class FollowupFlagImpl implements FollowupFlag {
    /** Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well. */
    private _additionalData?: Record<string, unknown> | undefined;
    /** The completedDateTime property */
    private _completedDateTime?: DateTimeTimeZone | undefined;
    /** The dueDateTime property */
    private _dueDateTime?: DateTimeTimeZone | undefined;
    /** The flagStatus property */
    private _flagStatus?: FollowupFlagStatus | undefined;
    /** The startDateTime property */
    private _startDateTime?: DateTimeTimeZone | undefined;
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
     * Gets the completedDateTime property value. The completedDateTime property
     * @returns a DateTimeTimeZoneInterface
     */
    public get completedDateTime() {
        return this._completedDateTime;
    };
    /**
     * Sets the completedDateTime property value. The completedDateTime property
     * @param value Value to set for the completedDateTime property.
     */
    public set completedDateTime(value: DateTimeTimeZone | undefined) {
        if(value) {
            this._completedDateTime = value instanceof DateTimeTimeZoneImpl? value as DateTimeTimeZoneImpl: new DateTimeTimeZoneImpl(value);
        }
    };
    /**
     * Instantiates a new followupFlag and sets the default values.
     * @param followupFlagParameterValue 
     */
    public constructor(followupFlagParameterValue?: FollowupFlag | undefined) {
        this._additionalData = followupFlagParameterValue?.additionalData ? followupFlagParameterValue?.additionalData! : {};
        this._completedDateTime = followupFlagParameterValue?.completedDateTime;
        this._dueDateTime = followupFlagParameterValue?.dueDateTime;
        this._flagStatus = followupFlagParameterValue?.flagStatus;
        this._startDateTime = followupFlagParameterValue?.startDateTime;
    };
    /**
     * Gets the dueDateTime property value. The dueDateTime property
     * @returns a DateTimeTimeZoneInterface
     */
    public get dueDateTime() {
        return this._dueDateTime;
    };
    /**
     * Sets the dueDateTime property value. The dueDateTime property
     * @param value Value to set for the dueDateTime property.
     */
    public set dueDateTime(value: DateTimeTimeZone | undefined) {
        if(value) {
            this._dueDateTime = value instanceof DateTimeTimeZoneImpl? value as DateTimeTimeZoneImpl: new DateTimeTimeZoneImpl(value);
        }
    };
    /**
     * Gets the flagStatus property value. The flagStatus property
     * @returns a followupFlagStatus
     */
    public get flagStatus() {
        return this._flagStatus;
    };
    /**
     * Sets the flagStatus property value. The flagStatus property
     * @param value Value to set for the flagStatus property.
     */
    public set flagStatus(value: FollowupFlagStatus | undefined) {
        if(value) {
            this._flagStatus = value;
        }
    };
    /**
     * The deserialization information for the current model
     * @returns a Record<string, (node: ParseNode) => void>
     */
    public getFieldDeserializers() : Record<string, (node: ParseNode) => void> {
        return {
            "completedDateTime": n => { this.completedDateTime = n.getObjectValue<DateTimeTimeZoneImpl>(createDateTimeTimeZoneFromDiscriminatorValue); },
            "dueDateTime": n => { this.dueDateTime = n.getObjectValue<DateTimeTimeZoneImpl>(createDateTimeTimeZoneFromDiscriminatorValue); },
            "flagStatus": n => { this.flagStatus = n.getEnumValue<FollowupFlagStatus>(FollowupFlagStatus); },
            "startDateTime": n => { this.startDateTime = n.getObjectValue<DateTimeTimeZoneImpl>(createDateTimeTimeZoneFromDiscriminatorValue); },
        };
    };
    /**
     * Serializes information the current object
     * @param writer Serialization writer to use to serialize this model
     */
    public serialize(writer: SerializationWriter) : void {
        if(!writer) throw new Error("writer cannot be undefined");
        if(this.completedDateTime){
            writer.writeObjectValue<DateTimeTimeZoneImpl>("completedDateTime", (this.completedDateTime instanceof DateTimeTimeZoneImpl? this.completedDateTime as DateTimeTimeZoneImpl: new DateTimeTimeZoneImpl(this.completedDateTime)));
        }
        if(this.dueDateTime){
            writer.writeObjectValue<DateTimeTimeZoneImpl>("dueDateTime", (this.dueDateTime instanceof DateTimeTimeZoneImpl? this.dueDateTime as DateTimeTimeZoneImpl: new DateTimeTimeZoneImpl(this.dueDateTime)));
        }
        if(this.flagStatus){
            writer.writeEnumValue<FollowupFlagStatus>("flagStatus", this.flagStatus);
        }
        if(this.startDateTime){
            writer.writeObjectValue<DateTimeTimeZoneImpl>("startDateTime", (this.startDateTime instanceof DateTimeTimeZoneImpl? this.startDateTime as DateTimeTimeZoneImpl: new DateTimeTimeZoneImpl(this.startDateTime)));
        }
        writer.writeAdditionalData(this.additionalData);
    };
    /**
     * Gets the startDateTime property value. The startDateTime property
     * @returns a DateTimeTimeZoneInterface
     */
    public get startDateTime() {
        return this._startDateTime;
    };
    /**
     * Sets the startDateTime property value. The startDateTime property
     * @param value Value to set for the startDateTime property.
     */
    public set startDateTime(value: DateTimeTimeZone | undefined) {
        if(value) {
            this._startDateTime = value instanceof DateTimeTimeZoneImpl? value as DateTimeTimeZoneImpl: new DateTimeTimeZoneImpl(value);
        }
    };
>>>>>>> main
}
