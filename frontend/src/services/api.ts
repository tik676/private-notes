import axios from "axios";

const API_URL = "http://localhost:2288";

const api = axios.create({
  baseURL: API_URL,
});

// 👉 Добавляем токен к каждому запросу
api.interceptors.request.use((config) => {
  const token = localStorage.getItem("access_token");
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
});

// 🔄 Автоматическое обновление токена по refresh_token
api.interceptors.response.use(
  (res) => res,
  async (err) => {
    const originalRequest = err.config;

    if (
      err.response?.status === 401 &&
      !originalRequest._retry &&
      localStorage.getItem("refresh_token")
    ) {
      originalRequest._retry = true;
      try {
        const res = await axios.post(`${API_URL}/refresh-token`, {
          refresh_token: localStorage.getItem("refresh_token"),
        });

        const { access_token, refresh_token } = res.data;
        localStorage.setItem("access_token", access_token);
        localStorage.setItem("refresh_token", refresh_token);

        originalRequest.headers.Authorization = `Bearer ${access_token}`;
        return api(originalRequest);
      } catch (refreshErr) {
        localStorage.removeItem("access_token");
        localStorage.removeItem("refresh_token");
        window.location.href = "/login";
        return Promise.reject(refreshErr);
      }
    }

    return Promise.reject(err);
  }
);

// 💾 Получение своих заметок
export const getMyNotes = async () => {
  const res = await api.get("/me");
  return res.data;
};

// 🔐 Экспорт по умолчанию для вызова других методов
export default api;
