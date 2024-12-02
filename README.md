# Service_des_emprunts
L'application de gestion d'une bibliothèque


## API Reference

#### PUT emprunts : rendre un livre

```http
  PUT /

```

```json
  {
    "empruntId":5, 
    "returned" : true
}
```
## MQTT Reference
#### Etre notifié quand un livre est rendu

- queue : "emprunts_finished_queue"
 
- routing key : "emprunts.v1.finished"

example de rendu 

```json
  {
    "CreatedAt": "2024-12-02T12:10:44.741317Z",
    "DateEmprunt": "2024-11-10T10:00:00Z",
    "DateRetourEffectif": "2024-12-02T21:19:02.087307Z",
    "DateRetourPrevu": "2024-11-25T10:00:00Z",
    "DeletedAt": null,
    "IDEmprunt": 5,
    "LivreID": 105,
    "UpdatedAt": "2024-12-02T21:19:02.099121Z",
    "UtilisateurID": 1
}
```

