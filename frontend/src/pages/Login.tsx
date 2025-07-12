// src/pages/Login.tsx
import { useNavigate } from "react-router-dom";
import LoginForm from "../components/LoginForm";
import { useTheme } from "../hooks/useTheme";

export default function Login() {
  const navigate = useNavigate();
  const { toggleTheme, theme } = useTheme();

  return (
    <div className="relative flex items-center justify-center min-h-screen bg-[var(--bg-color)] text-[var(--text-color)] transition-colors px-4 overflow-hidden">
      {/* 🔵 Анимация фона */}
      <div
        className={`pointer-events-none absolute top-0 left-1/2 transform -translate-x-1/2 w-[500px] h-[500px] rounded-full blur-[120px] opacity-30 z-0 ${
          theme === "dark" ? "bg-purple-500" : "bg-blue-300"
        }`}
      />

      {/* ⬅ Кнопка назад */}
      <div className="absolute top-4 left-4 z-10">
        <button
          onClick={() => navigate("/")}
          className="bg-[var(--card-bg)] hover:bg-[var(--button-hover)] text-[var(--text-color)] border border-gray-400 px-3 py-1 rounded-xl shadow-sm transition"
        >
          ⬅ Назад
        </button>
      </div>

      {/* 🌓 Сменить тему */}
      <div className="absolute top-4 right-4 z-10">
        <button
          onClick={toggleTheme}
          className="bg-[var(--card-bg)] hover:bg-[var(--button-hover)] text-[var(--text-color)] border border-gray-400 px-3 py-1 rounded-xl shadow-sm transition"
        >
          🌓 Сменить тему
        </button>
      </div>

      <LoginForm />
    </div>
  );
}
