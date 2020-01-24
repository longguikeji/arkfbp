package tpl

// Tpl7 ...
const Tpl7 = `
'use strict'

import yargs, { Argv } from 'yargs'

const path = require('path');
import { ark } from 'arkfbp/lib'
import { AppState } from 'arkfbp/lib/appState'
import { runWorkflowByFile } from 'arkfbp/lib/flow'
import { executeHook } from 'arkfbp/lib/hook'

// import Logger from './plugins/log'
import { serve } from './server'

// ark.registerPlugin(Logger)

function startServer(port, appState) {
    serve(Number(port), appState)
}

const gAppState = new AppState()

async function start(appState) {
    /**
     * 注册App级别的Hook
     * App的StartupFlow仅会在server启动的时候执行一次，你可以任意设置App的State
     * Server启动之后AppState会被引用
     */
    await executeHook(appState, path.resolve(__dirname, '/flows/hooks/app/beforeStart'))
    await executeHook(appState, path.resolve(__dirname, '/flows/hooks/app/started'))

    /**
     * 注册中间件
     */
}

start(gAppState).then(() => {
    // require('yargs')
    yargs.command('run', 'run workflow', (yargs) => {
        yargs.option('name', {
            describe: 'Workflow to execute',
            demand: true,
        }).option('inputs', {
            default: null,
            describe: 'Data to set as the inputs',
        })
    }, (args) => {
        const flowDirectory = __dirname + '/flows'
        const flowFilename = flowDirectory + '/' + args.name
        const inputs = args.inputs
        runWorkflowByFile(flowFilename, inputs, {
            appState: gAppState,
        }).then((data) => {
            console.info(data)
        })
    })
        .command('serve', 'Start the server.', (yargs) => {
            yargs.option('port', {
                describe: 'Port to bind on',
                default: '3000',
            })
        }, (args) => {
            startServer(args.port, gAppState)
        }).argv
})


`
