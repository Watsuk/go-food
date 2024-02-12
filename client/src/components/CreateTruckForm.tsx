import { useState } from "react";
import { Label } from "./ui/label";
import { Input } from "./ui/input";
import { Button } from "./ui/button";
import { createTruck } from "@/api/truck";

export default function CreateTruckForm() {
  const [name, setName] = useState("");
  const [openTime, setOpenTime] = useState("");
  const [closeTime, setCloseTime] = useState("");
  const [slotBuffer, setSlotBuffer] = useState(0);
  const [message, setMessage] = useState("");

  const handleSubmit = async (e) => {
    e.preventDefault();

    const jwt = localStorage.getItem("token");
    if (!jwt) {
      setMessage(
        "Token d'authentification manquant. Veuillez vous reconnecter."
      );
      return;
    }

    const user_id = Number(localStorage.getItem("user_id"));
    if (isNaN(user_id)) {
      setMessage("Erreur d'identification de l'utilisateur.");
      return;
    }

    try {
      await createTruck(jwt, {
        name,
        user_id,
        slot_buffer: Number(slotBuffer),
        open_time: openTime,
        close_time: closeTime,
      });

      setMessage("Truck créé avec succès.");
      setName("");
      setOpenTime("");
      setCloseTime("");
      setSlotBuffer(0);
    } catch (error) {
      setMessage(
        error instanceof Error
          ? error.message
          : "Erreur lors de la création du camion."
      );
    }
  };

  return (
    <form onSubmit={handleSubmit} className="flex flex-col gap-4">
      <Label>Nom du camion</Label>
      <Input
        value={name}
        onChange={(e) => setName(e.target.value)}
        placeholder="Nom du camion"
      />

      <Label>Heure d'ouverture</Label>
      <Input
        value={openTime}
        onChange={(e) => setOpenTime(e.target.value)}
        placeholder="HH:MM"
        type="time"
      />

      <Label>Heure de fermeture</Label>
      <Input
        value={closeTime}
        onChange={(e) => setCloseTime(e.target.value)}
        placeholder="HH:MM"
        type="time"
      />

      <Label> nombre de commande max en simultané</Label>
      <Input
        value={slotBuffer}
        onChange={(e) => setSlotBuffer(parseInt(e.target.value, 10))}
        placeholder="commande max en simultané"
        type="number"
      />

      <Button type="submit">Créer un camion</Button>
      {message && <div>{message}</div>}
    </form>
  );
}
