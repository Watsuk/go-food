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

export const createTruck = async (jwt: string, truckData: Truck): Promise<Truck> => {
  const response = await fetch(`${API_URL}/create-truck`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      "Authorization": jwt,
    },
    body: JSON.stringify({
      name: truckData.name,
      userID: truckData.user_id,
      slotBuffer: truckData.slot_buffer,
      openTime: truckData.open_time,
      closeTime: truckData.close_time,
    }),
  });

  if (!response.ok) {
    throw new Error("Failed to create truck");
  }

  return response.json();
};

// add jwt into header
export const getTrucks = async (jwt: string): Promise<[]> => {
  const response = await fetch(`${API_URL}/trucks`, {
    method: "GET",
    headers: {
      "Accept": "application/json",
      "Authorization": jwt,
    },

  });

  if (!response.ok) {
    throw new Error("Failed to fetch trucks");
  }

  return response.json();
};
export const getTrucksByTruckId = async (jwt: string, truckId: number): Promise<Truck> => {
  const response = await fetch(`${API_URL}/trucks/${truckId}`, {
    method: "GET",
    headers: {
      "Accept": "application/json",
      "Authorization": jwt,
    },
  });

  if (!response.ok) {
    throw new Error("Failed to fetch truck");
  }

  return response.json();
}