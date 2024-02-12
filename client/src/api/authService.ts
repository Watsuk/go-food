// authService.ts
const API_URL = "http://localhost:3000";

export const signin = async (email: string, password: string) => {
  const response = await fetch(`${API_URL}/login`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ email, password }),
  });
  if (!response.ok) {
    throw new Error("Signin failed");
  }
  return response.json();
};

export const register = async (
  userName: string,
  email: string,
  password: string,
  userRole: number
) => {
  const response = await fetch(`${API_URL}/register`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ username: userName, email, password, role: userRole }),
  });
  if (!response.ok) {
    throw new Error("Registration failed");
  }
  return response.json();
};

// export const signout = async () => {
//   localStorage.removeItem("token");
// };
