import {Detail} from './detail';
import {Innererror} from './innererror';
import {Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export class Main implements Parsable {
    /** Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.  */
    private _additionalData: Map<string, unknown>;
    private _code?: string | undefined;
    private _details?: Detail[] | undefined;
    /** The structure of this object is service-specific  */
    private _innererror?: Innererror | undefined;
    private _message?: string | undefined;
    private _target?: string | undefined;
    /**
     * Instantiates a new main and sets the default values.
     */
    public constructor() {
        this._additionalData = new Map<string, unknown>();
    };
    /**
     * Gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
     * @returns a Map<string, unknown>
     */
    public get additionalData() {
        return this._additionalData;
    };
    /**
     * Gets the code property value. 
     * @returns a string
     */
    public get code() {
        return this._code;
    };
    /**
     * Gets the details property value. 
     * @returns a detail
     */
    public get details() {
        return this._details;
    };
    /**
     * Gets the innererror property value. The structure of this object is service-specific
     * @returns a innererror
     */
    public get innererror() {
        return this._innererror;
    };
    /**
     * Gets the message property value. 
     * @returns a string
     */
    public get message() {
        return this._message;
    };
    /**
     * Gets the target property value. 
     * @returns a string
     */
    public get target() {
        return this._target;
    };
    /**
     * The deserialization information for the current model
     * @returns a Map<string, (item: T, node: ParseNode) => void>
     */
    public getFieldDeserializers<T>() : Map<string, (item: T, node: ParseNode) => void> {
        return new Map<string, (item: T, node: ParseNode) => void>([
            ["code", (o, n) => { (o as unknown as Main).code = n.getStringValue(); }],
            ["details", (o, n) => { (o as unknown as Main).details = n.getCollectionOfObjectValues<Detail>(Detail); }],
            ["innererror", (o, n) => { (o as unknown as Main).innererror = n.getObjectValue<Innererror>(Innererror); }],
            ["message", (o, n) => { (o as unknown as Main).message = n.getStringValue(); }],
            ["target", (o, n) => { (o as unknown as Main).target = n.getStringValue(); }],
        ]);
    };
    /**
     * Serializes information the current object
     * @param writer Serialization writer to use to serialize this model
     */
    public serialize(writer: SerializationWriter) : void {
        if(!writer) throw new Error("writer cannot be undefined");
        writer.writeStringValue("code", this.code);
        writer.writeCollectionOfObjectValues<Detail>("details", this.details);
        writer.writeObjectValue<Innererror>("innererror", this.innererror);
        writer.writeStringValue("message", this.message);
        writer.writeStringValue("target", this.target);
        writer.writeAdditionalData(this.additionalData);
    };
    /**
     * Sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
     * @param value Value to set for the AdditionalData property.
     */
    public set additionalData(value: Map<string, unknown>) {
        this._additionalData = value;
    };
    /**
     * Sets the code property value. 
     * @param value Value to set for the code property.
     */
    public set code(value: string | undefined) {
        this._code = value;
    };
    /**
     * Sets the details property value. 
     * @param value Value to set for the details property.
     */
    public set details(value: Detail[] | undefined) {
        this._details = value;
    };
    /**
     * Sets the innererror property value. The structure of this object is service-specific
     * @param value Value to set for the innererror property.
     */
    public set innererror(value: Innererror | undefined) {
        this._innererror = value;
    };
    /**
     * Sets the message property value. 
     * @param value Value to set for the message property.
     */
    public set message(value: string | undefined) {
        this._message = value;
    };
    /**
     * Sets the target property value. 
     * @param value Value to set for the target property.
     */
    public set target(value: string | undefined) {
        this._target = value;
    };
}
