// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgModifyRegistry } from "./types/registry/tx";
import { MsgUnbondRegistry } from "./types/registry/tx";
import { MsgCreateRegistry } from "./types/registry/tx";
import { MsgProposePrice } from "./types/registry/tx";
import { MsgJoinRegistryCoOperator } from "./types/registry/tx";
import { MsgJoinRegistryMember } from "./types/registry/tx";


const types = [
  ["/pricechain.cprc.registry.MsgModifyRegistry", MsgModifyRegistry],
  ["/pricechain.cprc.registry.MsgUnbondRegistry", MsgUnbondRegistry],
  ["/pricechain.cprc.registry.MsgCreateRegistry", MsgCreateRegistry],
  ["/pricechain.cprc.registry.MsgProposePrice", MsgProposePrice],
  ["/pricechain.cprc.registry.MsgJoinRegistryCoOperator", MsgJoinRegistryCoOperator],
  ["/pricechain.cprc.registry.MsgJoinRegistryMember", MsgJoinRegistryMember],
  
];
export const MissingWalletError = new Error("wallet is required");

export const registry = new Registry(<any>types);

const defaultFee = {
  amount: [],
  gas: "200000",
};

interface TxClientOptions {
  addr: string
}

interface SignAndBroadcastOptions {
  fee: StdFee,
  memo?: string
}

const txClient = async (wallet: OfflineSigner, { addr: addr }: TxClientOptions = { addr: "http://localhost:26657" }) => {
  if (!wallet) throw MissingWalletError;
  let client;
  if (addr) {
    client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
  }else{
    client = await SigningStargateClient.offline( wallet, { registry });
  }
  const { address } = (await wallet.getAccounts())[0];

  return {
    signAndBroadcast: (msgs: EncodeObject[], { fee, memo }: SignAndBroadcastOptions = {fee: defaultFee, memo: ""}) => client.signAndBroadcast(address, msgs, fee,memo),
    msgModifyRegistry: (data: MsgModifyRegistry): EncodeObject => ({ typeUrl: "/pricechain.cprc.registry.MsgModifyRegistry", value: MsgModifyRegistry.fromPartial( data ) }),
    msgUnbondRegistry: (data: MsgUnbondRegistry): EncodeObject => ({ typeUrl: "/pricechain.cprc.registry.MsgUnbondRegistry", value: MsgUnbondRegistry.fromPartial( data ) }),
    msgCreateRegistry: (data: MsgCreateRegistry): EncodeObject => ({ typeUrl: "/pricechain.cprc.registry.MsgCreateRegistry", value: MsgCreateRegistry.fromPartial( data ) }),
    msgProposePrice: (data: MsgProposePrice): EncodeObject => ({ typeUrl: "/pricechain.cprc.registry.MsgProposePrice", value: MsgProposePrice.fromPartial( data ) }),
    msgJoinRegistryCoOperator: (data: MsgJoinRegistryCoOperator): EncodeObject => ({ typeUrl: "/pricechain.cprc.registry.MsgJoinRegistryCoOperator", value: MsgJoinRegistryCoOperator.fromPartial( data ) }),
    msgJoinRegistryMember: (data: MsgJoinRegistryMember): EncodeObject => ({ typeUrl: "/pricechain.cprc.registry.MsgJoinRegistryMember", value: MsgJoinRegistryMember.fromPartial( data ) }),
    
  };
};

interface QueryClientOptions {
  addr: string
}

const queryClient = async ({ addr: addr }: QueryClientOptions = { addr: "http://localhost:1317" }) => {
  return new Api({ baseUrl: addr });
};

export {
  txClient,
  queryClient,
};
