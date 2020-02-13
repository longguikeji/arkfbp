import { Flow } from 'arkfbp/lib/flow'
import { Graph } from 'arkfbp/lib/graph'
import { StopNode } from 'arkfbp/lib/stopNode'

import { MyLoopNode } from './nodes/node1'
import { LoopBody, LoopBody2 } from './nodes/node2'

export class Main extends Flow {

    createGraph() {
        const g = new Graph()
        g.nodes = [
            {
                cls: MyLoopNode,
                id: 1,
                next: 4,
                body: [
                    {
                        cls: LoopBody,
                        id: 2,
                        next: LoopBody2,
                    },
                    {
                        cls: LoopBody2,
                        id: 3,
                    },
                ],
            },
            {
                'cls': StopNode,
                'id': 4,
            },
        ]

        return g
    }

}