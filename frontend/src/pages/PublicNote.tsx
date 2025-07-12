import { useParams } from "react-router-dom";
import { useEffect, useState } from "react";
import axios from "axios";

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

  if (loading) return <div>Загрузка...</div>;
  if (error) return <div style={{ color: "red" }}>{error}</div>;
  if (!note) return <div>Заметка не найдена</div>;

  return (
    <div>
      <h1>Публичная заметка</h1>
      <p>
        <strong>Содержимое:</strong> {note.content}
      </p>
      <p>
        <strong>Создана:</strong> {formatDate(note.created_at)}
      </p>
      <p>
        <strong>Истекает:</strong> {formatDate(note.expires_at)}
      </p>
    </div>
  );
}
