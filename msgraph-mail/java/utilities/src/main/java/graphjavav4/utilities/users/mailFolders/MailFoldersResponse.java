package graphjavav4.utilities.users.mailFolders;

import com.microsoft.kiota.serialization.Parsable;
import com.microsoft.kiota.serialization.ParseNode;
import com.microsoft.kiota.serialization.SerializationWriter;
import java.util.function.BiConsumer;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.Objects;
public class MailFoldersResponse implements Parsable {
    @javax.annotation.Nullable
    public List<MailFolder> value;
    @javax.annotation.Nullable
    public String nextLink;
    /**
     * Serialiazes information the current object
     * @param writer Serialization writer to use to serialize this model
     * @return a void
     */
    public void serialize(@javax.annotation.Nonnull final SerializationWriter writer) {
        Objects.requireNonNull(writer);
        writer.writeCollectionOfObjectValues("value", value);
        writer.writeStringValue("nextLink", nextLink);
    }
    /** Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.  */
    @javax.annotation.Nonnull
    private final Map<String, Object> additionalData = new HashMap<>();
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
        fields.put("value", (o, n) -> { ((MailFoldersResponse)o).value = n.getCollectionOfObjectValues(MailFolder.class); });
        fields.put("nextLink", (o, n) -> { ((MailFoldersResponse)o).nextLink = n.getStringValue(); });
        return fields;
    }
}
