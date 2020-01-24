package tpl

// Tpl9 ...
const Tpl9 = `
'use strict'

import bodyParser from 'body-parser'
import cookieParser from 'cookie-parser'
import express from 'express'
import formidableMiddleware from 'express-formidable'
import installRoutes from './router'

export async function serve(port, appState) {
    const app = express()

    app.use(cookieParser())
    app.use(formidableMiddleware())
    app.use(bodyParser.json({ limit: '1mb' }))
    app.use(bodyParser.urlencoded({
        extended: true,
    }))

    await installRoutes(app, appState)

    app.listen(port, () => {
        console.info(` + "`server started at :${port}`" + `)
    })

}
`
