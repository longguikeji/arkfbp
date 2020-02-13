import { FunctionNode } from 'arkfbp-browser/lib/functionNode'

export class SayHi extends FunctionNode {

    async run() {
        console.info('say hi')
    }

}