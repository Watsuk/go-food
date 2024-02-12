import { useState } from "react";
import { Label } from "./ui/label";
import { Input } from "./ui/input";
import { Button } from "./ui/button";
import { createTruck } from "@/api/truck";

export default function CreateTruckForm() {
  const [name, setName] = useState("");
  const [openTime, setOpenTime] = useState("");
  const [closeTime, setCloseTime] = useState("");
  const [slotBuffer, setSlotBuffer] = useState("");
  const [message, setMessage] = useState("");

  const handleSubmit = async (e) => {
    e.preventDefault();

    const token = localStorage.getItem("token");
    const userId = localStorage.getItem("user_id");

    try {
      // Remplace les valeurs par celles appropriées, notamment `userID` si nécessaire
      // const truck = await createTruck(token!, {
      //   name: name,
      //   open_time: openTime,
      //   close_time: closeTime,
      //   slot_buffer: parseFloat(slotBuffer, 10),
      //   user_id: parseFloat(userId),
      // });

      setMessage("Truck created successfully");
      // Réinitialise le formulaire ou navigue vers une autre page si nécessaire
    } catch (error) {
      setMessage(error.message);
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

      <Label>Intervalle de temps entre les slots (en minutes)</Label>
      <Input
        value={slotBuffer}
        onChange={(e) => setSlotBuffer(e.target.value)}
        placeholder="Intervalle"
        type="number"
      />

      <Button type="submit">Créer un camion</Button>
      {message && <div>{message}</div>}
    </form>
  );
}
