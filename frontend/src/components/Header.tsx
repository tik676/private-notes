import { useEffect, useState } from "react";
import { Link, useNavigate } from "react-router-dom";
import { useTheme } from "../hooks/useTheme";

export default function Header() {
  const { toggleTheme, theme } = useTheme();
  const [user, setUser] = useState<{ id: number; name: string } | null>(null);
  const navigate = useNavigate();

  useEffect(() => {
    try {
      const stored = localStorage.getItem("user");
      if (stored) setUser(JSON.parse(stored));
    } catch {
      setUser(null);
    }
  }, []);

  const handleLogout = () => {
    localStorage.removeItem("access_token");
    localStorage.removeItem("refresh_token");
    localStorage.removeItem("user");
    navigate("/login");
  };

  return (
    <header className="relative z-10 w-full px-4 py-4 shadow-md bg-[var(--card-bg)] text-[var(--text-color)] flex justify-between items-center">
      <div className="text-xl font-bold select-none pointer-events-none">TimeVault</div>

      <div className="flex items-center gap-3">
        <button
          onClick={toggleTheme}
          className="text-sm bg-[var(--button-bg)] hover:bg-[var(--button-hover)] text-white px-3 py-1 rounded-xl transition shadow"
        >
          {theme === "dark" ? "🌞 Светлая" : "🌙 Тёмная"}
        </button>

        {user ? (
          <>
            <span className="text-sm hidden sm:inline">
              👤 {user.name} (ID: {user.id})
            </span>
            <Link
              to="/dashboard"
              className="bg-blue-600 hover:bg-blue-700 text-white px-3 py-1 rounded-xl text-sm transition shadow hidden sm:inline"
            >
              🗒 Мои заметки
            </Link>
            <Link
              to="/create"
              className="bg-green-600 hover:bg-green-700 text-white px-3 py-1 rounded-xl text-sm transition shadow"
            >
              ➕ Новая заметка
            </Link>
            <button
              onClick={handleLogout}
              className="bg-red-600 hover:bg-red-700 text-white px-3 py-1 rounded-xl text-sm transition shadow"
            >
              🚪 Выйти
            </button>
          </>
        ) : (
          <>
            <Link
              to="/login"
              className="bg-blue-600 hover:bg-blue-700 text-white px-3 py-1 rounded-xl text-sm transition shadow"
            >
              Войти
            </Link>
            <Link
              to="/register"
              className="bg-green-600 hover:bg-green-700 text-white px-3 py-1 rounded-xl text-sm transition shadow"
            >
              Зарегистрироваться
            </Link>
          </>
        )}
      </div>
    </header>
  );
}
