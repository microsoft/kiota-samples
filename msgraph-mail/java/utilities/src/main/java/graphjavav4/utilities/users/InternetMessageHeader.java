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
    @javax.annotation.Nonnull
    private Map<String, Object> additionalData = new HashMap<>();
    /** Represents the key in a key-value pair.  */
    @javax.annotation.Nullable
    public String name;
    /** The value in a key-value pair.  */
    @javax.annotation.Nullable
    public String value;
    /**
     * Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
     * @return a Map<String, Object>
     */
    @javax.annotation.Nonnull
    public Map<String, Object> getAdditionalData() {
        return additionalData;
    }
    /**
     * The deserialization information for the current model
     * @return a Map<String, BiConsumer<T, ParseNode>>
     */
    @javax.annotation.Nonnull
    public <T> Map<String, BiConsumer<T, ParseNode>> getFieldDeserializers() {
        return new HashMap<>(2) {{
            this.put("name", (o, n) -> { ((InternetMessageHeader)o).name = n.getStringValue(); });
            this.put("value", (o, n) -> { ((InternetMessageHeader)o).value = n.getStringValue(); });
        }};
    }
    /**
     * Serializes information the current object
     * @param writer Serialization writer to use to serialize this model
     * @return a void
     */
    public void serialize(@javax.annotation.Nonnull final SerializationWriter writer) {
        Objects.requireNonNull(writer);
        writer.writeStringValue("name", name);
        writer.writeStringValue("value", value);
        writer.writeAdditionalData(this.additionalData);
    }
}
