import { Routes, Route } from "react-router-dom";
import { useEffect, useState } from "react";

import Register from "./pages/Register";
import Login from "./pages/Login";
import CreateNote from "./pages/CreateNote";
import PublicNote from "./pages/PublicNote";
import NoteEdit from "./pages/NoteEdit";
import Welcome from "./pages/Welcome";
import Dashboard from "./pages/Dashboard";

import "./index.css";
import "./App.css";

function App() {
  const [user, setUser] = useState<any>(null);

  useEffect(() => {
    try {
      const stored = localStorage.getItem("user");
      if (stored) setUser(JSON.parse(stored));
    } catch {
      setUser(null);
    }
  }, []);

  return (
    <div className="min-h-screen bg-[var(--bg-color)] text-[var(--text-color)] transition-colors">
      <Routes>
        <Route path="/" element={user ? <Dashboard /> : <Welcome />} />
        <Route path="/dashboard" element={<Dashboard />} />
        <Route path="/register" element={<Register />} />
        <Route path="/login" element={<Login />} />
        <Route path="/create" element={<CreateNote />} />
        <Route path="/note/:id" element={<PublicNote />} />
        <Route path="/edit/:id" element={<NoteEdit />} />
      </Routes>
    </div>
  );

}

export default App;
