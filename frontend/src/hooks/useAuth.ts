export function useAuth() {
  let user = null;
  try {
    user = JSON.parse(localStorage.getItem("user") || "null");
  } catch {
    user = null;
  }

  const logout = () => {
    localStorage.removeItem("user");
    window.location.reload();
  };

  return { user, logout };
}
