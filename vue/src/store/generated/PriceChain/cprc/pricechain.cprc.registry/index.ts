import { txClient, queryClient, MissingWalletError , registry} from './module'

import { Params } from "./module/types/registry/params"
import { Registry } from "./module/types/registry/registry"
import { RegistryMember } from "./module/types/registry/registry_member"
import { RegistryOwner } from "./module/types/registry/registry_owner"
import { RegistryStakedAmount } from "./module/types/registry/registry_staked_amount"
import { StakedAmountPerWallet } from "./module/types/registry/staked_amount_per_wallet"


export { Params, Registry, RegistryMember, RegistryOwner, RegistryStakedAmount, StakedAmountPerWallet };

async function initTxClient(vuexGetters) {
	return await txClient(vuexGetters['common/wallet/signer'], {
		addr: vuexGetters['common/env/apiTendermint']
	})
}

async function initQueryClient(vuexGetters) {
	return await queryClient({
		addr: vuexGetters['common/env/apiCosmos']
	})
}

function mergeResults(value, next_values) {
	for (let prop of Object.keys(next_values)) {
		if (Array.isArray(next_values[prop])) {
			value[prop]=[...value[prop], ...next_values[prop]]
		}else{
			value[prop]=next_values[prop]
		}
	}
	return value
}

function getStructure(template) {
	let structure = { fields: [] }
	for (const [key, value] of Object.entries(template)) {
		let field: any = {}
		field.name = key
		field.type = typeof value
		structure.fields.push(field)
	}
	return structure
}

const getDefaultState = () => {
	return {
				Params: {},
				Registry: {},
				RegistryAll: {},
				RegistryOwner: {},
				RegistryOwnerAll: {},
				RegistryMember: {},
				RegistryMemberAll: {},
				RegistryStakedAmount: {},
				RegistryStakedAmountAll: {},
				StakedAmountPerWallet: {},
				StakedAmountPerWalletAll: {},
				
				_Structure: {
						Params: getStructure(Params.fromPartial({})),
						Registry: getStructure(Registry.fromPartial({})),
						RegistryMember: getStructure(RegistryMember.fromPartial({})),
						RegistryOwner: getStructure(RegistryOwner.fromPartial({})),
						RegistryStakedAmount: getStructure(RegistryStakedAmount.fromPartial({})),
						StakedAmountPerWallet: getStructure(StakedAmountPerWallet.fromPartial({})),
						
		},
		_Registry: registry,
		_Subscriptions: new Set(),
	}
}

// initial state
const state = getDefaultState()

export default {
	namespaced: true,
	state,
	mutations: {
		RESET_STATE(state) {
			Object.assign(state, getDefaultState())
		},
		QUERY(state, { query, key, value }) {
			state[query][JSON.stringify(key)] = value
		},
		SUBSCRIBE(state, subscription) {
			state._Subscriptions.add(JSON.stringify(subscription))
		},
		UNSUBSCRIBE(state, subscription) {
			state._Subscriptions.delete(JSON.stringify(subscription))
		}
	},
	getters: {
				getParams: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Params[JSON.stringify(params)] ?? {}
		},
				getRegistry: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Registry[JSON.stringify(params)] ?? {}
		},
				getRegistryAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.RegistryAll[JSON.stringify(params)] ?? {}
		},
				getRegistryOwner: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.RegistryOwner[JSON.stringify(params)] ?? {}
		},
				getRegistryOwnerAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.RegistryOwnerAll[JSON.stringify(params)] ?? {}
		},
				getRegistryMember: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.RegistryMember[JSON.stringify(params)] ?? {}
		},
				getRegistryMemberAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.RegistryMemberAll[JSON.stringify(params)] ?? {}
		},
				getRegistryStakedAmount: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.RegistryStakedAmount[JSON.stringify(params)] ?? {}
		},
				getRegistryStakedAmountAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.RegistryStakedAmountAll[JSON.stringify(params)] ?? {}
		},
				getStakedAmountPerWallet: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.StakedAmountPerWallet[JSON.stringify(params)] ?? {}
		},
				getStakedAmountPerWalletAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.StakedAmountPerWalletAll[JSON.stringify(params)] ?? {}
		},
				
		getTypeStructure: (state) => (type) => {
			return state._Structure[type].fields
		},
		getRegistry: (state) => {
			return state._Registry
		}
	},
	actions: {
		init({ dispatch, rootGetters }) {
			console.log('Vuex module: pricechain.cprc.registry initialized!')
			if (rootGetters['common/env/client']) {
				rootGetters['common/env/client'].on('newblock', () => {
					dispatch('StoreUpdate')
				})
			}
		},
		resetState({ commit }) {
			commit('RESET_STATE')
		},
		unsubscribe({ commit }, subscription) {
			commit('UNSUBSCRIBE', subscription)
		},
		async StoreUpdate({ state, dispatch }) {
			state._Subscriptions.forEach(async (subscription) => {
				try {
					const sub=JSON.parse(subscription)
					await dispatch(sub.action, sub.payload)
				}catch(e) {
					throw new Error('Subscriptions: ' + e.message)
				}
			})
		},
		
		
		
		 		
		
		
		async QueryParams({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryParams()).data
				
					
				commit('QUERY', { query: 'Params', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryParams', payload: { options: { all }, params: {...key},query }})
				return getters['getParams']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryParams API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryRegistry({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryRegistry( key.id)).data
				
					
				commit('QUERY', { query: 'Registry', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryRegistry', payload: { options: { all }, params: {...key},query }})
				return getters['getRegistry']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryRegistry API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryRegistryAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryRegistryAll(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.queryRegistryAll({...query, 'pagination.key':(<any> value).pagination.next_key})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'RegistryAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryRegistryAll', payload: { options: { all }, params: {...key},query }})
				return getters['getRegistryAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryRegistryAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryRegistryOwner({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryRegistryOwner( key.id)).data
				
					
				commit('QUERY', { query: 'RegistryOwner', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryRegistryOwner', payload: { options: { all }, params: {...key},query }})
				return getters['getRegistryOwner']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryRegistryOwner API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryRegistryOwnerAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryRegistryOwnerAll(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.queryRegistryOwnerAll({...query, 'pagination.key':(<any> value).pagination.next_key})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'RegistryOwnerAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryRegistryOwnerAll', payload: { options: { all }, params: {...key},query }})
				return getters['getRegistryOwnerAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryRegistryOwnerAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryRegistryMember({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryRegistryMember( key.id)).data
				
					
				commit('QUERY', { query: 'RegistryMember', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryRegistryMember', payload: { options: { all }, params: {...key},query }})
				return getters['getRegistryMember']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryRegistryMember API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryRegistryMemberAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryRegistryMemberAll(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.queryRegistryMemberAll({...query, 'pagination.key':(<any> value).pagination.next_key})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'RegistryMemberAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryRegistryMemberAll', payload: { options: { all }, params: {...key},query }})
				return getters['getRegistryMemberAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryRegistryMemberAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryRegistryStakedAmount({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryRegistryStakedAmount( key.index)).data
				
					
				commit('QUERY', { query: 'RegistryStakedAmount', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryRegistryStakedAmount', payload: { options: { all }, params: {...key},query }})
				return getters['getRegistryStakedAmount']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryRegistryStakedAmount API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryRegistryStakedAmountAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryRegistryStakedAmountAll(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.queryRegistryStakedAmountAll({...query, 'pagination.key':(<any> value).pagination.next_key})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'RegistryStakedAmountAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryRegistryStakedAmountAll', payload: { options: { all }, params: {...key},query }})
				return getters['getRegistryStakedAmountAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryRegistryStakedAmountAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryStakedAmountPerWallet({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryStakedAmountPerWallet( key.index)).data
				
					
				commit('QUERY', { query: 'StakedAmountPerWallet', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryStakedAmountPerWallet', payload: { options: { all }, params: {...key},query }})
				return getters['getStakedAmountPerWallet']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryStakedAmountPerWallet API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryStakedAmountPerWalletAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryStakedAmountPerWalletAll(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.queryStakedAmountPerWalletAll({...query, 'pagination.key':(<any> value).pagination.next_key})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'StakedAmountPerWalletAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryStakedAmountPerWalletAll', payload: { options: { all }, params: {...key},query }})
				return getters['getStakedAmountPerWalletAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryStakedAmountPerWalletAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		async sendMsgProposePrice({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgProposePrice(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgProposePrice:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgProposePrice:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgJoinRegistryMember({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgJoinRegistryMember(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgJoinRegistryMember:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgJoinRegistryMember:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgCreateRegistry({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgCreateRegistry(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateRegistry:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgCreateRegistry:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgJoinRegistryCoOperator({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgJoinRegistryCoOperator(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgJoinRegistryCoOperator:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgJoinRegistryCoOperator:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgModifyRegistry({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgModifyRegistry(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgModifyRegistry:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgModifyRegistry:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgUnbondRegistry({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgUnbondRegistry(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUnbondRegistry:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgUnbondRegistry:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		
		async MsgProposePrice({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgProposePrice(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgProposePrice:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgProposePrice:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgJoinRegistryMember({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgJoinRegistryMember(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgJoinRegistryMember:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgJoinRegistryMember:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgCreateRegistry({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgCreateRegistry(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateRegistry:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgCreateRegistry:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgJoinRegistryCoOperator({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgJoinRegistryCoOperator(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgJoinRegistryCoOperator:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgJoinRegistryCoOperator:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgModifyRegistry({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgModifyRegistry(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgModifyRegistry:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgModifyRegistry:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgUnbondRegistry({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgUnbondRegistry(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUnbondRegistry:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgUnbondRegistry:Create Could not create message: ' + e.message)
				}
			}
		},
		
	}
}
