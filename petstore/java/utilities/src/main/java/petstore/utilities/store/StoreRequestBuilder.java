package petstore.utilities.store;

import com.microsoft.kiota.BaseRequestBuilder;
import com.microsoft.kiota.RequestAdapter;
import java.util.HashMap;
import java.util.Objects;
import petstore.utilities.store.inventory.InventoryRequestBuilder;
import petstore.utilities.store.order.OrderRequestBuilder;
/**
 * Builds and executes requests for operations under /store
 */
@jakarta.annotation.Generated("com.microsoft.kiota")
public class StoreRequestBuilder extends BaseRequestBuilder {
    /**
     * The inventory property
     * @return a {@link InventoryRequestBuilder}
     */
    @jakarta.annotation.Nonnull
    public InventoryRequestBuilder inventory() {
        return new InventoryRequestBuilder(pathParameters, requestAdapter);
    }
    /**
     * The order property
     * @return a {@link OrderRequestBuilder}
     */
    @jakarta.annotation.Nonnull
    public OrderRequestBuilder order() {
        return new OrderRequestBuilder(pathParameters, requestAdapter);
    }
    /**
     * Instantiates a new {@link StoreRequestBuilder} and sets the default values.
     * @param pathParameters Path parameters for the request
     * @param requestAdapter The request adapter to use to execute the requests.
     */
    public StoreRequestBuilder(@jakarta.annotation.Nonnull final HashMap<String, Object> pathParameters, @jakarta.annotation.Nonnull final RequestAdapter requestAdapter) {
        super(requestAdapter, "{+baseurl}/store", pathParameters);
    }
    /**
     * Instantiates a new {@link StoreRequestBuilder} and sets the default values.
     * @param rawUrl The raw URL to use for the request builder.
     * @param requestAdapter The request adapter to use to execute the requests.
     */
    public StoreRequestBuilder(@jakarta.annotation.Nonnull final String rawUrl, @jakarta.annotation.Nonnull final RequestAdapter requestAdapter) {
        super(requestAdapter, "{+baseurl}/store", rawUrl);
    }
}
