package graphjavav4.utilities.users.item;

import com.microsoft.kiota.HttpCore;
import graphjavav4.utilities.users.item.inferenceClassification.InferenceClassificationRequestBuilder;
import graphjavav4.utilities.users.item.mailFolders.item.MailFolderRequestBuilder;
import graphjavav4.utilities.users.item.mailFolders.MailFoldersRequestBuilder;
import graphjavav4.utilities.users.item.messages.item.MessageRequestBuilder;
import graphjavav4.utilities.users.item.messages.MessagesRequestBuilder;
import java.util.Objects;
/** Builds and executes requests for operations under /users/{user-id}  */
public class UserRequestBuilder {
    /** Current path for the request  */
    private final String currentPath;
    /** The http core service to use to execute the requests.  */
    private final HttpCore httpCore;
    @javax.annotation.Nonnull
    public InferenceClassificationRequestBuilder inferenceClassification() {
        return new InferenceClassificationRequestBuilder(currentPath + pathSegment, httpCore, false);
    }
    /** Whether the current path is a raw URL  */
    private final boolean isRawUrl;
    @javax.annotation.Nonnull
    public MailFoldersRequestBuilder mailFolders() {
        return new MailFoldersRequestBuilder(currentPath + pathSegment, httpCore, false);
    }
    @javax.annotation.Nonnull
    public MessagesRequestBuilder messages() {
        return new MessagesRequestBuilder(currentPath + pathSegment, httpCore, false);
    }
    /** Path segment to use to build the URL for the current request builder  */
    private final String pathSegment;
    /**
     * Instantiates a new UserRequestBuilder and sets the default values.
     * @param currentPath Current path for the request
     * @param httpCore The http core service to use to execute the requests.
     * @return a void
     */
    public UserRequestBuilder(final String currentPath, final HttpCore httpCore) {
        this(currentPath, httpCore, true);
    }
    /**
     * Instantiates a new UserRequestBuilder and sets the default values.
     * @param currentPath Current path for the request
     * @param httpCore The http core service to use to execute the requests.
     * @param isRawUrl Whether the current path is a raw URL
     * @return a void
     */
    public UserRequestBuilder(@javax.annotation.Nonnull final String currentPath, @javax.annotation.Nonnull final HttpCore httpCore, final boolean isRawUrl) {
        Objects.requireNonNull(currentPath);
        Objects.requireNonNull(httpCore);
        this.pathSegment = "";
        this.httpCore = httpCore;
        this.currentPath = currentPath;
        this.isRawUrl = isRawUrl;
    }
    /**
     * Gets an item from the graphjavav4.utilities.users.item.mailFolders.item collection
     * @param id Unique identifier of the item
     * @return a mailFolderRequestBuilder
     */
    @javax.annotation.Nonnull
    public MailFolderRequestBuilder mailFolders(@javax.annotation.Nonnull final String id) {
        Objects.requireNonNull(id);
        return new MailFolderRequestBuilder(currentPath + pathSegment + "/mailFolders/" + id, httpCore, false);
    }
    /**
     * Gets an item from the graphjavav4.utilities.users.item.messages.item collection
     * @param id Unique identifier of the item
     * @return a messageRequestBuilder
     */
    @javax.annotation.Nonnull
    public MessageRequestBuilder messages(@javax.annotation.Nonnull final String id) {
        Objects.requireNonNull(id);
        return new MessageRequestBuilder(currentPath + pathSegment + "/messages/" + id, httpCore, false);
    }
}
