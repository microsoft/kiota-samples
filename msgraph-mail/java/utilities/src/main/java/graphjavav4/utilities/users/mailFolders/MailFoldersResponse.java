package graphjavav4.utilities.users.mailFolders;

import com.microsoft.kiota.serialization.Parsable;
import com.microsoft.kiota.serialization.ParseNode;
import com.microsoft.kiota.serialization.SerializationWriter;
import graphjavav4.utilities.users.MailFolder;
import java.util.function.BiConsumer;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.Objects;
public class MailFoldersResponse implements Parsable {
    /** Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.  */
    private IDictionary<string, object> _additionalData;
    private String _nextLink;
    private List<MailFolder> _value;
    /**
     * Instantiates a new MailFoldersResponse and sets the default values.
     * @return a void
     */
    public MailFoldersResponse() {
        this._additionalData = new Dictionary<string, object>();
    }
    /**
     * Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
     * @return a IDictionary<string, object>
     */
    @javax.annotation.Nonnull
    public IDictionary<string, object> getAdditionalData() {
        return this._additionalData;
    }
    /**
     * Gets the nextLink property value. 
     * @return a string
     */
    @javax.annotation.Nullable
    public String getNextLink() {
        return this._nextLink;
    }
    /**
     * Gets the value property value. 
     * @return a mailFolder
     */
    @javax.annotation.Nullable
    public List<MailFolder> getValue() {
        return this._value;
    }
    /**
     * The deserialization information for the current model
     * @return a Map<String, BiConsumer<T, ParseNode>>
     */
    @javax.annotation.Nonnull
    public <T> Map<String, BiConsumer<T, ParseNode>> getFieldDeserializers() {
        return new HashMap<>(2) {{
            this.put("@odata.nextLink", (o, n) -> { ((MailFoldersResponse)o).setNextLink(n.getStringValue()); });
            this.put("value", (o, n) -> { ((MailFoldersResponse)o).setValue(n.getCollectionOfObjectValues(MailFolder.class)); });
        }};
    }
    /**
     * Serializes information the current object
     * @param writer Serialization writer to use to serialize this model
     * @return a void
     */
    public void serialize(@javax.annotation.Nonnull final SerializationWriter writer) {
        Objects.requireNonNull(writer);
        writer.writeStringValue("@odata.nextLink", this.getNextLink());
        writer.writeCollectionOfObjectValues("value", this.getValue());
        writer.writeAdditionalData(this.additionalData);
    }
    /**
     * Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
     * @param value Value to set for the AdditionalData property.
     * @return a void
     */
    public void setAdditionalData(@javax.annotation.Nullable final IDictionary<string, object> value) {
        this._additionalData = value;
    }
    /**
     * Sets the nextLink property value. 
     * @param value Value to set for the nextLink property.
     * @return a void
     */
    public void setNextLink(@javax.annotation.Nullable final String value) {
        this._nextLink = value;
    }
    /**
     * Sets the value property value. 
     * @param value Value to set for the value property.
     * @return a void
     */
    public void setValue(@javax.annotation.Nullable final List<MailFolder> value) {
        this._value = value;
    }
}
