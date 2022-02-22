package graphjavav4.utilities.users.item.mailFolders.item.multiValueExtendedProperties;

import com.microsoft.kiota.serialization.Parsable;
import com.microsoft.kiota.serialization.ParseNode;
import com.microsoft.kiota.serialization.SerializationWriter;
import graphjavav4.utilities.models.microsoft.graph.MultiValueLegacyExtendedProperty;
import java.util.function.BiConsumer;
import java.util.HashMap;
import java.util.Map;
import java.util.Objects;
public class MultiValueExtendedPropertiesResponse implements Parsable {
    /** Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.  */
    private Map<String, Object> _additionalData;
    private String _nextLink;
    private java.util.List<MultiValueLegacyExtendedProperty> _value;
    /**
     * Instantiates a new multiValueExtendedPropertiesResponse and sets the default values.
     * @return a void
     */
    public MultiValueExtendedPropertiesResponse() {
        this.setAdditionalData(new HashMap<>());
    }
    @javax.annotation.Nonnull
    public static MultiValueExtendedPropertiesResponse createFromDiscriminatorValue(@javax.annotation.Nonnull final ParseNode parseNode) {
        Objects.requireNonNull(parseNode);
        return new MultiValueExtendedPropertiesResponse();
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
     * Gets the @odata.nextLink property value. 
     * @return a string
     */
    @javax.annotation.Nullable
    public String getNextLink() {
        return this._nextLink;
    }
    /**
     * Gets the value property value. 
     * @return a multiValueLegacyExtendedProperty
     */
    @javax.annotation.Nullable
    public java.util.List<MultiValueLegacyExtendedProperty> getValue() {
        return this._value;
    }
    /**
     * The deserialization information for the current model
     * @return a Map<String, BiConsumer<T, ParseNode>>
     */
    @javax.annotation.Nonnull
    public <T> Map<String, BiConsumer<T, ParseNode>> getFieldDeserializers() {
        return new HashMap<>(2) {{
            this.put("@odata.nextLink", (o, n) -> { ((MultiValueExtendedPropertiesResponse)o).setNextLink(n.getStringValue()); });
            this.put("value", (o, n) -> { ((MultiValueExtendedPropertiesResponse)o).setValue(n.getCollectionOfObjectValues(MultiValueLegacyExtendedProperty::createFromDiscriminatorValue)); });
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
     * Sets the @odata.nextLink property value. 
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
    public void setValue(@javax.annotation.Nullable final java.util.List<MultiValueLegacyExtendedProperty> value) {
        this._value = value;
    }
}
