package graphjavav4.utilities.users;

import java.util.Objects;
import com.microsoft.kiota.serialization.SerializationWriter;
import com.microsoft.kiota.serialization.ParseNode;
import com.microsoft.kiota.serialization.Parsable;
import java.util.function.BiConsumer;
import java.util.Map;
import java.util.HashMap;
import java.util.List;
public class OutlookItem extends Entity implements Parsable {
    @javax.annotation.Nullable
    public List<Object> categories;
    @javax.annotation.Nullable
    public String changeKey;
    @javax.annotation.Nullable
    public String createdDateTime;
    @javax.annotation.Nullable
    public String lastModifiedDateTime;
    public void serialize(@javax.annotation.Nonnull final SerializationWriter writer) {
        Objects.requireNonNull(writer);
        super.serialize(writer);
        writer.writeCollectionOfPrimitiveValues("categories", categories);
        writer.writeStringValue("changeKey", changeKey);
        writer.writeStringValue("createdDateTime", createdDateTime);
        writer.writeStringValue("lastModifiedDateTime", lastModifiedDateTime);
    }
    @javax.annotation.Nonnull
    public <T> Map<String, BiConsumer<T, ParseNode>> getDeserializeFields() {
        final Map<String, BiConsumer<T, ParseNode>> fields = new HashMap<>(super.getDeserializeFields());
        fields.put("categories", (o, n) -> { ((OutlookItem)o).categories = n.getCollectionOfPrimitiveValues(Object.class); });
        fields.put("changeKey", (o, n) -> { ((OutlookItem)o).changeKey = n.getStringValue(); });
        fields.put("createdDateTime", (o, n) -> { ((OutlookItem)o).createdDateTime = n.getStringValue(); });
        fields.put("lastModifiedDateTime", (o, n) -> { ((OutlookItem)o).lastModifiedDateTime = n.getStringValue(); });
        return fields;
    }
}
