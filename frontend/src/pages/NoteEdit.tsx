import { useEffect, useState } from "react";
import { useNavigate, useParams } from "react-router-dom";
import api from "../services/api";
import { useTheme } from "../hooks/useTheme";

export default function NoteEdit() {
  const { id } = useParams();
  const navigate = useNavigate();
  const { theme } = useTheme();

  const [content, setContent] = useState("");
  const [expiresAt, setExpiresAt] = useState("");
  const [isPrivate, setIsPrivate] = useState(false);
  const [password, setPassword] = useState("");
  const [error, setError] = useState("");
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchNote = async () => {
      try {
        const res = await api.get(`/notes/${id}`);
        const note = res.data;
        setContent(note.content);
        setExpiresAt(new Date(note.expires_at).toISOString().slice(0, 16));
        setIsPrivate(note.is_private);
      } catch (err) {
        setError("Ошибка загрузки заметки");
      } finally {
        setLoading(false);
      }
    };

    fetchNote();
  }, [id]);

  const handleUpdate = async () => {
    setError("");

    if (isPrivate && password.trim() === "") {
      setError("Введите новый пароль для приватной заметки");
      return;
    }

    try {
      await api.put(`/notes/${id}`, {
        content,
        expires_at: new Date(expiresAt).toISOString(),
        is_private: isPrivate,
        password: isPrivate ? password : undefined,
      });
      navigate("/dashboard");
    } catch (err: any) {
      setError(err.response?.data || "Ошибка при обновлении заметки");
    }
  };

  if (loading) {
    return <p className="text-center mt-10">Загрузка...</p>;
  }

  return (
    <div className="relative flex items-center justify-center min-h-screen bg-[var(--bg-color)] text-[var(--text-color)] transition-colors px-4 overflow-hidden">
      <div
        className={`pointer-events-none absolute top-0 left-1/2 transform -translate-x-1/2 w-[600px] h-[600px] rounded-full blur-[140px] opacity-20 z-0 ${
          theme === "dark" ? "bg-purple-700" : "bg-blue-300"
        }`}
      />

      <div className="relative z-10 w-full max-w-lg bg-[var(--card-bg)] p-8 rounded-2xl shadow-xl">
        <h2 className="text-2xl font-bold mb-6">Редактировать заметку</h2>

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
            placeholder="Новый пароль (если нужно изменить)"
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
            onClick={handleUpdate}
            className="bg-blue-600 hover:bg-blue-700 text-white px-4 py-2 rounded-lg"
          >
            💾 Сохранить
          </button>
        </div>
      </div>
    </div>
  );
}