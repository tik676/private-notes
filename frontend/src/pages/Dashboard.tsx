import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import Header from "../components/Header";
import api from "../services/api";

interface Note {
  id: number;
  content: string;
  expires_at: string;
  is_private: boolean;
}

export default function Dashboard() {
  const navigate = useNavigate();
  const [notes, setNotes] = useState<Note[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");

  useEffect(() => {
    const token = localStorage.getItem("access_token");
    if (!token) {
      navigate("/login");
      return;
    }

    const fetchNotes = async () => {
      try {
        const res = await api.get("/me");
        setNotes(res.data);
      } catch (err) {
        console.error("Ошибка при получении заметок:", err);
        setError("Ошибка загрузки заметок");
      } finally {
        setLoading(false);
      }
    };

    fetchNotes();
  }, []);

  const deleteNote = async (id: number) => {
    try {
      await api.delete(`/notes/${id}`);
      setNotes(notes.filter((note) => note.id !== id));
    } catch (err) {
      console.error("Ошибка при удалении заметки:", err);
      alert("Не удалось удалить заметку");
    }
  };

  return (
    <div className="relative min-h-screen bg-[var(--bg-color)] text-[var(--text-color)] px-4 py-8 transition-colors overflow-hidden">
      {/* Анимация фона */}
      <div className="pointer-events-none absolute top-0 left-1/2 transform -translate-x-1/2 w-[700px] h-[700px] rounded-full blur-[140px] opacity-20 z-0 bg-gradient-to-br from-blue-400 to-purple-600 dark:from-purple-800 dark:to-indigo-600" />

      <Header />
      <main className="relative z-10 max-w-5xl mx-auto pt-6">
        <h2 className="text-3xl font-bold mb-6">📓 Мои заметки</h2>

        {loading ? (
          <p className="text-center">Загрузка...</p>
        ) : error ? (
          <p className="text-red-500 text-center">{error}</p>
        ) : notes.length === 0 ? (
          <p className="text-gray-500 text-center">У вас пока нет заметок</p>
        ) : (
          <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
            {notes.map((note) => (
              <div
                key={note.id}
                className="bg-[var(--card-bg)] text-[var(--text-color)] rounded-xl shadow-md p-5 flex flex-col justify-between"
              >
                <div>
                  <p className="text-lg break-words whitespace-pre-wrap mb-3">{note.content}</p>
                  <p className="text-xs text-gray-500">
                    ⏳ Срок: {new Date(note.expires_at).toLocaleString("ru-RU")}
                  </p>
                  <p className="text-xs">
                    🔒 {note.is_private ? "Приватная" : "Публичная"}
                  </p>
                </div>
                <div className="flex gap-2 mt-4">
                  <button
                    onClick={() => navigate(`/edit/${note.id}`)}
                    className="bg-blue-500 hover:bg-blue-600 text-white px-3 py-1 rounded"
                  >
                    ✏️ Редактировать
                  </button>
                  <button
                    onClick={() => deleteNote(note.id)}
                    className="bg-red-500 hover:bg-red-600 text-white px-3 py-1 rounded"
                  >
                    🗑 Удалить
                  </button>
                </div>
              </div>
            ))}
          </div>
        )}
      </main>
    </div>
  );
}