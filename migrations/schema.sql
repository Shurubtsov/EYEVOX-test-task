CREATE TABLE chats(
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) UNIQUE NOT NULL,
    founder_nickname VARCHAR(50) NOT NULL
);

CREATE TABLE messages(
    id SERIAL PRIMARY KEY,
    creator_nickname VARCHAR(50) NOT NULL,
    chat_id INTEGER NOT NULL,
    text_message TEXT,
    CONSTRAINT fk_chatid
        FOREIGN KEY(chat_id) REFERENCES chats(id)
)