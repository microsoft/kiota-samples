import {Main} from './main';
import {ParseNode} from '@microsoft/kiota-abstractions';

export function createMainFromDiscriminatorValue(parseNode: ParseNode | undefined) : Main {
    if(!parseNode) throw new Error("parseNode cannot be undefined");
    return new Main();
}
