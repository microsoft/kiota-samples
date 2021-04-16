package graphjavav4.utilities.users;

import java.util.Objects;
import graphjavav4.utilities.users.mailFolders.MailFolder;
import com.microsoft.kiota.serialization.SerializationWriter;
import com.microsoft.kiota.serialization.ParseNode;
import com.microsoft.kiota.serialization.Parsable;
import java.util.function.BiConsumer;
import java.util.Map;
import java.util.HashMap;
import java.util.List;
public class MailFoldersResponse implements Parsable {
    @javax.annotation.Nullable
    public List<MailFolder> value;
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
        fields.put("value", (o, n) -> { ((MailFoldersResponse)o).value = n.getCollectionOfObjectValues(MailFolder.class); });
        fields.put("nextLink", (o, n) -> { ((MailFoldersResponse)o).nextLink = n.getStringValue(); });
        return fields;
    }
}
