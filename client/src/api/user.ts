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
