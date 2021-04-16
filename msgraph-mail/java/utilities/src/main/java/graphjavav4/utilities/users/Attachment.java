package graphjavav4.utilities.users;

import com.microsoft.kiota.serialization.Parsable;
import com.microsoft.kiota.serialization.ParseNode;
import com.microsoft.kiota.serialization.SerializationWriter;
import java.util.function.BiConsumer;
import java.util.HashMap;
import java.util.Map;
import java.util.Objects;
public class Attachment extends Entity implements Parsable {
    @javax.annotation.Nullable
    public String contentType;
    @javax.annotation.Nullable
    public Boolean isInline;
    @javax.annotation.Nullable
    public String lastModifiedDateTime;
    @javax.annotation.Nullable
    public String name;
    @javax.annotation.Nullable
    public Integer size;
    public void serialize(@javax.annotation.Nonnull final SerializationWriter writer) {
        Objects.requireNonNull(writer);
        super.serialize(writer);
        writer.writeStringValue("contentType", contentType);
        writer.writeBooleanValue("isInline", isInline);
        writer.writeStringValue("lastModifiedDateTime", lastModifiedDateTime);
        writer.writeStringValue("name", name);
        writer.writeIntegerValue("size", size);
    }
    @javax.annotation.Nonnull
    public <T> Map<String, BiConsumer<T, ParseNode>> getDeserializeFields() {
        final Map<String, BiConsumer<T, ParseNode>> fields = new HashMap<>(super.getDeserializeFields());
        fields.put("contentType", (o, n) -> { ((Attachment)o).contentType = n.getStringValue(); });
        fields.put("isInline", (o, n) -> { ((Attachment)o).isInline = n.getBooleanValue(); });
        fields.put("lastModifiedDateTime", (o, n) -> { ((Attachment)o).lastModifiedDateTime = n.getStringValue(); });
        fields.put("name", (o, n) -> { ((Attachment)o).name = n.getStringValue(); });
        fields.put("size", (o, n) -> { ((Attachment)o).size = n.getIntegerValue(); });
        return fields;
    }
}
