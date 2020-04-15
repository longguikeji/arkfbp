
import { Flow } from 'arkfbp/lib/flow'
import { Graph } from 'arkfbp/lib/graph'
import { StartNode } from 'arkfbp/lib/StartNode'
import { StopNode } from 'arkfbp/lib/StopNode'

import { Node1 } from './nodes/node1'

export class Main extends Flow {

    createNodes() {
        return [
            {
                cls: StartNode,
                id: '1',
                next: '2',
            },
            {
                cls: Node1,
                id: '2',
                next: '3',
            },
            {
                cls: StopNode,
                id: '3',
            },
        ]
    }

    createGraph() {
        const g = new Graph()
        g.nodes = this.createNodes()
        return g
    }

}