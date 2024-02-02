package petstore.utilities.models;

import com.microsoft.kiota.serialization.ValuedEnum;
import java.util.Objects;

/**
 * Order Status
 */
@jakarta.annotation.Generated("com.microsoft.kiota")
public enum OrderStatus implements ValuedEnum {
    Placed("placed"),
    Approved("approved"),
    Delivered("delivered");
    public final String value;
    OrderStatus(final String value) {
        this.value = value;
    }
    @jakarta.annotation.Nonnull
    public String getValue() { return this.value; }
    @jakarta.annotation.Nullable
    public static OrderStatus forValue(@jakarta.annotation.Nonnull final String searchValue) {
        Objects.requireNonNull(searchValue);
        switch(searchValue) {
            case "placed": return Placed;
            case "approved": return Approved;
            case "delivered": return Delivered;
            default: return null;
        }
    }
}
