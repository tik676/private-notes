import { useState, useEffect } from "react";
import { useParams, useNavigate, useLocation } from "react-router-dom";
import api from "../services/api";
import { useTheme } from "../hooks/useTheme";

export default function UnlockNote() {
  const { id } = useParams();
  const [password, setPassword] = useState("");
  const [error, setError] = useState("");
  const [note, setNote] = useState<any>(null);
  const { theme } = useTheme();
  const navigate = useNavigate();
  const location = useLocation();

  const [backPath, setBackPath] = useState("/");

  useEffect(() => {
    const token = localStorage.getItem("access_token");
    setBackPath(token ? "/dashboard" : "/");
  }, []);

  const handleUnlock = async () => {
    setError("");
    try {
      const res = await api.post(`/notes/${id}/unlock`, { password });
      setNote(res.data);

      const userRaw = localStorage.getItem("user");
      const user = userRaw ? JSON.parse(userRaw) : null;
      const isRedirectToEdit = new URLSearchParams(location.search).get("redirect") === "edit";

      if (user && res.data.user_id === user.id && isRedirectToEdit) {
        navigate(`/edit/${id}`);
      }
    } catch (err: any) {
      console.log("Ошибка при разблокировке:", err);
      const responseData = err.response?.data;

      if (typeof responseData === "string") {
        setError(responseData);
      } else if (typeof responseData === "object" && responseData !== null && "error" in responseData) {
        setError(responseData.error);
      } else {
        setError("Ошибка разблокировки");
      }
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
        <h2 className="text-xl font-bold mb-4">🔐 Введите пароль для заметки</h2>

        {!note ? (
          <>
            <input
              type="password"
              placeholder="Пароль"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              className="w-full mb-4 p-3 rounded-lg border border-gray-300 dark:border-gray-600 bg-white dark:bg-gray-800 dark:text-white"
            />
            {error && <p className="text-red-500 mb-3">{error}</p>}

            <button
              onClick={handleUnlock}
              className="bg-blue-600 hover:bg-blue-700 text-white w-full py-2 rounded-lg"
            >
              🔓 Разблокировать
            </button>
          </>
        ) : (
          <div>
            <h3 className="text-lg font-semibold mb-2">Содержимое:</h3>
            <p className="whitespace-pre-wrap break-words bg-gray-100 dark:bg-gray-700 p-3 rounded-md mb-4">
              {note.content}
            </p>
            <button
              onClick={() => navigate(backPath)}
              className="bg-gray-300 dark:bg-gray-700 text-black dark:text-white w-full py-2 rounded-lg"
            >
              ⬅ Назад
            </button>
          </div>
        )}
      </div>
    </div>
  );
}
