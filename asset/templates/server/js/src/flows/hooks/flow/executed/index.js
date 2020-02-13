import { Flow } from 'arkfbp/lib/flow'
import { Graph } from 'arkfbp/lib/graph'

import { Node1 } from './nodes/node1'

export class Main extends Flow {

    createGraph() {
        const g = new Graph()
        g.nodes = [
            {
                cls: Node1,
                id: 1,
            },
        ]

        return g
    }

}