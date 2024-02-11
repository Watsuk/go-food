import { Label } from "@/components/ui/label";
import { useState, FormEvent } from "react";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { signin } from "@/api/authService";

export default function Sigin() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [message, setMessage] = useState("");

  const handleSubmit = async (event: FormEvent) => {
    event.preventDefault();
    try {
      const data = await signin(email, password);
      if (data) {
        setMessage("Signin successful");
        localStorage.setItem("token", data.token);
        setTimeout(() => {
          window.location.href = "/dashboard";
        }, 2000);
      }
    } catch (error) {
      setMessage(error instanceof Error ? error.message : "Unknown error");
    }
  };

  return (
    <form onSubmit={handleSubmit} className="flex flex-col gap-4">
      <div>
        <Label>Email</Label>
        <Input
          value={email}
          onChange={(e) => setEmail(e.target.value)}
          placeholder="Enter your email"
          type="email"
        />
      </div>
      <div>
        <Label>Password</Label>
        <Input
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          placeholder="Enter your password"
          type="password"
        />
      </div>
      <Button type="submit">Sign in</Button>
      {message && <div>{message}</div>}
    </form>
  );
}
