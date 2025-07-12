import { useState } from "react";

type Props = {
  onSubmit: (password: string) => void;
};

export default function PublicNotePasswordForm({ onSubmit }: Props) {
  const [password, setPassword] = useState("");

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    onSubmit(password);
  };

  return (
    <form onSubmit={handleSubmit}>
      <h3>Эта заметка приватная</h3>
      <input
        type="password"
        placeholder="Введите пароль"
        value={password}
        onChange={(e) => setPassword(e.target.value)}
        required
      />
      <button type="submit">Показать</button>
    </form>
  );
}
