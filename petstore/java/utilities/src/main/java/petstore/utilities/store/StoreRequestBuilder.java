package petstore.utilities.store;

import com.microsoft.kiota.RequestAdapter;
import java.util.HashMap;
import java.util.Objects;
import petstore.utilities.store.inventory.InventoryRequestBuilder;
import petstore.utilities.store.order.item.WithOrderItemRequestBuilder;
import petstore.utilities.store.order.OrderRequestBuilder;
/**
 * Builds and executes requests for operations under /store
 */
public class StoreRequestBuilder {
    /** The inventory property */
    @javax.annotation.Nonnull
    public InventoryRequestBuilder inventory() {
        return new InventoryRequestBuilder(pathParameters, requestAdapter);
    }
    /** The order property */
    @javax.annotation.Nonnull
    public OrderRequestBuilder order() {
        return new OrderRequestBuilder(pathParameters, requestAdapter);
    }
    /** Path parameters for the request */
    private HashMap<String, Object> pathParameters;
    /** The request adapter to use to execute the requests. */
    private RequestAdapter requestAdapter;
    /** Url template to use to build the URL for the current request builder */
    private String urlTemplate;
    /**
     * Instantiates a new StoreRequestBuilder and sets the default values.
     * @param pathParameters Path parameters for the request
     * @param requestAdapter The request adapter to use to execute the requests.
     * @return a void
     */
    @javax.annotation.Nullable
    public StoreRequestBuilder(@javax.annotation.Nonnull final HashMap<String, Object> pathParameters, @javax.annotation.Nonnull final RequestAdapter requestAdapter) {
        Objects.requireNonNull(pathParameters);
        Objects.requireNonNull(requestAdapter);
        this.urlTemplate = "{+baseurl}/store";
        final HashMap<String, Object> urlTplParams = new HashMap<String, Object>(pathParameters);
        this.pathParameters = urlTplParams;
        this.requestAdapter = requestAdapter;
    }
    /**
     * Instantiates a new StoreRequestBuilder and sets the default values.
     * @param rawUrl The raw URL to use for the request builder.
     * @param requestAdapter The request adapter to use to execute the requests.
     * @return a void
     */
    @javax.annotation.Nullable
    public StoreRequestBuilder(@javax.annotation.Nonnull final String rawUrl, @javax.annotation.Nonnull final RequestAdapter requestAdapter) {
        this.urlTemplate = "{+baseurl}/store";
        final HashMap<String, Object> urlTplParams = new HashMap<String, Object>();
        urlTplParams.put("request-raw-url", rawUrl);
        this.pathParameters = urlTplParams;
        this.requestAdapter = requestAdapter;
    }
    /**
     * Gets an item from the petstore.utilities.store.order.item collection
     * @param id Unique identifier of the item
     * @return a WithOrderItemRequestBuilder
     */
    @javax.annotation.Nonnull
    public WithOrderItemRequestBuilder order(@javax.annotation.Nonnull final String id) {
        Objects.requireNonNull(id);
        final HashMap<String, Object> urlTplParams = new HashMap<String, Object>(this.pathParameters);
        urlTplParams.put("orderId", id);
        return new WithOrderItemRequestBuilder(urlTplParams, requestAdapter);
    }
}
