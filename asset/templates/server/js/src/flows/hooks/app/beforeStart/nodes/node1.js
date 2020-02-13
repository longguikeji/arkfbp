import { FunctionNode } from 'arkfbp/lib/functionNode'

export class Node1 extends FunctionNode {

    async run() {
        this.$appState.commit((state) => {
            return state
        })
    }

}