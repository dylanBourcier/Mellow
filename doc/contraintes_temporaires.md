# ⏳ Contraintes temporelles – Projet Mellow

Ce document décrit les règles de cohérence à appliquer sur les champs de type date ou datetime dans la base de données, afin d’assurer la validité temporelle des données.

---

## 🧑 Utilisateurs (`users`)

- `birthdate` doit être :
  - antérieure à la date actuelle
  - cohérente avec un âge minimal : **13 ans requis**
- À valider côté **backend Go**, car SQLite ne gère pas de `CHECK (NOW() > ...)`

```go
if birthDate.After(time.Now()) {
    return errors.New("La date de naissance ne peut pas être dans le futur")
}
```

---

## 🗓️ Événements (`events`)

- `event_date` doit être dans le **futur** ou le **présent**
- Empêcher la création d’un événement dans le passé
- Vérifier que la réponse (`event_response`) arrive **avant ou à la date de l’événement**

```go
if eventDate.Before(time.Now()) {
    return errors.New("Un événement ne peut pas être créé dans le passé")
}
```

---

## 📝 Posts et commentaires (`posts`, `comments`)

- `creation_date` doit être ≤ `time.Now()`
- Les dates sont généralement générées automatiquement par le backend

---

## 📩 Messages (`messages`)

- `creation_date` ne peut pas être dans le futur
- Si horodatés manuellement (par test ou import), une validation est recommandée

---

## 📌 Remarques

- SQLite ne permet **pas** de `CHECK (field <= CURRENT_DATE)`
- Ces contraintes doivent être **appliquées côté application (Go)** pour être fiables
- Penser à tester ces cas via des **tests unitaires / fonctionnels**

---

