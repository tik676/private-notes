import { useParams, useNavigate } from "react-router-dom";
import { useEffect, useState } from "react";
import axios from "axios";
import { useTheme } from "../hooks/useTheme";

type Note = {
  id: number;
  content: string;
  created_at: string;
  expires_at: string;
  is_private: boolean;
};

export default function PublicNote() {
  const { id } = useParams();
  const [note, setNote] = useState<Note | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");
  const [backPath, setBackPath] = useState("/");
  const { theme } = useTheme();
  const navigate = useNavigate();

  useEffect(() => {
    const token = localStorage.getItem("access_token");
    setBackPath(token ? "/dashboard" : "/");
  }, []);

  useEffect(() => {
    const fetchNote = async () => {
      try {
        const res = await axios.get(`http://localhost:2288/notes/public/${id}`);
        setNote(res.data);
      } catch (err) {
        setError("Заметка не найдена или доступ запрещён");
      } finally {
        setLoading(false);
      }
    };

    fetchNote();
  }, [id]);

  const formatDate = (date: string) =>
    new Date(date).toLocaleString("ru-RU", { timeZone: "Asia/Almaty" });

  return (
    <div className="min-h-screen flex items-center justify-center bg-[var(--bg-color)] text-[var(--text-color)] px-4 transition-colors">
      <div
        className={`pointer-events-none absolute top-0 left-1/2 transform -translate-x-1/2 w-[600px] h-[600px] rounded-full blur-[140px] opacity-20 z-0 ${
          theme === "dark" ? "bg-purple-700" : "bg-blue-300"
        }`}
      />

      <div className="relative z-10 w-full max-w-md bg-[var(--card-bg)] p-6 rounded-2xl shadow-xl">
        <h2 className="text-xl font-bold mb-4">📄 Публичная заметка</h2>

        {loading ? (
          <p>Загрузка...</p>
        ) : error ? (
          <p className="text-red-500">{error}</p>
        ) : note ? (
          <>
            <p className="whitespace-pre-wrap break-words bg-gray-100 dark:bg-gray-700 p-3 rounded-md mb-4">
              {note.content}
            </p>
            <p className="text-sm mb-2">
              🕓 <strong>Создана:</strong> {formatDate(note.created_at)}
            </p>
            <p className="text-sm mb-4">
              ⏳ <strong>Истекает:</strong> {formatDate(note.expires_at)}
            </p>
            <button
              onClick={() => navigate(backPath)}
              className="bg-gray-300 dark:bg-gray-700 text-black dark:text-white w-full py-2 rounded-lg"
            >
              ⬅ Назад
            </button>
          </>
        ) : (
          <p>Заметка не найдена</p>
        )}
      </div>
    </div>
  );
}
