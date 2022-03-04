package graphjavav4.utilities.users.item.messages.item.singlevalueextendedproperties;

import com.microsoft.kiota.serialization.AdditionalDataHolder;
import com.microsoft.kiota.serialization.Parsable;
import com.microsoft.kiota.serialization.ParseNode;
import com.microsoft.kiota.serialization.SerializationWriter;
import graphjavav4.utilities.models.microsoft.graph.SingleValueLegacyExtendedProperty;
import java.util.function.BiConsumer;
import java.util.HashMap;
import java.util.Map;
import java.util.Objects;
public class SingleValueExtendedPropertiesResponse implements AdditionalDataHolder, Parsable {
    /** Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.  */
    private Map<String, Object> _additionalData;
    private String _nextLink;
    private java.util.List<SingleValueLegacyExtendedProperty> _value;
    /**
     * Instantiates a new singleValueExtendedPropertiesResponse and sets the default values.
     * @return a void
     */
    public SingleValueExtendedPropertiesResponse() {
        this.setAdditionalData(new HashMap<>());
    }
    /**
     * Creates a new instance of the appropriate class based on discriminator value
     * @param parseNode The parse node to use to read the discriminator value and create the object
     * @return a singleValueExtendedPropertiesResponse
     */
    @javax.annotation.Nonnull
    public static SingleValueExtendedPropertiesResponse createFromDiscriminatorValue(@javax.annotation.Nonnull final ParseNode parseNode) {
        Objects.requireNonNull(parseNode);
        return new SingleValueExtendedPropertiesResponse();
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
     * The deserialization information for the current model
     * @return a Map<String, BiConsumer<T, ParseNode>>
     */
    @javax.annotation.Nonnull
    public <T> Map<String, BiConsumer<T, ParseNode>> getFieldDeserializers() {
        return new HashMap<>(2) {{
            this.put("@odata.nextLink", (o, n) -> { ((SingleValueExtendedPropertiesResponse)o).setNextLink(n.getStringValue()); });
            this.put("value", (o, n) -> { ((SingleValueExtendedPropertiesResponse)o).setValue(n.getCollectionOfObjectValues(SingleValueLegacyExtendedProperty::createFromDiscriminatorValue)); });
        }};
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
     * @return a singleValueLegacyExtendedProperty
     */
    @javax.annotation.Nullable
    public java.util.List<SingleValueLegacyExtendedProperty> getValue() {
        return this._value;
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
    public void setValue(@javax.annotation.Nullable final java.util.List<SingleValueLegacyExtendedProperty> value) {
        this._value = value;
    }
}
