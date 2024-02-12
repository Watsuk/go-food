const API_URL = "http://localhost:3000";

interface OrderData {
  userID: number;
  truckID: number;
  products: number[];
  quantities: number[];
  comment: string;
  hour: string;
}

export const createOrder = async (jwt: string, orderData: OrderData): Promise<void> => {
  const response = await fetch(`${API_URL}/order`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      "Authorization": jwt,
    },
    body: JSON.stringify({
      userID: orderData.userID,
      truckID: orderData.truckID,
      products: orderData.products,
      quantities: orderData.quantities,
      comment: orderData.comment,
      hour: orderData.hour,
    }),
  });

  if (!response.ok) {
    throw new Error("Failed to create order");
  }
};

export const acceptOrder = async (jwt: string, orderID: number, accept: boolean): Promise<void> => {
  const status = accept ? "1" : "0"; // Supposons que "1" représente accepter et "0" représenter rejeter pour l'URL
  const response = await fetch(`${API_URL}/order/${orderID}/accept/${status}`, {
    method: "PATCH",
    headers: {
      "Accept": "application/json",
      "Authorization": jwt,
    },
  });

  if (!response.ok) {
    throw new Error("Failed to update order status");
  }
};

interface Order {
  id: number;
  userID: number;
  truckID: number;
  price: number;
  hours: string;
  status: string;
  orderData: any;
  updatedAt: string;
  deletedAt?: string;
}

export const getOrderById = async (jwt: string, orderID: number): Promise<Order> => {
  const response = await fetch(`${API_URL}/order/${orderID}`, {
    method: "GET",
    headers: {
      "Accept": "application/json",
      "Authorization": jwt,
    },
  });

  if (!response.ok) {
    throw new Error("Failed to fetch order");
  }

  return response.json();
};

export const getOrdersByTruck = async (jwt: string, truckID: number): Promise<Order[]> => {
  const response = await fetch(`${API_URL}/orders/truck/${truckID}`, {
    method: "GET",
    headers: {
      "Accept": "application/json",
      "Authorization": jwt,
    },
  });

  if (!response.ok) {
    throw new Error("Failed to fetch orders");
  }

  return response.json();
};

export const getOrdersByUser = async (jwt: string, userID: number): Promise<Order[]> => {
  const response = await fetch(`${API_URL}/orders/user/${userID}`, {
    method: "GET",
    headers: {
      "Accept": "application/json",
      "Authorization": jwt,

    },
  });

  if (!response.ok) {
    throw new Error("Failed to fetch orders");
  }

  return response.json();
};

export const getOrders = async (jwt: string): Promise<Order[]> => {
  const response = await fetch(`${API_URL}/orders`, {
    method: "GET",
    headers: {
      "Accept": "application/json",
      "Authorization": jwt,
    },
  });

  if (!response.ok) {
    throw new Error("Failed to fetch orders");
  }

  return response.json();
};

export const completeOrder = async (jwt: string, orderID: number): Promise<void> => {
  const response = await fetch(`${API_URL}/order/${orderID}/completed`, {
    method: "PATCH",
    headers: {
      "Accept": "application/json",
      "Authorization": jwt,
    },
  });

  if (!response.ok) {
    throw new Error(`Failed to complete order with ID ${orderID}`);
  }
};

export const handOverOrder = async (jwt: string, orderID: number): Promise<void> => {
  const response = await fetch(`${API_URL}/order/${orderID}/handedover`, {
    method: "PATCH",
    headers: {
      "Accept": "application/json",
      "Authorization": jwt,
    },
  });

  if (!response.ok) {
    throw new Error(`Failed to hand over order with ID ${orderID}`);
  }
};

export const deleteOrder = async (jwt: string, orderID: number): Promise<void> => {
  const response = await fetch(`${API_URL}/order/${orderID}`, {
    method: "DELETE",
    headers: {
      "Accept": "application/json",
      "Authorization": jwt,
    },
  });

  if (!response.ok) {
    throw new Error("Failed to delete order");
  }

  return response.json();
};