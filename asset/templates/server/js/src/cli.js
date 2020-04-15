'use strict'

import yargs, { Argv } from 'yargs'

const path = require('path');
import { ark } from 'arkfbp/lib'
import { AppState } from 'arkfbp/lib/appState'
import { runWorkflowByFile } from 'arkfbp/lib/flow'
import { executeHook } from 'arkfbp/lib/hook'

import { serve } from './server'

function startServer(port, appState) {
    // console.info(`Serve on port ${port}.`)
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
    yargs.command('run', 'run flow', (yargs) => {
        yargs.option('name', {
            describe: 'Flow to execute',
            demand: true,
        }).option('inputs', {
            default: null,
            describe: 'Data to set as the inputs',
        }).option('debug', {
            default: false,
            describe: 'Enable debug mode or not',
            type: 'boolean',
        }).option('instance-id', {
            default: null,
            describe: 'instance id of the workflow',
        }).option('state-log-file', {
            default: null,
            describe: 'state log file',
        })
    }, (args) => {
        const flowDirectory = __dirname + '/flows'
        const flowFilename = flowDirectory + '/' + args.name
        const inputs = args.inputs
        runWorkflowByFile(flowFilename, inputs, {
            appState: gAppState,
            debug: args.debug,
            debugStatePersistentFile: args.stateLogFile,
        }).then((data) => {
            console.info(data)
        })
    })
        .command('test', 'run test flow', (yargs) => {
            yargs.option('name', {
                describe: 'Test flow to execute',
                demand: true,
            }).option('inputs', {
                default: null,
                describe: 'Data to set as the inputs',
            })
        }, (args) => {
            const flowDirectory = __dirname + '/testFlows'
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

