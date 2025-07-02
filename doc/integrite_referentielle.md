# 🔐 Règles d’Intégrité Référentielle – Projet Mellow

Ce document liste toutes les relations entre tables (clés étrangères) ainsi que les comportements définis pour les suppressions (`ON DELETE`) ou mises à jour (`ON UPDATE`).

---

## 🧑‍🤝‍🧑 users

Aucune dépendance entrante. C’est une table de référence pour plusieurs autres :

- `sessions.user_id` → ON DELETE CASCADE
- `posts.user_id` → ON DELETE CASCADE
- `comments.user_id` → ON DELETE CASCADE
- `notifications.user_id` → ON DELETE CASCADE
- `follow_requests.sender_id / receiver_id` → ON DELETE CASCADE
- `groups_member.user_id` → ON DELETE CASCADE
- `events_response.user_id` → ON DELETE CASCADE
- `messages.sender_id / receiver_id` → ON DELETE CASCADE
- `reports.user_id` → ON DELETE CASCADE
- `posts_viewer.user_id` → ON DELETE CASCADE

---

## 📬 sessions

- `sessions.user_id` → users(user_id)  
  🔁 `ON DELETE CASCADE`

---

## 📢 posts

- `posts.user_id` → users(user_id)  
  🔁 `ON DELETE CASCADE`
- `posts.group_id` → groups(group_id)  
  🔁 `ON DELETE SET NULL`

---

## 👁️ posts_viewer

- `posts_viewer.post_id` → posts(post_id)  
  🔁 `ON DELETE CASCADE`
- `posts_viewer.user_id` → users(user_id)  
  🔁 `ON DELETE CASCADE`

---

## 💬 comments

- `comments.user_id` → users(user_id)  
  🔁 `ON DELETE CASCADE`
- `comments.post_id` → posts(post_id)  
  🔁 `ON DELETE CASCADE`

---

## 🚩 reports

- `reports.post_id` → posts(post_id)  
  🔁 `ON DELETE CASCADE`
- `reports.user_id` → users(user_id)  
  🔁 `ON DELETE CASCADE`
- `reports.group_id` → groups(group_id)  
  🔁 `ON DELETE CASCADE`

---

## 👥 groups

- `groups.user_id` → users(user_id)  
  🔁 `ON DELETE CASCADE`

---

## 🧑‍💼 groups_member

- `groups_member.group_id` → groups(group_id)  
  🔁 `ON DELETE CASCADE`
- `groups_member.user_id` → users(user_id)  
  🔁 `ON DELETE CASCADE`

---

## 🔔 notifications

- `notifications.user_id` → users(user_id)  
  🔁 `ON DELETE CASCADE`

---

## 📨 messages

- `messages.sender_id` → users(user_id)  
  🔁 `ON DELETE CASCADE`
- `messages.receiver_id` → users(user_id)  
  🔁 `ON DELETE CASCADE`

---

## 📅 events

- `events.group_id` → groups(group_id)  
  🔁 `ON DELETE CASCADE`
- `events.user_id` → users(user_id)  
  🔁 `ON DELETE CASCADE`

---

## ✅ events_response

- `events_response.event_id` → events(event_id)  
  🔁 `ON DELETE CASCADE`
- `events_response.user_id` → users(user_id)  
  🔁 `ON DELETE CASCADE`
- `events_response.group_id` → groups(group_id)  
  🔁 `ON DELETE CASCADE`

---

## ➕ follow_requests

- `follow_requests.sender_id` → users(user_id)  
  🔁 `ON DELETE CASCADE`
- `follow_requests.receiver_id` → users(user_id)  
  🔁 `ON DELETE CASCADE`
- `follow_requests.group_id` → groups(group_id)  
  🔁 `ON DELETE CASCADE`

---

🎯 Ces règles garantissent que les données liées ne deviennent jamais orphelines et que les suppressions se propagent correctement.
