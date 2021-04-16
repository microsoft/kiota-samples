package graphjavav4.utilities.users.mailFolders.messages.attachments;

import com.microsoft.kiota.serialization.Parsable;
import com.microsoft.kiota.serialization.ParseNode;
import com.microsoft.kiota.serialization.SerializationWriter;
import graphjavav4.utilities.users.Attachment;
import java.util.function.BiConsumer;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.Objects;
public class AttachmentsResponse implements Parsable {
    @javax.annotation.Nullable
    public List<Attachment> value;
    @javax.annotation.Nullable
    public String nextLink;
    public void serialize(@javax.annotation.Nonnull final SerializationWriter writer) {
        Objects.requireNonNull(writer);
        writer.writeCollectionOfObjectValues("value", value);
        writer.writeStringValue("nextLink", nextLink);
    }
    @javax.annotation.Nonnull
    public <T> Map<String, BiConsumer<T, ParseNode>> getDeserializeFields() {
        final Map<String, BiConsumer<T, ParseNode>> fields = new HashMap<>(2);
        fields.put("value", (o, n) -> { ((AttachmentsResponse)o).value = n.getCollectionOfObjectValues(Attachment.class); });
        fields.put("nextLink", (o, n) -> { ((AttachmentsResponse)o).nextLink = n.getStringValue(); });
        return fields;
    }
}
