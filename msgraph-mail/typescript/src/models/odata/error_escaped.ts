import {Main} from './error_escaped/main';
import {ApiError, Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export class Error_escaped extends ApiError implements Parsable {
    /** Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.  */
    private _additionalData: Map<string, unknown>;
    private _error_escaped?: Main | undefined;
    /**
     * Instantiates a new Error_escaped and sets the default values.
     */
    public constructor() {
        super();
        Object.setPrototypeOf(this, ApiError.prototype);
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
     * Gets the error property value. 
     * @returns a main
     */
    public get error_escaped() {
        return this._error_escaped;
    };
    /**
     * The deserialization information for the current model
     * @returns a Map<string, (item: T, node: ParseNode) => void>
     */
    public getFieldDeserializers<T>() : Map<string, (item: T, node: ParseNode) => void> {
        return new Map<string, (item: T, node: ParseNode) => void>([
            ["error", (o, n) => { (o as unknown as Error_escaped).error_escaped = n.getObjectValue<Main>(Main); }],
        ]);
    };
    /**
     * Serializes information the current object
     * @param writer Serialization writer to use to serialize this model
     */
    public serialize(writer: SerializationWriter) : void {
        if(!writer) throw new Error("writer cannot be undefined");
        writer.writeObjectValue<Main>("error", this.error_escaped);
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
     * Sets the error property value. 
     * @param value Value to set for the error_escaped property.
     */
    public set error_escaped(value: Main | undefined) {
        this._error_escaped = value;
    };
}
