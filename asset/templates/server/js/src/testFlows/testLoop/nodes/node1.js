import { TestNode } from 'arkfbp/lib/testNode'
import assert from 'assert'

import { Main as QueryGithubRepoStarFlow } from './../../../flows/queryGithubRepoStar'

export class Node1 extends TestNode {

    flow = QueryGithubRepoStarFlow
    start = 1
    stop = 1

    setUp() {
        console.info('setUp')
    }

    tearDown() {
        console.info('tearDown')
    }

    testA() {
        console.info(this.outputs)
        assert.strictEqual(1, 1)
    }

    testB() {
        assert.strictEqual(1, 2)
    }

}