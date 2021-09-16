
#Fluxo
![fluxo grpcApi](docs/fluxo.png "Fluxo grpcApi")

#Como subir a aplicação

```shell
$ make start
```

# Rota
````shell
http://localhost:8000/checkout
````
### Exemplo de payload
```json
{
    "products": [
        {
            "id": 1,
            "quantity": 1
        },
			  {
            "id": 2,
            "quantity": 1
        }
    ]
} 
```

### Exemplo de retorno
```json
{
  "total_amount": 108968,
  "total_amount_with_discount": 103521,
  "total_discount": 5447,
  "products": [
    {
      "id": 2,
      "quantity": 1,
      "unit_amount": 93811,
      "total_amount": 93811,
      "discount": 4690,
      "is_gift": false
    },
    {
      "id": 1,
      "quantity": 1,
      "unit_amount": 15157,
      "total_amount": 15157,
      "discount": 757,
      "is_gift": false
    }
  ]
}
```

#Execução de testes
```shell
$ make test
```