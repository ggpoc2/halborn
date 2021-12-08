# You can run all of these commands from your home directory
cd $HOME
rm -rf .halbornd
rm -rf .halborncli

# Initialize the genesis.json file that will help you to bootstrap the network
halbornd init testing --chain-id=testnet

# Create a key to hold your validator account
halborncli keys add validator
halborncli keys add ctfer

halborncli config indent true
halborncli config output json
halborncli config trust-node true

# Add that key into the genesis.app_state.accounts array in the genesis file
# NOTE: this command lets you set the number of coins. Make sure this account has some coins
# with the genesis.app_state.staking.params.bond_denom denom, the default is staking
halbornd add-genesis-account $(halborncli keys show validator -a) 1000000000stake
halbornd add-genesis-account $(halborncli keys show ctfer -a) 1000ctc

# Generate the transaction that creates your validator
halbornd gentx --name validator

# Add the generated bonding transaction to the genesis file
halbornd collect-gentxs

# Now its safe to start `halbornd`
halbornd start
