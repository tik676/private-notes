import { useEffect, useState } from "react";
import { getMyNotes } from "../services/api";

type Note = {
  id: number;
  content: string;
  is_private: boolean;
  created_at: string;
  expires_at: string;
};

export default function NoteList() {
  const [notes, setNotes] = useState<Note[]>([]);
  const [error, setError] = useState("");

  useEffect(() => {
    const fetchNotes = async () => {
      try {
        const data = await getMyNotes();
        setNotes(data);
      } catch (err) {
        setError("Ошибка загрузки заметок");
      }
    };

    fetchNotes();
  }, []);

  if (error) {
    return <div>{error}</div>;
  }

  return (
    <div>
      <h2>Мои заметки</h2>
      {notes.length === 0 && <p>Нет заметок</p>}
      <ul>
        {notes.map((note) => (
          <li key={note.id}>
            <strong>{note.is_private ? "🔒 " : "🌐 "}</strong>
            {note.content}
          </li>
        ))}
      </ul>
    </div>
  );
}
