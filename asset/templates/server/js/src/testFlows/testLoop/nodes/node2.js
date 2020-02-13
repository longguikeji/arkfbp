import { StopNode } from 'arkfbp/lib/stopNode'

export class Node2 extends StopNode {

    async run() {
        return {
            'loop': this.state.fetch().sum,
        }
    }

}