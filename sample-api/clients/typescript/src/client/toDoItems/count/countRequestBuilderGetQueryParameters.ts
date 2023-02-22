
/**
 * Get the number of the resource
 */
export class CountRequestBuilderGetQueryParameters {
    /** Filter items by property values */
    public filter?: string | undefined;
    /** Search items by search phrases */
    public search?: string | undefined;
    /**
     * Maps the query parameters names to their encoded names for the URI template parsing.
     * @param originalName The original query parameter name in the class.
     * @returns a string
     */
    public getQueryParameter(originalName: string | undefined) : string {
        if(!originalName) throw new Error("originalName cannot be undefined");
        switch(originalName) {
            case "filter": return "%24filter";
            case "search": return "%24search";
            default: return originalName;
        }
    };
}
