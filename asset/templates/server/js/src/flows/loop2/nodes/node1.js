import { LoopNode } from 'arkfbp/lib/loopNode'

export class MyLoopNode extends LoopNode {

    _i = 0

    async initStatement() {
        this._i = 0
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
        this.state.commit((state) => {
            state.values.push(this._i)
            return state
        })
    }

}