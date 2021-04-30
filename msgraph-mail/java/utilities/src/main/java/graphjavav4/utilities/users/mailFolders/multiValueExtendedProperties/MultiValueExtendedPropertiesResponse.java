package graphjavav4.utilities.users.mailFolders.multiValueExtendedProperties;

import com.microsoft.kiota.serialization.Parsable;
import com.microsoft.kiota.serialization.ParseNode;
import com.microsoft.kiota.serialization.SerializationWriter;
import graphjavav4.utilities.users.MultiValueLegacyExtendedProperty;
import java.util.function.BiConsumer;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.Objects;
public class MultiValueExtendedPropertiesResponse implements Parsable {
    /** Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.  */
    @javax.annotation.Nonnull
    private final Map<String, Object> additionalData = new HashMap<>();
    @javax.annotation.Nullable
    public String nextLink;
    @javax.annotation.Nullable
    public List<MultiValueLegacyExtendedProperty> value;
    /**
     * Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
     * @return a Map<String, Object>
     */
    @javax.annotation.Nonnull
    public Map<String, Object> getAdditionalData() {
        return additionalData;
    }
    /**
     * The serialization information for the current model
     * @return a Map<String, BiConsumer<T, ParseNode>>
     */
    @javax.annotation.Nonnull
    public <T> Map<String, BiConsumer<T, ParseNode>> getDeserializeFields() {
        final Map<String, BiConsumer<T, ParseNode>> fields = new HashMap<>(2);
        fields.put("@odata.nextLink", (o, n) -> { ((MultiValueExtendedPropertiesResponse)o).nextLink = n.getStringValue(); });
        fields.put("value", (o, n) -> { ((MultiValueExtendedPropertiesResponse)o).value = n.getCollectionOfObjectValues(MultiValueLegacyExtendedProperty.class); });
        return fields;
    }
    /**
     * Serialiazes information the current object
     * @param writer Serialization writer to use to serialize this model
     * @return a void
     */
    public void serialize(@javax.annotation.Nonnull final SerializationWriter writer) {
        Objects.requireNonNull(writer);
        writer.writeStringValue("@odata.nextLink", nextLink);
        writer.writeCollectionOfObjectValues("value", value);
        writer.writeAdditionalData(this.additionalData);
    }
}
