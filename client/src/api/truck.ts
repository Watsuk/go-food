import { Truck } from "@/types/type";

const API_URL = "http://localhost:3000";

export const deleteTruck = async (jwt: string, truckID: number): Promise<void> => {
  const response = await fetch(`${API_URL}/delete-truck/${truckID}`, {
    method: "DELETE",
    headers: {
      "Accept": "application/json",
      "Authorization": jwt,
    },
  });
  if (!response.ok) {
    throw new Error("Failed to delete truck");
  }
  return;
};

export const createTruck = async (jwt: string, truckData): Promise<any> => {
  const response = await fetch(`${API_URL}/create-truck`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      "Authorization": jwt,
    },
    body: JSON.stringify({
      name: truckData.name,
      user_id: truckData.user_id,
      slot_buffer: truckData.slot_buffer,
      open_time: truckData.open_time,
      close_time: truckData.close_time,
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

export const editTruck = async (jwt: string, truckID: number): Promise<void> => {
  const response = await fetch(`${API_URL}/trucks/${truckID}`, {
    method: "PATCH",
    headers: {
      "Accept": "application/json",
      "Authorization": jwt,
    },
  });
  if (!response.ok) {
    throw new Error("Failed to edit truck");
  }
  return;
};

export const getNumberCurrentOrdersByTruckID = async (jwt: string, truckId: number): Promise<void> => {
  const response = await fetch(`${API_URL}/truck/${truckId}/number-current-orders`, {
    method: "GET",
    headers: {
      "Accept": "application/json",
      "Authorization": jwt,
    },
  });

  if (!response.ok) {
    throw new Error("Failed to fetch order currently use by the truck");
  }

  return response.json();
}

export const getTrucksByUserId = async (jwt: string, userID: number): Promise<Truck[]> => {
  const response = await fetch(`${API_URL}/trucks/user/${userID}`, {
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