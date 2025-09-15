# 🐳 Docker – Lancement et Profils

## ✅ Prérequis

- Docker et Docker Compose installés
- Être à la racine du dépôt (fichier `docker-compose.yml` présent)

---

## 🚀 Démarrage rapide (images finales)

Lance les services optimisés (backend et frontend) avec persistance des données.

```bash
# Depuis la racine du projet
docker compose up -d --build
```

- Backend exposé sur `http://localhost:3225`
- Frontend exposé sur `http://localhost:3000`
- Volumes persistants:
  - `backend_data` → base SQLite dans `/app/data/social.db`
  - `backend_uploads` → fichiers uploadés dans `/app/uploads`

Commandes utiles:

- Logs temps réel: `docker compose logs -f`
- Arrêt sans supprimer les volumes: `docker compose down`
- Arrêt + suppression des volumes: `docker compose down -v`

---

## 🛠️ Mode développement (hot‑reload)

Utilise les profils `dev` définis dans `docker-compose.yml` pour monter le code hôte dans les conteneurs et activer le rechargement à chaud.

```bash
# Lancer le stack de dev
docker compose --profile dev up -d --build

# Suivre les logs
docker compose --profile dev logs -f
```

- Services:
  - `backend-dev` (Go): lance `go run ./` avec volumes montés (`./backend:/app`)
  - `frontend-dev` (Next.js): lance `next dev` avec `npm ci` au démarrage et hot‑reload (`./frontend:/app`)
- Ports:
  - Backend: `3225`
  - Frontend: `3000`
- Particularités dev côté frontend:
  - `BACKEND_ORIGIN=http://backend-dev:3225`
  - `WATCHPACK_POLLING=true` et `CHOKIDAR_USEPOLLING=true` pour un watcher fiable en conteneur

Pour arrêter le mode dev:

```bash
docker compose --profile dev down
```

---

## 🔧 Détails de build et configuration

- Backend (Go): `backend/docker/Dockerfile`
  - Cible `final`: binaire statique + image distroless (runtime minimal)
  - Cible `dev`: dépendances build et `go run` pour itération rapide
  - Variables importantes: `PORT`, `DB_PATH`, `MIGRATIONS_PATH`
- Frontend (Next.js): `frontend/Dockerfile`
  - Cible `final`: build Next.js en mode standalone (image distroless node non‑root)
  - Cible `dev`: `npm ci` + `next dev`
  - `BACKEND_ORIGIN` est injecté au build (et utilisé pour le rewrite `/api/*` → backend), voir `frontend/next.config.mjs`

Rebuild forcé sans cache:

```bash
docker compose build --no-cache && docker compose up -d
```

---

## 🔌 Réseaux, ports et accès

- Les services communiquent via le réseau par défaut de Compose (`backend`, `frontend`, `*-dev`)
- Accès depuis l’hôte:
  - Frontend: `http://localhost:3000`
  - Backend: `http://localhost:3225`

---

## 🗂️ Volumes et persistance

Déclarés en bas de `docker-compose.yml`:

- `backend_data`: base SQLite persistée
- `backend_uploads`: répertoire des uploads
- `frontend_node_modules` (dev): préserver `node_modules` entre rebuilds dev

Pour repartir proprement (attention, destructif):

```bash
docker compose down -v && docker volume prune
```

---

## 🧪 Dépannage rapide

- Port déjà utilisé → changer `ports` dans `docker-compose.yml` ou libérer le port
- Changement d’API non pris en compte (frontend prod) → `docker compose up -d --build`
- Watcher Next ne détecte pas les changements → profil `dev` déjà configure le polling; vérifier que `frontend-dev` tourne
- Permissions d’écriture DB/uploads → gérées dans l’image backend finale (utilisateur root dans conteneur)

---

## 📌 Rappel des services (compose)

- Production: `backend`, `frontend`
- Développement: `backend-dev`, `frontend-dev` (profil `dev`)
