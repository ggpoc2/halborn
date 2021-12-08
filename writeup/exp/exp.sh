halborncli tx halborn create-advertisement $ID $CONTENT --from $KEY --chain-id mainnet --fees 10ctc --node $RPC
halborncli tx sign tx.json --from $KEY --chain-id mainnet --node $RPC > signtx.json
halborncli tx broadcast signtx.json --node $RPC
halborncli tx halborn withdraw $ID 100ctc --from $KEY --chain-id mainnet --fees 10ctc --node $RPC
halborncli tx halborn ctf $ID --from $KEY --chain-id mainnet --fees 10ctc --node $RPC
