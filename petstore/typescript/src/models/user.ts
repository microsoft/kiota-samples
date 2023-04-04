import {AdditionalDataHolder, Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export class User implements AdditionalDataHolder, Parsable {
    /** Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well. */
    private _additionalData: Record<string, unknown>;
    /** The email property */
    private _email?: string | undefined;
    /** The firstName property */
    private _firstName?: string | undefined;
    /** The id property */
    private _id?: number | undefined;
    /** The lastName property */
    private _lastName?: string | undefined;
    /** The password property */
    private _password?: string | undefined;
    /** The phone property */
    private _phone?: string | undefined;
    /** The username property */
    private _username?: string | undefined;
    /** User Status */
    private _userStatus?: number | undefined;
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
     * Instantiates a new User and sets the default values.
     */
    public constructor() {
        this._additionalData = {};
    };
    /**
     * Gets the email property value. The email property
     * @returns a string
     */
    public get email() {
        return this._email;
    };
    /**
     * Sets the email property value. The email property
     * @param value Value to set for the email property.
     */
    public set email(value: string | undefined) {
        this._email = value;
    };
    /**
     * Gets the firstName property value. The firstName property
     * @returns a string
     */
    public get firstName() {
        return this._firstName;
    };
    /**
     * Sets the firstName property value. The firstName property
     * @param value Value to set for the firstName property.
     */
    public set firstName(value: string | undefined) {
        this._firstName = value;
    };
    /**
     * The deserialization information for the current model
     * @returns a Record<string, (node: ParseNode) => void>
     */
    public getFieldDeserializers() : Record<string, (node: ParseNode) => void> {
        return {
            "email": n => { this.email = n.getStringValue(); },
            "firstName": n => { this.firstName = n.getStringValue(); },
            "id": n => { this.id = n.getNumberValue(); },
            "lastName": n => { this.lastName = n.getStringValue(); },
            "password": n => { this.password = n.getStringValue(); },
            "phone": n => { this.phone = n.getStringValue(); },
            "username": n => { this.username = n.getStringValue(); },
            "userStatus": n => { this.userStatus = n.getNumberValue(); },
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
     * Gets the lastName property value. The lastName property
     * @returns a string
     */
    public get lastName() {
        return this._lastName;
    };
    /**
     * Sets the lastName property value. The lastName property
     * @param value Value to set for the lastName property.
     */
    public set lastName(value: string | undefined) {
        this._lastName = value;
    };
    /**
     * Gets the password property value. The password property
     * @returns a string
     */
    public get password() {
        return this._password;
    };
    /**
     * Sets the password property value. The password property
     * @param value Value to set for the password property.
     */
    public set password(value: string | undefined) {
        this._password = value;
    };
    /**
     * Gets the phone property value. The phone property
     * @returns a string
     */
    public get phone() {
        return this._phone;
    };
    /**
     * Sets the phone property value. The phone property
     * @param value Value to set for the phone property.
     */
    public set phone(value: string | undefined) {
        this._phone = value;
    };
    /**
     * Serializes information the current object
     * @param writer Serialization writer to use to serialize this model
     */
    public serialize(writer: SerializationWriter) : void {
        if(!writer) throw new Error("writer cannot be undefined");
        writer.writeStringValue("email", this.email);
        writer.writeStringValue("firstName", this.firstName);
        writer.writeNumberValue("id", this.id);
        writer.writeStringValue("lastName", this.lastName);
        writer.writeStringValue("password", this.password);
        writer.writeStringValue("phone", this.phone);
        writer.writeStringValue("username", this.username);
        writer.writeNumberValue("userStatus", this.userStatus);
        writer.writeAdditionalData(this.additionalData);
    };
    /**
     * Gets the username property value. The username property
     * @returns a string
     */
    public get username() {
        return this._username;
    };
    /**
     * Sets the username property value. The username property
     * @param value Value to set for the username property.
     */
    public set username(value: string | undefined) {
        this._username = value;
    };
    /**
     * Gets the userStatus property value. User Status
     * @returns a integer
     */
    public get userStatus() {
        return this._userStatus;
    };
    /**
     * Sets the userStatus property value. User Status
     * @param value Value to set for the userStatus property.
     */
    public set userStatus(value: number | undefined) {
        this._userStatus = value;
    };
}
