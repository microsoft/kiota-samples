package graphjavav4.utilities.users.inferenceClassification;

import com.microsoft.kiota.serialization.Parsable;
import com.microsoft.kiota.serialization.ParseNode;
import com.microsoft.kiota.serialization.SerializationWriter;
import graphjavav4.utilities.users.Entity;
import java.util.function.BiConsumer;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.Objects;
public class InferenceClassification extends Entity implements Parsable {
    /** A set of overrides for a user to always classify messages from specific senders in certain ways: focused, or other. Read-only. Nullable.  */
    private List<InferenceClassificationOverride> _overrides;
    /**
     * Instantiates a new InferenceClassification and sets the default values.
     * @return a void
     */
    public InferenceClassification() {
        super();
    }
    /**
     * Gets the overrides property value. A set of overrides for a user to always classify messages from specific senders in certain ways: focused, or other. Read-only. Nullable.
     * @return a inferenceClassificationOverride
     */
    @javax.annotation.Nullable
    public List<InferenceClassificationOverride> getOverrides() {
        return this._overrides;
    }
    /**
     * The deserialization information for the current model
     * @return a Map<String, BiConsumer<T, ParseNode>>
     */
    @javax.annotation.Nonnull
    public <T> Map<String, BiConsumer<T, ParseNode>> getFieldDeserializers() {
        return new HashMap<>(super.getFieldDeserializers()) {{
            this.put("overrides", (o, n) -> { ((InferenceClassification)o).setOverrides(n.getCollectionOfObjectValues(InferenceClassificationOverride.class)); });
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
        writer.writeCollectionOfObjectValues("overrides", this.getOverrides());
    }
    /**
     * Sets the overrides property value. A set of overrides for a user to always classify messages from specific senders in certain ways: focused, or other. Read-only. Nullable.
     * @param value Value to set for the overrides property.
     * @return a void
     */
    public void setOverrides(@javax.annotation.Nullable final List<InferenceClassificationOverride> value) {
        this._overrides = value;
    }
}
