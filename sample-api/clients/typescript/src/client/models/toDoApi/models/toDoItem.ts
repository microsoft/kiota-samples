import {AdditionalDataHolder, Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export class ToDoItem implements AdditionalDataHolder, Parsable {
    /** Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well. */
    private _additionalData: Record<string, unknown>;
    /** The dueDate property */
    private _dueDate?: Date | undefined;
    /** The id property */
    private _id?: string | undefined;
    /** The isComplete property */
    private _isComplete?: boolean | undefined;
    /** The name property */
    private _name?: string | undefined;
    /** The priority property */
    private _priority?: number | undefined;
    /** The type property */
    private _type?: string | undefined;
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
     * Instantiates a new ToDoItem and sets the default values.
     */
    public constructor() {
        this._additionalData = {};
    };
    /**
     * Gets the dueDate property value. The dueDate property
     * @returns a Date
     */
    public get dueDate() {
        return this._dueDate;
    };
    /**
     * Sets the dueDate property value. The dueDate property
     * @param value Value to set for the dueDate property.
     */
    public set dueDate(value: Date | undefined) {
        this._dueDate = value;
    };
    /**
     * The deserialization information for the current model
     * @returns a Record<string, (node: ParseNode) => void>
     */
    public getFieldDeserializers() : Record<string, (node: ParseNode) => void> {
        return {
            "dueDate": n => { this.dueDate = n.getDateValue(); },
            "id": n => { this.id = n.getStringValue(); },
            "isComplete": n => { this.isComplete = n.getBooleanValue(); },
            "name": n => { this.name = n.getStringValue(); },
            "priority": n => { this.priority = n.getNumberValue(); },
            "type": n => { this.type = n.getStringValue(); },
        };
    };
    /**
     * Gets the id property value. The id property
     * @returns a Guid
     */
    public get id() {
        return this._id;
    };
    /**
     * Sets the id property value. The id property
     * @param value Value to set for the id property.
     */
    public set id(value: string | undefined) {
        this._id = value;
    };
    /**
     * Gets the isComplete property value. The isComplete property
     * @returns a boolean
     */
    public get isComplete() {
        return this._isComplete;
    };
    /**
     * Sets the isComplete property value. The isComplete property
     * @param value Value to set for the isComplete property.
     */
    public set isComplete(value: boolean | undefined) {
        this._isComplete = value;
    };
    /**
     * Gets the name property value. The name property
     * @returns a string
     */
    public get name() {
        return this._name;
    };
    /**
     * Sets the name property value. The name property
     * @param value Value to set for the name property.
     */
    public set name(value: string | undefined) {
        this._name = value;
    };
    /**
     * Gets the priority property value. The priority property
     * @returns a integer
     */
    public get priority() {
        return this._priority;
    };
    /**
     * Sets the priority property value. The priority property
     * @param value Value to set for the priority property.
     */
    public set priority(value: number | undefined) {
        this._priority = value;
    };
    /**
     * Serializes information the current object
     * @param writer Serialization writer to use to serialize this model
     */
    public serialize(writer: SerializationWriter) : void {
        if(!writer) throw new Error("writer cannot be undefined");
        writer.writeDateValue("dueDate", this.dueDate);
        writer.writeStringValue("id", this.id);
        writer.writeBooleanValue("isComplete", this.isComplete);
        writer.writeStringValue("name", this.name);
        writer.writeNumberValue("priority", this.priority);
        writer.writeStringValue("type", this.type);
        writer.writeAdditionalData(this.additionalData);
    };
    /**
     * Gets the type property value. The type property
     * @returns a string
     */
    public get type() {
        return this._type;
    };
    /**
     * Sets the type property value. The type property
     * @param value Value to set for the type property.
     */
    public set type(value: string | undefined) {
        this._type = value;
    };
}
