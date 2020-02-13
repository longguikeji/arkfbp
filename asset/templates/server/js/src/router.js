import fs from 'fs'

import { importWorkflowByFile, runWorkflow } from 'arkfbp/lib/flow'
import { executeHook } from 'arkfbp/lib/hook'
import { Response } from 'arkfbp/lib/response'

function r(app, appState, path, method, flowName) {
    const cb = async (req, res) => {
        await executeHook(appState, __dirname + '/flows/hooks/flow/beforeCreate')

        const flowFilename = __dirname + '/flows' + '/' + flowName
        console.info(flowFilename, '....>>')
        const ns = await importWorkflowByFile(flowFilename)
        const rr = new Response()
        const flow = new ns.Main({
            request: req,
            response: rr,
            appState,
        })

        await executeHook(appState, __dirname + '/flows/hooks/flow/created')
        const data = await runWorkflow(flow)
        await executeHook(appState, __dirname + '/flows/hooks/flow/executed')

        /**
         * catefully merge the original response propertities in ark with
         */

        const flowResponse = flow.response
        /**
         * merge status code
         */
        res.status(flowResponse.status)

        /**
         * merge headers
         */
        for (const key in flowResponse.headers) {
            if (flowResponse.headers.hasOwnProperty(key)) {
                const values = flowResponse.headers[key]
                if (Array.isArray(values)) {
                    for (const value of values) {
                        res.set(key, value)
                    }
                } else {
                    res.set(key, flowResponse.headers[key])
                }
            }
        }

        /**
         * merge data
         */
        if (flowResponse.data) {
            res.send(flowResponse.data)
        } else {
            res.send(data)
        }
    }

    switch (method) {
        case 'get':
            app.get(path, cb)
            break
        case 'post':
            app.post(path, cb)
            break
        default:
            return
    }
}

export default async function (app, appState) {

    app.get('/_routes/', async (req, res) => {
        let s = ''
        let idx = 0

        function getMethods(r) {
            const methods = []
            for (const method in r.methods) {
                if (r.methods.hasOwnProperty(method)) {
                    methods.push(method.toUpperCase())
                }
            }

            return methods
        }

        app._router.stack.forEach((r) => {
            if (r.route && r.route.path) {
                s += `${idx + 1} `

                for (const method of getMethods(r.route)) {
                    s += `[${method}]`
                }

                s += ' ' + r.route.path + '<br />'

                idx += 1
            }
        })

        res.send(s)
    })

    const routeFiles = fs.readdirSync(__dirname + '/routes')
    for (const filename of routeFiles) {
        if (filename.indexOf('.map') >= 0) {
            continue
        }

        const routes = await import(__dirname + '/routes' + '/' + filename)
        const namespace = routes.default.namespace
        routes.default.routes.forEach((route) => {
            for (const key in route) {
                if (typeof route[key] === 'string') {
                    r(app, appState, namespace + '/' + key, 'get', route[key])
                } else if (typeof route[key] === 'object') {
                    for (const method in route[key]) {
                        if (route[key].hasOwnProperty(method)) {
                            const flowName = route[key][method]
                            r(app, appState, namespace + '/' + key, method, flowName)
                        }
                    }
                }
            }
        })
    }
}