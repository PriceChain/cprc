// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import PriceChainCprcPricechainCprcPrcibc from './PriceChain/cprc/pricechain.cprc.prcibc'
import PriceChainCprcPricechainCprcPrcmint from './PriceChain/cprc/pricechain.cprc.prcmint'
import PriceChainCprcPricechainCprcRegistry from './PriceChain/cprc/pricechain.cprc.registry'


export default { 
  PriceChainCprcPricechainCprcPrcibc: load(PriceChainCprcPricechainCprcPrcibc, 'pricechain.cprc.prcibc'),
  PriceChainCprcPricechainCprcPrcmint: load(PriceChainCprcPricechainCprcPrcmint, 'pricechain.cprc.prcmint'),
  PriceChainCprcPricechainCprcRegistry: load(PriceChainCprcPricechainCprcRegistry, 'pricechain.cprc.registry'),
  
}


function load(mod, fullns) {
    return function init(store) {        
        if (store.hasModule([fullns])) {
            throw new Error('Duplicate module name detected: '+ fullns)
        }else{
            store.registerModule([fullns], mod)
            store.subscribe((mutation) => {
                if (mutation.type == 'common/env/INITIALIZE_WS_COMPLETE') {
                    store.dispatch(fullns+ '/init', null, {
                        root: true
                    })
                }
            })
        }
    }
}
