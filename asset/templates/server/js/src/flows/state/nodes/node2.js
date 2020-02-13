import { FunctionNode } from 'arkfbp/lib/functionNode'

export class Node2 extends FunctionNode {

    async run() {
        console.info(this.state.getInputs('1'))
        console.info(this.state.getOutputs('1'))
        console.info(this.state.fetch().now)
    }

}