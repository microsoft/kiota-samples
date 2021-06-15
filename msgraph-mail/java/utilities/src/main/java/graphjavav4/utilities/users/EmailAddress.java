package graphjavav4.utilities.users;

import com.microsoft.kiota.serialization.Parsable;
import com.microsoft.kiota.serialization.ParseNode;
import com.microsoft.kiota.serialization.SerializationWriter;
import java.util.function.BiConsumer;
import java.util.HashMap;
import java.util.Map;
import java.util.Objects;
public class EmailAddress implements Parsable {
    /** Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.  */
    private IDictionary<string, object> _additionalData;
    /** The email address of an entity instance.  */
    private String _address;
    /** The display name of an entity instance.  */
    private String _name;
    /**
     * Instantiates a new emailAddress and sets the default values.
     * @return a void
     */
    public EmailAddress() {
        this._additionalData = new Dictionary<string, object>();
    }
    /**
     * Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
     * @return a IDictionary<string, object>
     */
    @javax.annotation.Nonnull
    public IDictionary<string, object> getAdditionalData() {
        return this._additionalData;
    }
    /**
     * Gets the address property value. The email address of an entity instance.
     * @return a string
     */
    @javax.annotation.Nullable
    public String getAddress() {
        return this._address;
    }
    /**
     * Gets the name property value. The display name of an entity instance.
     * @return a string
     */
    @javax.annotation.Nullable
    public String getName() {
        return this._name;
    }
    /**
     * The deserialization information for the current model
     * @return a Map<String, BiConsumer<T, ParseNode>>
     */
    @javax.annotation.Nonnull
    public <T> Map<String, BiConsumer<T, ParseNode>> getFieldDeserializers() {
        return new HashMap<>(2) {{
            this.put("address", (o, n) -> { ((EmailAddress)o).setAddress(n.getStringValue()); });
            this.put("name", (o, n) -> { ((EmailAddress)o).setName(n.getStringValue()); });
        }};
    }
    /**
     * Serializes information the current object
     * @param writer Serialization writer to use to serialize this model
     * @return a void
     */
    public void serialize(@javax.annotation.Nonnull final SerializationWriter writer) {
        Objects.requireNonNull(writer);
        writer.writeStringValue("address", this.getAddress());
        writer.writeStringValue("name", this.getName());
        writer.writeAdditionalData(this.additionalData);
    }
    /**
     * Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
     * @param value Value to set for the AdditionalData property.
     * @return a void
     */
    public void setAdditionalData(@javax.annotation.Nullable final IDictionary<string, object> value) {
        this._additionalData = value;
    }
    /**
     * Sets the address property value. The email address of an entity instance.
     * @param value Value to set for the address property.
     * @return a void
     */
    public void setAddress(@javax.annotation.Nullable final String value) {
        this._address = value;
    }
    /**
     * Sets the name property value. The display name of an entity instance.
     * @param value Value to set for the name property.
     * @return a void
     */
    public void setName(@javax.annotation.Nullable final String value) {
        this._name = value;
    }
}
