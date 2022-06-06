thank you
https://github.com/Kong/deck/issues/371#issuecomment-838473146
``` docker run --rm --network=kong_net -e DECK_KONG_ADDR=http://kong:8001 -v ${PWD}/:/mnt/deck kong/deck dump -o /mnt/deck/dump.yaml ```