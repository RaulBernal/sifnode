#!/bin/sh

SIFCHAIN_ID=sifchain-1 \
  COSMOS_BASE_DENOM=uatom \
  COSMOS_CHANNEL_ID=channel-0 \
  COSMOS_COUNTERPARTY_CHANNEL_ID=channel-192 \
  COSMOS_CHAIN_ID=cosmoshub-4 \
  AKASH_CHANNEL_ID=channel-2   \
  AKASH_COUNTERPARTY_CHANNEL_ID=channel-24 \
  AKASH_CHAIN_ID=akashnet-2 \
  PERSISTENCE_CHANNEL_ID=channel-7 \
  PERSISTENCE_COUNTERPARTY_CHANNEL_ID=channel-26 \
  PERSISTENCE_CHAIN_ID=core-1 \
  IRIS_BASE_DENOM=uiris \
  IRIS_CHANNEL_ID=channel-8 \
  IRIS_COUNTERPARTY_CHANNEL_ID=channel-19 \
  IRIS_CHAIN_ID=irishub-1 \
  SENTINEL_CHANNEL_ID=channel-1 \
  SENTINEL_COUNTERPARTY_CHANNEL_ID=channel-36 \
  SENTINEL_CHAIN_ID=sentinelhub-2 \
  REGEN_CHAIN_ID=regen-1 \
  REGEN_CHANNEL_ID=channel-10 \
  REGEN_COUNTERPARTY_CHANNEL_ID=channel-28 \
  CRYPTO_ORG_CHAIN_ID=crypto-org-chain-mainnet-1 \
  CRYPTO_ORG_CHANNEL_ID=channel-9 \
  CRYPTO_ORG_COUNTERPARTY_CHANNEL_ID=channel-33 ./template/generate-all-ibc.sh