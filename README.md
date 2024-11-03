# API Machinery

Zdrojové kódy z článku p *API Machinery* na [okontajneroch.sk](https://okontajneroch.sk). Na spustenie je potrebné 
mať nainštalované:

- Go 1.23
- deepcopy-gen (`go install k8s.io/code-generator/cmd/deepcopy-gen@latest`)

Príklady z článku su rozdelene v adresároch, v `./examples`. Ak chceme spustiť príklad z adresára 
`./examples/api-machinery-serialize`, postupujeme nasledovne:

```bash
go run ./examples/api-machinery-serialize
```

Ak chceme sputiť `./examples/api-machinery-deserialize`, postupujeme nasledovne:

```bash
go run ./examples/api-machinery-deserialize << EOF
apiVersion: okontajneroch.sk/v1
kind: Starfighter
metadata:
   name: x-wing-1
spec:
   faction: rebellion
   type: x-wing
   pilot: Luke Skywalker
EOF
```
