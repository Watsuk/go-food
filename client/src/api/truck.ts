const API_URL = "http://localhost:3000";

export const deleteTruck = async (truckID: number): Promise<void> => {
    const response = await fetch(`${API_URL}/delete-truck/${truckID}`, {
      method: "DELETE", // Utilisez la méthode DELETE pour supprimer une ressource
      headers: {
        "Accept": "application/json",
      },
    });
    if (!response.ok) {
      throw new Error("Failed to delete truck");
    }
    // Aucun corps de réponse attendu pour une requête DELETE, donc pas de .json() à parser
  };
  
  interface Truck {
    name: string;
    userID: number;
    slotBuffer: number;
    openTime: string;
    closeTime: string;
  }
  
  export const createTruck = async (truckData: Truck): Promise<Truck> => {
    const response = await fetch(`${API_URL}/create-truck`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        name: truckData.name,
        userID: truckData.userID,
        slotBuffer: truckData.slotBuffer,
        openTime: truckData.openTime,
        closeTime: truckData.closeTime,
      }),
    });
  
    if (!response.ok) {
      throw new Error("Failed to create truck");
    }
  
    return response.json();
  };
  