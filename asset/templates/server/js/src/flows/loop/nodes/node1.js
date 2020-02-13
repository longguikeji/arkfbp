import { LoopNode } from 'arkfbp/lib/loopNode'

export class Node1 extends LoopNode {

    _i = 0

    _sum = 0

    async initStatement() {
        this._i = 1
        this.state.commit((state) => {
            state.values = []
            return state
        })
    }

    async conditionStatement() {
        return this._i < 10
    }

    async postStatement() {
        this._i += 1
    }

    async process() {
        this._sum += this._i
        this.state.commit((state) => {
            state.sum = this._sum
            return state
        })
    }

}