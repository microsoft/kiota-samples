import {AdditionalDataHolder, Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export class Post implements AdditionalDataHolder, Parsable {
    /** Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well. */
    private _additionalData: Record<string, unknown>;
    /** The body property */
    private _body?: string | undefined;
    /** The id property */
    private _id?: number | undefined;
    /** The title property */
    private _title?: string | undefined;
    /** The userId property */
    private _userId?: number | undefined;
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
     * Gets the body property value. The body property
     * @returns a string
     */
    public get body() {
        return this._body;
    };
    /**
     * Sets the body property value. The body property
     * @param value Value to set for the body property.
     */
    public set body(value: string | undefined) {
        this._body = value;
    };
    /**
     * Instantiates a new Post and sets the default values.
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
            "body": n => { this.body = n.getStringValue(); },
            "id": n => { this.id = n.getNumberValue(); },
            "title": n => { this.title = n.getStringValue(); },
            "userId": n => { this.userId = n.getNumberValue(); },
        };
    };
    /**
     * Gets the id property value. The id property
     * @returns a integer
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
     * Serializes information the current object
     * @param writer Serialization writer to use to serialize this model
     */
    public serialize(writer: SerializationWriter) : void {
        if(!writer) throw new Error("writer cannot be undefined");
        writer.writeStringValue("body", this.body);
        writer.writeNumberValue("id", this.id);
        writer.writeStringValue("title", this.title);
        writer.writeNumberValue("userId", this.userId);
        writer.writeAdditionalData(this.additionalData);
    };
    /**
     * Gets the title property value. The title property
     * @returns a string
     */
    public get title() {
        return this._title;
    };
    /**
     * Sets the title property value. The title property
     * @param value Value to set for the title property.
     */
    public set title(value: string | undefined) {
        this._title = value;
    };
    /**
     * Gets the userId property value. The userId property
     * @returns a integer
     */
    public get userId() {
        return this._userId;
    };
    /**
     * Sets the userId property value. The userId property
     * @param value Value to set for the userId property.
     */
    public set userId(value: number | undefined) {
        this._userId = value;
    };
}
