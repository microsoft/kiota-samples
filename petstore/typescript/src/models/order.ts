import {Order_status} from './order_status';
import {AdditionalDataHolder, Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export class Order implements AdditionalDataHolder, Parsable {
    /** Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well. */
    private _additionalData: Record<string, unknown>;
    /** The complete property */
    private _complete?: boolean | undefined;
    /** The id property */
    private _id?: number | undefined;
    /** The petId property */
    private _petId?: number | undefined;
    /** The quantity property */
    private _quantity?: number | undefined;
    /** The shipDate property */
    private _shipDate?: Date | undefined;
    /** Order Status */
    private _status?: Order_status | undefined;
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
    public set additionalData(value: Record<string, unknown>) {
        this._additionalData = value;
    };
    /**
     * Gets the complete property value. The complete property
     * @returns a boolean
     */
    public get complete() {
        return this._complete;
    };
    /**
     * Sets the complete property value. The complete property
     * @param value Value to set for the complete property.
     */
    public set complete(value: boolean | undefined) {
        this._complete = value;
    };
    /**
     * Instantiates a new Order and sets the default values.
     */
    public constructor() {
        this._additionalData = {};
    };
    /**
     * The deserialization information for the current model
     * @returns a Record<string, (node: ParseNode) => void>
     */
    public getFieldDeserializers() : Record<string, (node: ParseNode) => void> {
        return {
            "complete": n => { this.complete = n.getBooleanValue(); },
            "id": n => { this.id = n.getNumberValue(); },
            "petId": n => { this.petId = n.getNumberValue(); },
            "quantity": n => { this.quantity = n.getNumberValue(); },
            "shipDate": n => { this.shipDate = n.getDateValue(); },
            "status": n => { this.status = n.getEnumValue<Order_status>(Order_status); },
        };
    };
    /**
     * Gets the id property value. The id property
     * @returns a int64
     */
    public get id() {
        return this._id;
    };
    /**
     * Sets the id property value. The id property
     * @param value Value to set for the id property.
     */
    public set id(value: number | undefined) {
        this._id = value;
    };
    /**
     * Gets the petId property value. The petId property
     * @returns a int64
     */
    public get petId() {
        return this._petId;
    };
    /**
     * Sets the petId property value. The petId property
     * @param value Value to set for the petId property.
     */
    public set petId(value: number | undefined) {
        this._petId = value;
    };
    /**
     * Gets the quantity property value. The quantity property
     * @returns a integer
     */
    public get quantity() {
        return this._quantity;
    };
    /**
     * Sets the quantity property value. The quantity property
     * @param value Value to set for the quantity property.
     */
    public set quantity(value: number | undefined) {
        this._quantity = value;
    };
    /**
     * Serializes information the current object
     * @param writer Serialization writer to use to serialize this model
     */
    public serialize(writer: SerializationWriter) : void {
        if(!writer) throw new Error("writer cannot be undefined");
        writer.writeBooleanValue("complete", this.complete);
        writer.writeNumberValue("id", this.id);
        writer.writeNumberValue("petId", this.petId);
        writer.writeNumberValue("quantity", this.quantity);
        writer.writeDateValue("shipDate", this.shipDate);
        writer.writeEnumValue<Order_status>("status", this.status);
        writer.writeAdditionalData(this.additionalData);
    };
    /**
     * Gets the shipDate property value. The shipDate property
     * @returns a Date
     */
    public get shipDate() {
        return this._shipDate;
    };
    /**
     * Sets the shipDate property value. The shipDate property
     * @param value Value to set for the shipDate property.
     */
    public set shipDate(value: Date | undefined) {
        this._shipDate = value;
    };
    /**
     * Gets the status property value. Order Status
     * @returns a Order_status
     */
    public get status() {
        return this._status;
    };
    /**
     * Sets the status property value. Order Status
     * @param value Value to set for the status property.
     */
    public set status(value: Order_status | undefined) {
        this._status = value;
    };
}
