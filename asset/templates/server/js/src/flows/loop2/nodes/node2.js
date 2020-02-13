import { FunctionNode } from 'arkfbp/lib/functionNode'


export class LoopBody extends FunctionNode {

    async run() {
        console.info('loop body 1')
    }

}

export class LoopBody2 extends FunctionNode {

    async run() {
        console.info('loop body 2')
    }

}