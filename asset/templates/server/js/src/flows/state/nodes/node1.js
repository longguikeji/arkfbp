import { FunctionNode } from 'arkfbp/lib/functionNode'


export class Node1 extends FunctionNode {

    async run() {
        this.state.commit((state) => {
            state.now = new Date()
            return state
        })

        return 10
    }

}