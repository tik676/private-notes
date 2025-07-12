import { useNavigate } from "react-router-dom";
import { useTheme } from "../hooks/useTheme";
import { motion } from "framer-motion";

export default function Welcome() {
  const navigate = useNavigate();
  const { theme, toggleTheme } = useTheme();

  return (
    <div className="relative min-h-screen bg-[var(--bg-color)] text-[var(--text-color)] transition-colors overflow-hidden flex items-center justify-center px-4">
      {/* 🔵 Анимированные линии на фоне */}
      <div
        className="animated-lines"
        style={{
          "--line-color": theme === "light" ? "#c9e7ff" : "#2c2c3a",
        } as React.CSSProperties}
      />

      {/* 🔘 Кнопка смены темы */}
      <div className="absolute top-4 right-4 z-10">
        <motion.button
          whileTap={{ scale: 0.95 }}
          onClick={toggleTheme}
          className="bg-[var(--card-bg)] hover:bg-[var(--button-hover)] text-[var(--text-color)] border border-gray-400 px-3 py-1 rounded-xl shadow-sm transition"
        >
          🌓 Сменить тему
        </motion.button>
      </div>

      {/* 💬 Контент Welcome */}
      <motion.div
        className="relative z-10 max-w-5xl w-full flex flex-col items-center gap-10"
        initial={{ opacity: 0, y: 30 }}
        animate={{ opacity: 1, y: 0 }}
        transition={{ duration: 0.6 }}
      >
        <div className="text-center">
          <h1 className="text-4xl md:text-5xl font-bold mb-4">
            Приватные заметки — только для вас
          </h1>
          <p className="mb-6 text-lg text-gray-600 dark:text-gray-300">
            Создавайте заметки, защищайте паролем и контролируйте срок их жизни. Всё просто.
          </p>
          <div className="flex flex-col sm:flex-row gap-4 justify-center">
            <button
              onClick={() => navigate("/register")}
              className="bg-blue-600 hover:bg-blue-700 text-white px-6 py-3 rounded-lg transition"
            >
              🚀 Зарегистрироваться
            </button>
            <button
              onClick={() => navigate("/login")}
              className="bg-gray-300 hover:bg-gray-400 text-black dark:bg-gray-700 dark:text-white px-6 py-3 rounded-lg transition"
            >
              🔐 Войти
            </button>
          </div>
        </div>
      </motion.div>
    </div>
  );
}