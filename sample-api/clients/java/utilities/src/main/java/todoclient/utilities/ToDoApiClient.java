package todoclient.utilities;

import com.microsoft.kiota.ApiClientBuilder;
import com.microsoft.kiota.RequestAdapter;
import com.microsoft.kiota.serialization.JsonParseNodeFactory;
import com.microsoft.kiota.serialization.JsonSerializationWriterFactory;
import com.microsoft.kiota.serialization.ParseNodeFactoryRegistry;
import com.microsoft.kiota.serialization.SerializationWriterFactoryRegistry;
import java.util.HashMap;
import java.util.Objects;
import todoclient.utilities.todoitems.item.ToDoItemItemRequestBuilder;
import todoclient.utilities.todoitems.ToDoItemsRequestBuilder;
/** The main entry point of the SDK, exposes the configuration and the fluent API.  */
public class ToDoApiClient {
    /** Path parameters for the request  */
    private final HashMap<String, Object> pathParameters;
    /** The request adapter to use to execute the requests.  */
    private final RequestAdapter requestAdapter;
    @javax.annotation.Nonnull
    public ToDoItemsRequestBuilder toDoItems() {
        return new ToDoItemsRequestBuilder(pathParameters, requestAdapter);
    }
    /** Url template to use to build the URL for the current request builder  */
    private final String urlTemplate;
    /**
     * Instantiates a new ToDoApiClient and sets the default values.
     * @param requestAdapter The request adapter to use to execute the requests.
     * @return a void
     */
    public ToDoApiClient(@javax.annotation.Nonnull final RequestAdapter requestAdapter) {
        Objects.requireNonNull(requestAdapter);
        this.pathParameters = new HashMap<>();
        this.urlTemplate = "{+baseurl}";
        this.requestAdapter = requestAdapter;
        ApiClientBuilder.registerDefaultSerializer(JsonSerializationWriterFactory.class);
        ApiClientBuilder.registerDefaultDeserializer(JsonParseNodeFactory.class);
        requestAdapter.setBaseUrl("https://localhost:7206");
    }
    /**
     * Gets an item from the todoclient.utilities.ToDoItems.item collection
     * @param id Unique identifier of the item
     * @return a ToDoItemItemRequestBuilder
     */
    @javax.annotation.Nonnull
    public ToDoItemItemRequestBuilder toDoItems(@javax.annotation.Nonnull final String id) {
        Objects.requireNonNull(id);
        var urlTplParams = new HashMap<String, Object>(this.pathParameters);
        urlTplParams.put("ToDoItem_id", id);
        return new ToDoItemItemRequestBuilder(urlTplParams, requestAdapter);
    }
}
