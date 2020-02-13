import { Ark } from 'arkfbp/lib/index'
import { debug } from 'debug'


export default function install(ark, options) {
    Object.defineProperty(Ark.prototype, 'logger', {
        get: () => {
            return debug
        },
        enumerable: true,
        configurable: true,
    })

}