package graphjavav4.utilities.users;

import com.microsoft.kiota.serialization.Parsable;
import com.microsoft.kiota.serialization.ParseNode;
import com.microsoft.kiota.serialization.SerializationWriter;
import java.util.function.BiConsumer;
import java.util.HashMap;
import java.util.Map;
import java.util.Objects;
public class Entity implements Parsable {
    /** Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.  */
    private IDictionary<string, object> _additionalData;
    /** Read-only.  */
    private String _id;
    /**
     * Instantiates a new entity and sets the default values.
     * @return a void
     */
    public Entity() {
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
     * Gets the id property value. Read-only.
     * @return a string
     */
    @javax.annotation.Nullable
    public String getId() {
        return this._id;
    }
    /**
     * The deserialization information for the current model
     * @return a Map<String, BiConsumer<T, ParseNode>>
     */
    @javax.annotation.Nonnull
    public <T> Map<String, BiConsumer<T, ParseNode>> getFieldDeserializers() {
        return new HashMap<>(1) {{
            this.put("id", (o, n) -> { ((Entity)o).setId(n.getStringValue()); });
        }};
    }
    /**
     * Serializes information the current object
     * @param writer Serialization writer to use to serialize this model
     * @return a void
     */
    public void serialize(@javax.annotation.Nonnull final SerializationWriter writer) {
        Objects.requireNonNull(writer);
        writer.writeStringValue("id", this.getId());
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
     * Sets the id property value. Read-only.
     * @param value Value to set for the id property.
     * @return a void
     */
    public void setId(@javax.annotation.Nullable final String value) {
        this._id = value;
    }
}
