import { deleteAccount } from "@/api/user";
import { Button } from "@/components/ui/button";
import { useState } from "react";

export default function Sigin() {
  const [message, setMessage] = useState("");

  const handleDelete = async () => {
    try {
      const data = await deleteAccount();
      if (data) {
        setMessage("Account deleted");
        localStorage.removeItem("JWT_TOKEN");
      }
    } catch (error) {
      setMessage(error instanceof Error ? error.message : "Unknown error");
    }
  };
  return (
    <div className="w-full h-full flex flex-col items-center justify-center gap-4">
      <h1>Dashboard</h1>
      <div className="h-[60vh] flex flex-col items-center justify-center gap-4">
        <Button onClick={handleDelete}>delete account</Button>
        {message && <div>{message}</div>}
      </div>
    </div>
  );
}
