package graphjavav4.utilities.users;

import com.microsoft.kiota.serialization.Parsable;
import com.microsoft.kiota.serialization.ParseNode;
import com.microsoft.kiota.serialization.SerializationWriter;
import java.util.function.BiConsumer;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.Objects;
public class MultiValueLegacyExtendedProperty extends Entity implements Parsable {
    /** A collection of property values.  */
    private List<String> _value;
    /**
     * Instantiates a new multiValueLegacyExtendedProperty and sets the default values.
     * @return a void
     */
    public MultiValueLegacyExtendedProperty() {
        super();
    }
    /**
     * Gets the value property value. A collection of property values.
     * @return a string
     */
    @javax.annotation.Nullable
    public List<String> getValue() {
        return this._value;
    }
    /**
     * The deserialization information for the current model
     * @return a Map<String, BiConsumer<T, ParseNode>>
     */
    @javax.annotation.Nonnull
    public <T> Map<String, BiConsumer<T, ParseNode>> getFieldDeserializers() {
        return new HashMap<>(super.getFieldDeserializers()) {{
            this.put("value", (o, n) -> { ((MultiValueLegacyExtendedProperty)o).setValue(n.getCollectionOfPrimitiveValues(String.class)); });
        }};
    }
    /**
     * Serializes information the current object
     * @param writer Serialization writer to use to serialize this model
     * @return a void
     */
    public void serialize(@javax.annotation.Nonnull final SerializationWriter writer) {
        Objects.requireNonNull(writer);
        super.serialize(writer);
        writer.writeCollectionOfPrimitiveValues("value", this.getValue());
    }
    /**
     * Sets the value property value. A collection of property values.
     * @param value Value to set for the value property.
     * @return a void
     */
    public void setValue(@javax.annotation.Nullable final List<String> value) {
        this._value = value;
    }
}
