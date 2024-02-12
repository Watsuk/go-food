const API_URL = "http://localhost:3000";

export const deleteAccount = async () => {
  //const token = localStorage.getItem("token");
  const response = await fetch(`${API_URL}/delete-account`, {
    method: "DELETE",
    headers: {
      "Content-Type": "application/json",
    },
  });
  if (!response.ok) {
    throw new Error("Delete account failed");
  }
  console.log(response.json());
  return response.json();
};

export const getUsers = async () => {
  const response = await fetch(`${API_URL}/users`, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
  });
  if (!response.ok) {
    throw new Error("Failed to fetch users");
  }
  return response.json();
};

export const getUserById = async (jwt: string, userID: number) => {
  const response = await fetch(`${API_URL}/user/${userID}`, {
    method: "GET",
    headers: {
      "Accept": "application/json",
      "Authorization": jwt,

    },
  });
  if (!response.ok) {
    throw new Error("Failed to fetch user");
  }
  return response.json();
};


