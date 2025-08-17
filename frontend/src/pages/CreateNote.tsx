import { useState } from "react";
import { useNavigate } from "react-router-dom";
import api from "../services/api";
import { useTheme } from "../hooks/useTheme";

export default function CreateNote() {
  const navigate = useNavigate();
  const { theme } = useTheme();

  const [content, setContent] = useState("");
  const [expiresAt, setExpiresAt] = useState("");
  const [isPrivate, setIsPrivate] = useState(false);
  const [password, setPassword] = useState("");
  const [error, setError] = useState("");

  const handleCreate = async () => {
    setError("");

    if (isPrivate && password.trim() === "") {
      setError("Введите пароль для приватной заметки");
      return;
    }

    try {
      await api.post("/notes", {
        content,
        expires_at: new Date(expiresAt).toISOString(),
        is_private: isPrivate, 
        password: isPrivate ? password : undefined, 
      });

      navigate("/dashboard");
    } catch (err: any) {
      setError(err.response?.data || "Ошибка при создании заметки");
    }
  };

  return (
    <div className="relative flex items-center justify-center min-h-screen bg-[var(--bg-color)] text-[var(--text-color)] transition-colors px-4 overflow-hidden">
      <div
        className={`pointer-events-none absolute top-0 left-1/2 transform -translate-x-1/2 w-[600px] h-[600px] rounded-full blur-[140px] opacity-20 z-0 ${
          theme === "dark" ? "bg-purple-700" : "bg-blue-300"
        }`}
      />

      <div className="relative z-10 w-full max-w-lg bg-[var(--card-bg)] p-8 rounded-2xl shadow-xl">
        <h2 className="text-2xl font-bold mb-6">Создать заметку</h2>

        <textarea
          rows={5}
          placeholder="Содержание заметки..."
          value={content}
          onChange={(e) => setContent(e.target.value)}
          className="w-full p-3 rounded-lg border border-gray-300 dark:border-gray-600 bg-white dark:bg-gray-800 dark:text-white mb-4"
        />

        <input
          type="datetime-local"
          value={expiresAt}
          onChange={(e) => setExpiresAt(e.target.value)}
          className="w-full p-3 rounded-lg border border-gray-300 dark:border-gray-600 bg-white dark:bg-gray-800 dark:text-white mb-4"
        />

        <label className="flex items-center gap-2 mb-4">
          <input
            type="checkbox"
            checked={isPrivate}
            onChange={(e) => setIsPrivate(e.target.checked)}
          />
          Приватная заметка
        </label>

        {isPrivate && (
          <input
            type="password"
            placeholder="Пароль для доступа"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            className="w-full p-3 rounded-lg border border-gray-300 dark:border-gray-600 bg-white dark:bg-gray-800 dark:text-white mb-4"
          />
        )}

        {error && <p className="text-red-500 mb-4">{error}</p>}

        <div className="flex justify-between items-center">
          <button
            onClick={() => navigate("/dashboard")}
            className="bg-gray-300 dark:bg-gray-700 text-black dark:text-white px-4 py-2 rounded-lg hover:opacity-90"
          >
            ⬅ Назад
          </button>

          <button
            onClick={handleCreate}
            className="bg-blue-600 hover:bg-blue-700 text-white px-4 py-2 rounded-lg"
          >
            ✅ Создать
          </button>
        </div>
      </div>
    </div>
  );
}
