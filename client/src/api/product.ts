const API_URL = "http://localhost:3000";

interface ProductData {
  truckID: number;
  label: string;
  description: string;
  price: number;
}

export const createProduct = async (jwt: string, productData: ProductData): Promise<ProductData> => {
  const response = await fetch(`${API_URL}/product`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      "Authorization": jwt,
    },
    body: JSON.stringify({
      truck_id: productData.truckID,
      label: productData.label,
      description: productData.description,
      price: productData.price,
    }),
  });

  if (!response.ok) {
    throw new Error("Failed to create product");
  }

  return response.json();
};

interface Product {
  id: number;
  truckID: number;
  label: string;
  description: string;
  price: number;
  createdAt: string;
  updatedAt: string;
}

export const getProductById = async (jwt: string, productID: number): Promise<Product> => {
  const response = await fetch(`${API_URL}/product/${productID}`, {
    method: "GET",
    headers: {
      "Accept": "application/json",
      "Authorization": jwt,
    },
  });

  if (!response.ok) {
    throw new Error("Failed to fetch product");
  }

  return response.json();
};

export const getProductByTruckId = async (jwt: string, truckID: number): Promise<Product> => {
  const response = await fetch(`${API_URL}/products/truck/${truckID}`, {
    method: "GET",
    headers: {
      "Accept": "application/json",
      "Authorization": jwt,
    },
  });

  if (!response.ok) {
    throw new Error("Failed to fetch products");
  }

  return response.json();
};

export const deleteProduct = async (jwt: string, productID: number): Promise<Product> => {
  const response = await fetch(`${API_URL}/product/${productID}`, {
    method: "DELETE",
    headers: {
      "Accept": "application/json",
      "Authorization": jwt,
    },
  });

  if (!response.ok) {
    throw new Error("Failed to delete product");
  }

  return response.json();
};