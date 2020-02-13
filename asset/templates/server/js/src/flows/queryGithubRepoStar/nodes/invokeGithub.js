import { APINode } from 'arkfbp/lib/apiNode'

export class InvokeGithub extends APINode {

    url = 'https://api.github.com/repos/longguikeji/arkid-core/stargazers'

    async run() {
        const data = this.$state.fetch()
        if (data.link) {
            this.url = data.link
        }

        await super.run()

        this.$state.commit((state) => {
            if (typeof state.datumx === 'undefined') {
                state.datumx = []
            }
            state.datumx.push(this.resp.data)
            return state
        })

        console.info('invoked github')

        return this.resp
    }
}