import { useState } from "react";
import axios from "axios";

export default function NoteCreate() {
  const [content, setContent] = useState("");
  const [expiresAt, setExpiresAt] = useState("");
  const [isPrivate, setIsPrivate] = useState(false);
  const [password, setPassword] = useState("");

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    try {
      const token = localStorage.getItem("access_token");
      await axios.post(
        "http://localhost:2288/note",
        {
          content,
          expires_at: new Date(expiresAt),
          is_private: isPrivate,
          password: isPrivate ? password : undefined,
        },
        {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        }
      );

      alert("Заметка создана!");
      setContent("");
      setExpiresAt("");
      setIsPrivate(false);
      setPassword("");
    } catch (err: any) {
      alert("Ошибка создания заметки: " + err?.response?.data || "неизвестная ошибка");
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      <textarea
        placeholder="Содержимое заметки"
        value={content}
        onChange={(e) => setContent(e.target.value)}
        required
      />
      <input
        type="datetime-local"
        value={expiresAt}
        onChange={(e) => setExpiresAt(e.target.value)}
        required
      />
      <label>
        Приватная?
        <input
          type="checkbox"
          checked={isPrivate}
          onChange={(e) => setIsPrivate(e.target.checked)}
        />
      </label>
      {isPrivate && (
        <input
          type="password"
          placeholder="Пароль"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          required
        />
      )}
      <button type="submit">Создать заметку</button>
    </form>
  );
}
