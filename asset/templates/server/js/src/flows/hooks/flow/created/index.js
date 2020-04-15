import { Flow } from 'arkfbp/lib/flow'
import { Graph } from 'arkfbp/lib/graph'

import { Node1 } from './nodes/node1'

export class Main extends Flow {

    createNodes() {
        return [
            {
                cls: Node1,
                id: '1',
            },
        ]
    }

    createGraph() {
        const g = new Graph()
        g.nodes = this.createNodes()
        return g
    }

}