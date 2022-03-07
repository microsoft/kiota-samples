package todoclient.utilities.todoitems;

import com.microsoft.kiota.serialization.Parsable;
import com.microsoft.kiota.serialization.ParseNode;
import com.microsoft.kiota.serialization.SerializationWriter;
import java.util.function.BiConsumer;
import java.util.HashMap;
import java.util.Map;
import java.util.Objects;
import todoclient.utilities.models.todoapi.models.ToDoItem;
public class ToDoItemsResponse implements Parsable {
    /** Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.  */
    private Map<String, Object> _additionalData;
    private java.util.List<ToDoItem> _value;
    /**
     * Instantiates a new ToDoItemsResponse and sets the default values.
     * @return a void
     */
    public ToDoItemsResponse() {
        this.setAdditionalData(new HashMap<>());
    }
    /**
     * Gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
     * @return a Map<String, Object>
     */
    @javax.annotation.Nonnull
    public Map<String, Object> getAdditionalData() {
        return this._additionalData;
    }
    /**
     * Gets the value property value. 
     * @return a ToDoItem
     */
    @javax.annotation.Nullable
    public java.util.List<ToDoItem> getValue() {
        return this._value;
    }
    /**
     * The deserialization information for the current model
     * @return a Map<String, BiConsumer<T, ParseNode>>
     */
    @javax.annotation.Nonnull
    public <T> Map<String, BiConsumer<T, ParseNode>> getFieldDeserializers() {
        return new HashMap<>(1) {{
            this.put("value", (o, n) -> { ((ToDoItemsResponse)o).setValue(n.getCollectionOfObjectValues(ToDoItem.class)); });
        }};
    }
    /**
     * Serializes information the current object
     * @param writer Serialization writer to use to serialize this model
     * @return a void
     */
    public void serialize(@javax.annotation.Nonnull final SerializationWriter writer) {
        Objects.requireNonNull(writer);
        writer.writeCollectionOfObjectValues("value", this.getValue());
        writer.writeAdditionalData(this.getAdditionalData());
    }
    /**
     * Sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
     * @param value Value to set for the AdditionalData property.
     * @return a void
     */
    public void setAdditionalData(@javax.annotation.Nullable final Map<String, Object> value) {
        this._additionalData = value;
    }
    /**
     * Sets the value property value. 
     * @param value Value to set for the value property.
     * @return a void
     */
    public void setValue(@javax.annotation.Nullable final java.util.List<ToDoItem> value) {
        this._value = value;
    }
}
