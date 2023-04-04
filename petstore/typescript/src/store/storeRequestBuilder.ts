import {InventoryRequestBuilder} from './inventory/inventoryRequestBuilder';
import {WithOrderItemRequestBuilder} from './order/item/withOrderItemRequestBuilder';
import {OrderRequestBuilder} from './order/orderRequestBuilder';
import {getPathParameters, RequestAdapter} from '@microsoft/kiota-abstractions';

/**
 * Builds and executes requests for operations under /store
 */
export class StoreRequestBuilder {
    /** The inventory property */
    public get inventory(): InventoryRequestBuilder {
        return new InventoryRequestBuilder(this.pathParameters, this.requestAdapter);
    }
    /** The order property */
    public get order(): OrderRequestBuilder {
        return new OrderRequestBuilder(this.pathParameters, this.requestAdapter);
    }
    /** Path parameters for the request */
    private pathParameters: Record<string, unknown>;
    /** The request adapter to use to execute the requests. */
    private requestAdapter: RequestAdapter;
    /** Url template to use to build the URL for the current request builder */
    private urlTemplate: string;
    /**
     * Instantiates a new StoreRequestBuilder and sets the default values.
     * @param pathParameters The raw url or the Url template parameters for the request.
     * @param requestAdapter The request adapter to use to execute the requests.
     */
    public constructor(pathParameters: Record<string, unknown> | string | undefined, requestAdapter: RequestAdapter) {
        if(!pathParameters) throw new Error("pathParameters cannot be undefined");
        if(!requestAdapter) throw new Error("requestAdapter cannot be undefined");
        this.urlTemplate = "{+baseurl}/store";
        const urlTplParams = getPathParameters(pathParameters);
        this.pathParameters = urlTplParams;
        this.requestAdapter = requestAdapter;
    };
    /**
     * Gets an item from the petstore.utilities.store.order.item collection
     * @param id Unique identifier of the item
     * @returns a WithOrderItemRequestBuilder
     */
    public orderById(id: string) : WithOrderItemRequestBuilder {
        if(!id) throw new Error("id cannot be undefined");
        const urlTplParams = getPathParameters(this.pathParameters);
        urlTplParams["orderId"] = id
        return new WithOrderItemRequestBuilder(urlTplParams, this.requestAdapter);
    };
}
