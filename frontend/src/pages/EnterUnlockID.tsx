import { useState } from "react";
import { useNavigate } from "react-router-dom";
import api from "../services/api";
import { useTheme } from "../hooks/useTheme";

export default function EnterUnlockID() {
  const [id, setId] = useState("");
  const [error, setError] = useState("");
  const { theme } = useTheme();
  const navigate = useNavigate();

  const handleCheckNote = async () => {
    setError("");

    if (!id.trim()) {
      setError("Введите ID заметки");
      return;
    }

    try {
      const res = await api.get(`/notes/${id}/check`);
      const isPrivate = res.data?.is_private;

      if (isPrivate) {
        navigate(`/note/${id}/unlock`);
      } else {
        navigate(`/note/${id}`);
      }
    } catch (err: any) {
      const msg = err.response?.data || "Ошибка при проверке заметки";
      setError(typeof msg === "string" ? msg : "Заметка не найдена");
    }
  };

  return (
    <div className="min-h-screen flex items-center justify-center bg-[var(--bg-color)] text-[var(--text-color)] px-4 transition-colors">
      <div
        className={`pointer-events-none absolute top-0 left-1/2 transform -translate-x-1/2 w-[600px] h-[600px] rounded-full blur-[140px] opacity-20 z-0 ${
          theme === "dark" ? "bg-purple-700" : "bg-blue-300"
        }`}
      />
      <div className="relative z-10 w-full max-w-md bg-[var(--card-bg)] p-6 rounded-2xl shadow-xl">
        <h2 className="text-xl font-bold mb-4">🔍 Введите ID заметки</h2>

        <input
          type="text"
          placeholder="ID"
          value={id}
          onChange={(e) => setId(e.target.value)}
          className="w-full mb-4 p-3 rounded-lg border border-gray-300 dark:border-gray-600 bg-white dark:bg-gray-800 dark:text-white"
        />

        {error && <p className="text-red-500 mb-3">{error}</p>}

        <button
          onClick={handleCheckNote}
          className="bg-blue-600 hover:bg-blue-700 text-white w-full py-2 rounded-lg"
        >
          Проверить и перейти
        </button>
      </div>
    </div>
  );
}
