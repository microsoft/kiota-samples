package graphjavav4.utilities.users;

import com.microsoft.kiota.serialization.Parsable;
import com.microsoft.kiota.serialization.ParseNode;
import com.microsoft.kiota.serialization.SerializationWriter;
import java.time.OffsetDateTime;
import java.util.function.BiConsumer;
import java.util.HashMap;
import java.util.Map;
import java.util.Objects;
public class Attachment extends Entity implements Parsable {
    /** The MIME type.  */
    @javax.annotation.Nullable
    public String contentType;
    /** true if the attachment is an inline attachment; otherwise, false.  */
    @javax.annotation.Nullable
    public Boolean isInline;
    /** The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z  */
    @javax.annotation.Nullable
    public OffsetDateTime lastModifiedDateTime;
    /** The display name of the attachment. This does not need to be the actual file name.  */
    @javax.annotation.Nullable
    public String name;
    /** The length of the attachment in bytes.  */
    @javax.annotation.Nullable
    public Integer size;
    /**
     * The serialization information for the current model
     * @return a Map<String, BiConsumer<T, ParseNode>>
     */
    @javax.annotation.Nonnull
    public <T> Map<String, BiConsumer<T, ParseNode>> getDeserializeFields() {
        final Map<String, BiConsumer<T, ParseNode>> fields = new HashMap<>(super.getDeserializeFields());
        fields.put("contentType", (o, n) -> { ((Attachment)o).contentType = n.getStringValue(); });
        fields.put("isInline", (o, n) -> { ((Attachment)o).isInline = n.getBooleanValue(); });
        fields.put("lastModifiedDateTime", (o, n) -> { ((Attachment)o).lastModifiedDateTime = n.getOffsetDateTimeValue(); });
        fields.put("name", (o, n) -> { ((Attachment)o).name = n.getStringValue(); });
        fields.put("size", (o, n) -> { ((Attachment)o).size = n.getIntegerValue(); });
        return fields;
    }
    /**
     * Serialiazes information the current object
     * @param writer Serialization writer to use to serialize this model
     * @return a void
     */
    public void serialize(@javax.annotation.Nonnull final SerializationWriter writer) {
        Objects.requireNonNull(writer);
        super.serialize(writer);
        writer.writeStringValue("contentType", contentType);
        writer.writeBooleanValue("isInline", isInline);
        writer.writeOffsetDateTimeValue("lastModifiedDateTime", lastModifiedDateTime);
        writer.writeStringValue("name", name);
        writer.writeIntegerValue("size", size);
    }
}
