import { Truck } from "@/types/type";

const API_URL = "http://localhost:3000";

export const deleteTruck = async (truckID: number): Promise<void> => {
  const response = await fetch(`${API_URL}/delete-truck/${truckID}`, {
    method: "DELETE",
    headers: {
      "Accept": "application/json",
    },
  });
  if (!response.ok) {
    throw new Error("Failed to delete truck");
  }
  return;
};

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

export const getTrucks = async (): Promise<Truck[]> => {
  const response = await fetch(`${API_URL}/trucks`, {
    method: "GET",
    headers: {
      "Accept": "application/json",

    },

  });

  if (!response.ok) {
    throw new Error("Failed to fetch trucks");
  }

  return response.json();
};

export const getTrucksByTruckId = async (truckId: number): Promise<Truck> => {
  const response = await fetch(`${API_URL}/trucks/${truckId}`, {
    method: "GET",
    headers: {
      "Accept": "application/json",
    },
  });

  if (!response.ok) {
    throw new Error("Failed to fetch truck");
  }

  return response.json();
}