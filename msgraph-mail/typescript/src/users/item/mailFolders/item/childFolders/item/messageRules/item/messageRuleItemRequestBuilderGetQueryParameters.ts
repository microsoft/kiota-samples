
<<<<<<< HEAD
export interface MessageRuleItemRequestBuilderGetQueryParameters {
    /** Select properties to be returned */
    select?: string[] | undefined;
=======
/** The collection of rules that apply to the user's Inbox folder. */
export class MessageRuleItemRequestBuilderGetQueryParameters {
    /** Select properties to be returned */
    public select?: string[] | undefined;
    /**
     * Maps the query parameters names to their encoded names for the URI template parsing.
     * @param originalName The original query parameter name in the class.
     * @returns a string
     */
    public getQueryParameter(originalName: string | undefined) : string {
        if(!originalName) throw new Error("originalName cannot be undefined");
        switch(originalName) {
            case "select": return "%24select";
            default: return originalName;
        }
    };
>>>>>>> main
}
