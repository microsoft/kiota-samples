package graphjavav4.utilities.users;

import com.microsoft.kiota.serialization.Parsable;
import com.microsoft.kiota.serialization.ParseNode;
import com.microsoft.kiota.serialization.SerializationWriter;
import java.util.function.BiConsumer;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.Objects;
public class OutlookItem extends Entity implements Parsable {
    /** The categories associated with the item  */
    @javax.annotation.Nullable
    public List<String> categories;
    /** Identifies the version of the item. Every time the item is changed, changeKey changes as well. This allows Exchange to apply changes to the correct version of the object. Read-only.  */
    @javax.annotation.Nullable
    public String changeKey;
    /** The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z  */
    @javax.annotation.Nullable
    public String createdDateTime;
    /** The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z  */
    @javax.annotation.Nullable
    public String lastModifiedDateTime;
    /**
     * Serialiazes information the current object
     * @param writer Serialization writer to use to serialize this model
     * @return a void
     */
    public void serialize(@javax.annotation.Nonnull final SerializationWriter writer) {
        Objects.requireNonNull(writer);
        super.serialize(writer);
        writer.writeCollectionOfPrimitiveValues("categories", categories);
        writer.writeStringValue("changeKey", changeKey);
        writer.writeStringValue("createdDateTime", createdDateTime);
        writer.writeStringValue("lastModifiedDateTime", lastModifiedDateTime);
    }
    /**
     * The serialization information for the current model
     * @return a Map<String, BiConsumer<T, ParseNode>>
     */
    @javax.annotation.Nonnull
    public <T> Map<String, BiConsumer<T, ParseNode>> getDeserializeFields() {
        final Map<String, BiConsumer<T, ParseNode>> fields = new HashMap<>(super.getDeserializeFields());
        fields.put("categories", (o, n) -> { ((OutlookItem)o).categories = n.getCollectionOfPrimitiveValues(String.class); });
        fields.put("changeKey", (o, n) -> { ((OutlookItem)o).changeKey = n.getStringValue(); });
        fields.put("createdDateTime", (o, n) -> { ((OutlookItem)o).createdDateTime = n.getStringValue(); });
        fields.put("lastModifiedDateTime", (o, n) -> { ((OutlookItem)o).lastModifiedDateTime = n.getStringValue(); });
        return fields;
    }
}
