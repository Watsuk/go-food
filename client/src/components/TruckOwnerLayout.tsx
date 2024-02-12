import React, { useEffect, useState } from "react";
import { getTrucksByUserId, getTruckMenu } from "@/api/truck"; // Supposons que getTruckMenu est une fonction d'API pour récupérer le menu d'un camion
import CreateTruckForm from "@/components/CreateTruckForm";
import { Truck } from "@/types/type";

const TruckOwnerLayout = () => {
  const [trucks, setTrucks] = useState<Truck[]>([]);
  const [loading, setLoading] = useState(true);
  const [selectedTruck, setSelectedTruck] = useState<Truck | null>(null);

  useEffect(() => {
    const fetchTrucks = async () => {
      const jwt = localStorage.getItem("token");
      const userID = parseInt(localStorage.getItem("user_id") || "0");

      if (!jwt || userID === 0) {
        console.error("JWT or UserID missing");
        setLoading(false);
        return;
      }

      try {
        const trucksData = await getTrucksByUserId(jwt, userID);
        setTrucks(trucksData || []);
      } catch (error) {
        console.error("Failed to fetch trucks data:", error);
      } finally {
        setLoading(false);
      }
    };

    fetchTrucks();
  }, []);

  const handleTruckSelect = async (truckId: number) => {
    const truck = trucks.find((t) => t.id === truckId);
    if (truck) {
      setSelectedTruck(truck);
    }
  };

  const handleDeselectTruck = () => {
    setSelectedTruck(null);
  };

  if (loading) {
    return <div>Chargement...</div>;
  }

  return (
    <div className="h-full w-full text-black p-4">
      {selectedTruck ? (
        <>
          <button onClick={handleDeselectTruck}>Retour</button>
        </>
      ) : trucks.length > 0 ? (
        <>
          <h1 className="text-2xl font-bold">Vos Trucks</h1>
          <div className="flex flex-wrap gap-4">
            {trucks.map((truck) => (
              <div key={truck.id} className="border p-4">
                <h2>{truck.name}</h2>
                <p>Heure d'ouverture: {truck.open_time}</p>
                <p>Heure de fermeture: {truck.close_time}</p>
                <p>Nombre de places: {truck.slot_buffer}</p>
                <button onClick={() => handleTruckSelect(truck.id)}>
                  Voir le Menu
                </button>
              </div>
            ))}
          </div>
          <CreateTruckForm />
        </>
      ) : (
        <>
          <h1 className="text-2xl font-bold">Vous n'avez pas de camion</h1>
          <CreateTruckForm />
        </>
      )}
    </div>
  );
};

export default TruckOwnerLayout;
