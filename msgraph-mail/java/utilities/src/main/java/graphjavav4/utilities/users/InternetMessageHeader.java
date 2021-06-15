package graphjavav4.utilities.users;

import com.microsoft.kiota.serialization.Parsable;
import com.microsoft.kiota.serialization.ParseNode;
import com.microsoft.kiota.serialization.SerializationWriter;
import java.util.function.BiConsumer;
import java.util.HashMap;
import java.util.Map;
import java.util.Objects;
public class InternetMessageHeader implements Parsable {
    /** Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.  */
    private IDictionary<string, object> _additionalData;
    /** Represents the key in a key-value pair.  */
    private String _name;
    /** The value in a key-value pair.  */
    private String _value;
    /**
     * Instantiates a new internetMessageHeader and sets the default values.
     * @return a void
     */
    public InternetMessageHeader() {
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
     * Gets the name property value. Represents the key in a key-value pair.
     * @return a string
     */
    @javax.annotation.Nullable
    public String getName() {
        return this._name;
    }
    /**
     * Gets the value property value. The value in a key-value pair.
     * @return a string
     */
    @javax.annotation.Nullable
    public String getValue() {
        return this._value;
    }
    /**
     * The deserialization information for the current model
     * @return a Map<String, BiConsumer<T, ParseNode>>
     */
    @javax.annotation.Nonnull
    public <T> Map<String, BiConsumer<T, ParseNode>> getFieldDeserializers() {
        return new HashMap<>(2) {{
            this.put("name", (o, n) -> { ((InternetMessageHeader)o).setName(n.getStringValue()); });
            this.put("value", (o, n) -> { ((InternetMessageHeader)o).setValue(n.getStringValue()); });
        }};
    }
    /**
     * Serializes information the current object
     * @param writer Serialization writer to use to serialize this model
     * @return a void
     */
    public void serialize(@javax.annotation.Nonnull final SerializationWriter writer) {
        Objects.requireNonNull(writer);
        writer.writeStringValue("name", this.getName());
        writer.writeStringValue("value", this.getValue());
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
     * Sets the name property value. Represents the key in a key-value pair.
     * @param value Value to set for the name property.
     * @return a void
     */
    public void setName(@javax.annotation.Nullable final String value) {
        this._name = value;
    }
    /**
     * Sets the value property value. The value in a key-value pair.
     * @param value Value to set for the value property.
     * @return a void
     */
    public void setValue(@javax.annotation.Nullable final String value) {
        this._value = value;
    }
}
