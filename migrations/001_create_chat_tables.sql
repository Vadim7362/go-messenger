-- Таблица чатов (групповых и личных)
CREATE TABLE chats (
  id SERIAL PRIMARY KEY,
  is_group BOOLEAN DEFAULT FALSE,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Кто участвует в каком чате (многие ко многим)
CREATE TABLE chat_members (
  id SERIAL PRIMARY KEY,
  chat_id INT REFERENCES chats(id) ON DELETE CASCADE,
  user_id INT REFERENCES users(id) ON DELETE CASCADE
);

-- Сообщения
CREATE TABLE messages (
  id SERIAL PRIMARY KEY,
  chat_id INT REFERENCES chats(id) ON DELETE CASCADE,
  sender_id INT REFERENCES users(id),
  content TEXT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);